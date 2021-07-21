package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinifyJSArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewMinifyJSArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with params", func(t *testing.T) {

		mapExpected := map[string]interface{}{
			"gzip":               false,
			"use_babel_polyfill": false,
			"keep_fn_name":       false,
			"keep_class_name":    false,
			"mangle":             false,
			"merge_vars":         false,
			"remove_console":     false,
			"remove_undefined":   false,
			"targets":            false,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewMinifyJSArgs()
		args.SetGzip(false)
		args.SetUseBabelPolyfill(false)
		args.SetKeepFnName(false)
		args.SetKeepClassName(false)
		args.SetMangle(false)
		args.SetMergeVars(false)
		args.SetRemoveConsole(false)
		args.SetRemoveUndefined(false)
		args.SetTargets(false)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
