package parse

import (
    "testing"
    //"fmt"
)

func TestDirective_TTL_1(t *testing.T) {
    test_dirctive := `$TTL  123`
    s, zs := setUp(test_dirctive)
    _, err := ParseDirective(zs, s)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    AssertEqual(t, zs.Ttl, int32(123))
}

func TestDirective_TTL_2(t *testing.T) {
    test_dirctive := `$TTL  123; asdfasdf asdf`
    s, zs := setUp(test_dirctive)
    _, err := ParseDirective(zs, s)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    AssertEqual(t, zs.Ttl, int32(123))
    AssertEqual(t, (<-s.Next()).Value, "")
}

func TestDirective_ORIGIN_1(t *testing.T) {
    test_dirctive := `$ORIGIN  foobar.com.;asdf asdf`
    s, zs := setUp(test_dirctive)
    _, err := ParseDirective(zs, s)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    AssertEqual(t, zs.Origin, "foobar.com.")
}

func TestDirective_ORIGIN_2(t *testing.T) {
    test_dirctive := `$ORIGIN  foobar.com`
    s, zs := setUp(test_dirctive)
    _, err := ParseDirective(zs, s)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    AssertEqual(t, zs.Origin, "foobar.com.example.com.")
}

