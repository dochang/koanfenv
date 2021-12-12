// Package koanfenv provides callbacks which convert environment variables to
// koanf keys.  These callbacks are used for
// `github.com/knadh/koanf/providers/env#Provider`.
package koanfenv

import "reflect"

func fillMap(keyMap *map[string]string, i interface{}) {
	v := reflect.ValueOf(i)
	v = reflect.Indirect(v)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		envTag := field.Tag.Get("env")
		if envTag == "" || envTag == "-" {
			continue
		}
		koanfKey := field.Tag.Get("koanf")
		if koanfKey == "-" {
			continue
		}
		if koanfKey == "" {
			koanfKey = field.Name
		}
		(*keyMap)[envTag] = koanfKey
	}
}

// ByStruct translates environment variables to koanf keys by struct field tags.
//
// First, all unexported fields are ignored.
//
// If key "env" is in an exported field tag and its value is not empty or "-",
// the environment variable, whose name is the value of "env", will be
// translated.  Otherwise, the field will be ignored.
//
// If key "koanf" is in an exported field tag and its value is not empty, the
// value will be used as the koanf key.
//
// If key "koanf" is in an exported field tag and its value is empty, or
// "koanf" is not in the tag, the field name will be used as the koanf key.
//
// If key "koanf" is in an exported field tag and its value is "-", the field
// will be ignored.
//
// Here are some examples:
//
//   // Translate "VAR" to "key".
//   Field string `koanf:"key" env:"VAR"`
//
//   // Translate "VAR" to "Field"
//   Field string `koanf:"" env:"VAR"`
//   Field string `env:"VAR"`
//
//   // Ignore these fields.
//   Field string `koanf:"-" env:"VAR"`
//   Field string `koanf:"key" env:"-"`
//   Field string `koanf:"key" env:""`
//   Field string `koanf:"key"`
//   Field string ``
//   Field string
func ByStruct(i interface{}) func(string) string {
	keyMap := map[string]string{}

	fillMap(&keyMap, i)

	return ByMap(keyMap)
}

// ByMap translates environment variables to koanf keys by a map.
func ByMap(keyMap map[string]string) func(string) string {
	return func(tag string) string {
		return keyMap[tag]
	}
}
