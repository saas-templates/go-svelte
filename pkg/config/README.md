# Config

A convenience wrapper around `viper` that provides easy-to-use struct-based config loading.

## Example

### Load Directly

```golang
package main

func main() {
	var cfg Config
	opts := []config.Option{
		config.WithEnv(),
	}
	if err := config.Load(&cfg, opts...); err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}

type Config struct {
	Addr   string `default:":8080"`
	StatsD struct {
		Host string `default:"localhost"`
		Port int    `default:"8125"`
	}
}
```

