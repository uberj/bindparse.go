package parse

import (
    "testing"
    "github.com/uberj/bindparse/scanner"
    "fmt"
    "reflect"
)

func AssertEqual(t *testing.T, is interface{}, shouldbe interface{}) bool {
    if reflect.TypeOf(is) != reflect.TypeOf(shouldbe) {
        panic(fmt.Sprintf("Type '%s' != '%s'", reflect.TypeOf(is), reflect.TypeOf(shouldbe)))
    }
    if is == shouldbe {
        return true
    } else {
        t.Fatalf(fmt.Sprintf("'%s' != '%s'", is, shouldbe))
        return false
    }
}

func setUp(input string) (*scanner.Scanner, *ZoneState) {
    return &scanner.Scanner{Source: input}, &ZoneState{Origin: "example.com."}
}

