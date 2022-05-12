package helpers

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Int returns a pointer to the int value passed in.
func Int(v int) *int {
	return &v
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the boolean value passed in.
func Bool(v bool) *bool {
	return &v
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

func IsOffline() bool {
	ok, err := strconv.ParseBool(os.Getenv("IS_OFFLINE"))
	if err != nil {
		return false
	}
	return ok
}

func IsTestEnv() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}
