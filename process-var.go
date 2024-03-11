package gox

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

var processVarPrefix = ""

func SetProcessVarEnvPrefix(str string) {
	processVarPrefix = str
}

func applyProcessVarReflect(valueOf reflect.Value, str string) {
	if len(str) > 0 {
		switch valueOf.Type().Kind() {
		case reflect.Bool:
			tmp, _ := strconv.ParseBool(str)
			valueOf.SetBool(tmp)
		case reflect.Int, reflect.Int64:
			if tmp, err := strconv.ParseInt(str, 10, 64); err == nil {
				valueOf.SetInt(tmp)
			}
		case reflect.Int32:
			if tmp, err := strconv.ParseInt(str, 10, 32); err == nil {
				valueOf.SetInt(tmp)
			}
		case reflect.Int16:
			if tmp, err := strconv.ParseInt(str, 10, 16); err == nil {
				valueOf.SetInt(tmp)
			}
		case reflect.Int8:
			if tmp, err := strconv.ParseInt(str, 10, 8); err == nil {
				valueOf.SetInt(tmp)
			}
		case reflect.Uint, reflect.Uint64:
			if tmp, err := strconv.ParseUint(str, 10, 64); err == nil {
				valueOf.SetUint(tmp)
			}
		case reflect.Uint32:
			if tmp, err := strconv.ParseUint(str, 10, 32); err == nil {
				valueOf.SetUint(tmp)
			}
		case reflect.Uint16:
			if tmp, err := strconv.ParseUint(str, 10, 16); err == nil {
				valueOf.SetUint(tmp)
			}
		case reflect.Uint8:
			if tmp, err := strconv.ParseUint(str, 10, 8); err == nil {
				valueOf.SetUint(tmp)
			}
		case reflect.Float64:
			if tmp, err := strconv.ParseFloat(str, 64); err == nil {
				valueOf.SetFloat(tmp)
			}
		case reflect.Float32:
			if tmp, err := strconv.ParseFloat(str, 32); err == nil {
				valueOf.SetFloat(tmp)
			}
		case reflect.String:
			valueOf.SetString(str)
		}
	}
}

func parseProcessVarValue[Type Primitive](val *Type, str string) {
	valueOf := reflect.ValueOf(val).Elem()
	applyProcessVarReflect(valueOf, str)
}

func applyProcessVarEnv[Type Primitive](val *Type, name string) {
	if env := os.Getenv(processVarPrefix + "_" + strings.ToUpper(name)); env != "" {
		parseProcessVarValue[Type](val, env)
	}
}

func applyProcessVarCmd[Type Primitive](val *Type, name string) {
	rawArgs := os.Args[1:]
	if len(rawArgs) > 0 {
		args := TransformMapKeys[string, string](SplitStringsToMap(rawArgs, "="), func(s string) string {
			str, _ := strings.CutPrefix(s, "--")
			return strings.ToLower(str)
		})

		if arg, ok := args[name]; ok {
			// Empty strings are treated as flags
			if arg == "" {
				arg = "true"
			}
			parseProcessVarValue[Type](val, arg)
		}
	}
}

// ProcessVar is a self-initializing variable which looks for it's value in
// multiple places at the time of instantiation. In other words, it's for those
// cases where you want to have over-writable variables on startup such as ENV,
// arguments, default, etc.
//
// Order of operations for finding the value is:
// - Default Value
// - Environment variable
// - Command line argument
func ProcessVar[Type Primitive](name string, defValue Type) Type {
	val := defValue

	applyProcessVarEnv[Type](&val, name)
	applyProcessVarCmd[Type](&val, name)

	return val
}

// StructProcessVar is a self-initializing variable which looks for it's value
// in multiple places at the time of instantiation. This version targets a struct
// and it's fields. If the struct contains exported fields it will search both the
// environment and command line arguments for a value. It finds these by prefixing
// both the env, and cli lookups with the given prefix to prevent global crowding.
//
// Order of operations for finding the value is:
// - Default Value
// - Environment variable
// - Command line argument
func StructProcessVar[Type any](prefix string, defValue Type) Type {
	val := defValue

	envPrefix := processVarPrefix + "_" + strings.ToUpper(prefix) + "_"

	valueOf := reflect.Indirect(reflect.ValueOf(&val))
	if valueOf.Type().Kind() != reflect.Struct {
		panic("cannot parse StructProcessVar on a non-struct type")
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)
		if field.IsExported() {
			tag, ok := field.Tag.Lookup("json")
			if !ok {
				tag = strings.ToLower(field.Name)
			}

			if IsPrimitive(field.Type) {
				if env := os.Getenv(envPrefix + strings.ToUpper(tag)); env != "" {
					applyProcessVarReflect(valueOf.Field(i), tag)
				}

				rawArgs := os.Args[1:]
				if len(rawArgs) > 0 {
					args := TransformMapKeys[string, string](SplitStringsToMap(rawArgs, "="), func(s string) string {
						str, _ := strings.CutPrefix(s, "--")
						return strings.ToLower(str)
					})

					name := strings.ToLower(prefix + "_" + tag)
					if arg, ok := args[name]; ok {
						// Empty strings are treated as flags
						if arg == "" {
							arg = "true"
						}
						applyProcessVarReflect(valueOf.Field(i), arg)
					}
				}
			}
		}
	}

	return val
}
