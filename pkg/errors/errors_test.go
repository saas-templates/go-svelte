package errors_test

import (
	goerrors "errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/saas-templates/go-svelte/pkg/errors"
)

func TestError_Error(t *testing.T) {
	t.Parallel()

	table := []struct {
		title string
		err   errors.Error
		want  string
	}{
		{
			title: "WithoutCause",
			err:   errors.ErrInvalid,
			want:  "request is not valid",
		},
		{
			title: "WithCause",
			err:   errors.ErrInvalid.WithMsgf("").WithCausef("foo"),
			want:  "bad_request: foo",
		},
	}

	for _, tt := range table {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.err.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestError_Is(t *testing.T) {
	t.Parallel()

	table := []struct {
		title string
		err   errors.Error
		other error
		want  bool
	}{
		{
			title: "WithDifferentCode",
			err:   errors.ErrInternal,
			other: errors.ErrInvalid,
			want:  false,
		},
		{
			title: "NonError",
			err:   errors.ErrInternal,
			other: goerrors.New("foo"),
			want:  true,
		},
		{
			title: "WithSameCode",
			err:   errors.ErrInvalid.WithCausef("cause 1"),
			other: errors.ErrInvalid.WithCausef("cause 2"),
			want:  true,
		},
	}

	for _, tt := range table {
		t.Run(tt.title, func(t *testing.T) {
			got := goerrors.Is(tt.err, tt.other)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestError_WithCausef(t *testing.T) {
	t.Parallel()

	table := []struct {
		title string
		err   errors.Error
		want  errors.Error
	}{
		{
			title: "WithCauseString",
			err:   errors.ErrInvalid.WithCausef("foo"),
			want: errors.Error{
				Code:    "bad_request",
				Message: "Request is not valid",
				Cause:   "foo",
			},
		},
		{
			title: "WithCauseFormatted",
			err:   errors.ErrConflict.WithCausef("hello %s", "world"),
			want: errors.Error{
				Code:    "conflict",
				Message: "An entity with conflicting identifier exists",
				Cause:   "hello world",
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.err)
		})
	}
}

func TestError_WithMsgf(t *testing.T) {
	t.Parallel()

	table := []struct {
		title string
		err   errors.Error
		want  errors.Error
	}{
		{
			title: "WithCauseString",
			err:   errors.ErrInvalid.WithMsgf("foo"),
			want: errors.Error{
				Code:    "bad_request",
				Message: "foo",
			},
		},
		{
			title: "WithCauseFormatted",
			err:   errors.ErrInvalid.WithMsgf("hello %s", "world"),
			want: errors.Error{
				Code:    "bad_request",
				Message: "hello world",
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.err)
		})
	}
}

func Test_Errorf(t *testing.T) {
	e := errors.Errorf("failed: %d", 100)
	assert.Error(t, e)
	assert.EqualError(t, e, "failed: 100")
}
