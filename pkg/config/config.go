package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

// Load loads configurations into the given structPtr.
func Load(structPtr interface{}, opts ...Option) error {
	l := &viperLoader{
		viper:       viper.New(),
		intoPtr:     structPtr,
		useDefaults: true,
	}

	for _, opt := range opts {
		if err := opt(l); err != nil {
			return err
		}
	}

	return l.load()
}

type viperLoader struct {
	viper       *viper.Viper
	configs     []configDef
	intoPtr     interface{}
	confFile    string
	confName    string
	useEnv      bool
	envPrefix   string
	useDefaults bool
}

func (l *viperLoader) load() error {
	v := l.viper

	keys, err := extractConfigDefs(l.intoPtr, l.useDefaults)
	if err != nil {
		return err
	}

	for _, cfg := range keys {
		v.SetDefault(cfg.Key, cfg.Default)
	}

	if l.useEnv {
		// for transforming app.host to app_host
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
		v.SetEnvPrefix(l.envPrefix)
		v.AutomaticEnv()
		for _, cfg := range keys {
			if err := v.BindEnv(cfg.Key); err != nil {
				return err
			}
		}
	}

	if l.confFile != "" {
		v.SetConfigFile(l.confFile)
		if err := v.ReadInConfig(); err != nil {
			return err
		}
	} else {
		if l.confName == "" {
			l.confName = "config"
		}
		v.AddConfigPath("./")
		v.AddConfigPath(getExecPath())
		v.SetConfigName(l.confName)
		_ = v.ReadInConfig()
	}

	return v.Unmarshal(l.intoPtr)
}

type configDef struct {
	Key     string      `json:"key"`
	Doc     string      `json:"doc"`
	Default interface{} `json:"default"`
}

func extractConfigDefs(structPtr interface{}, useDefaults bool) ([]configDef, error) {
	rv := reflect.ValueOf(structPtr)

	if err := ensureStructPtr(rv); err != nil {
		return nil, err
	}

	if useDefaults {
		defaults.SetDefaults(structPtr)
	}

	return readRecursive(deref(rv), "")
}

func readRecursive(rv reflect.Value, rootKey string) ([]configDef, error) {
	rt := rv.Type()

	var acc []configDef
	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		fv := deref(rv.Field(i))

		key := toCamelCase(ft.Name)
		if rootKey != "" {
			key = fmt.Sprintf("%s.%s", rootKey, key)
		}

		if fv.Kind() == reflect.Struct {
			nestedConfigs, err := readRecursive(fv, key)
			if err != nil {
				return nil, err
			}
			acc = append(acc, nestedConfigs...)
		} else {
			acc = append(acc, configDef{
				Key:     key,
				Doc:     ft.Tag.Get("doc"),
				Default: fv.Interface(),
			})
		}
	}

	return acc, nil
}

func toCamelCase(s string) string {
	var result string
	for i, r := range s {
		if i > 0 && (r >= 'A' && r < 'Z') {
			result += "_"
		}
		result += strings.ToLower(string(r))
	}
	return result
}

func deref(rv reflect.Value) reflect.Value {
	if rv.Kind() == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}
	return rv
}

func ensureStructPtr(value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("need a pointer to struct, not '%s'", value.Kind())
	} else {
		value = reflect.Indirect(value)
		if value.Kind() != reflect.Struct {
			return fmt.Errorf("need a pointer to struct, not pointer to '%s'", value.Kind())
		}
	}
	return nil
}

func getExecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(execPath)
}
