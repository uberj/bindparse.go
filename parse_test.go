package main

import (
    "testing"
    "fmt"
    "github.com/uberj/bindparse/scanner"
    "github.com/uberj/bindparse/parse"
    "github.com/uberj/bindparse/rdtype"
)

func AssertEqual(t *testing.T, is string, shouldbe string) bool {
    if is == shouldbe {
        return true
    } else {
        t.Fatalf(fmt.Sprintf("%s != %s", is, shouldbe))
        return false
    }
}

func TestParseSOA_1(t *testing.T) {
    test_zone := `example.com.  3600  IN   SOA   ns.example.org. sysadmins.example.org. (
        2014021804     ; Serial
            180     ; Refresh
            180     ; Retry
            1209600     ; Expire
            60     ; Minimum
    )`
    s := &scanner.Scanner{Source: test_zone}
    zs := &parse.ZoneState{Origin: "example.com."}
    rdd := parse.NewRDTypeDispatch()
    soa, err := parse.ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_2(t *testing.T) {
    test_zone := `example.com.  3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     ; Serial
            180     ; Refresh
            180     ; Retry
            1209600     ; Expire
            60     ; Minimum
    )`
    s := &scanner.Scanner{Source: string(test_zone)}
    zs := &parse.ZoneState{Origin: "example.com."}
    rdd := parse.NewRDTypeDispatch()
    soa, err := parse.ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_3(t *testing.T) {
    test_zone := `example.com.  3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s := &scanner.Scanner{Source: string(test_zone)}
    zs := &parse.ZoneState{Origin: "example.com."}
    rdd := parse.NewRDTypeDispatch()
    soa, err := parse.ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_4(t *testing.T) {
    test_zone := `@ 3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s := &scanner.Scanner{Source: string(test_zone)}
    zs := &parse.ZoneState{Origin: "example.com."}
    rdd := parse.NewRDTypeDispatch()
    soa, err := parse.ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_5(t *testing.T) {
    test_zone := `@ IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s := &scanner.Scanner{Source: string(test_zone)}
    zs := &parse.ZoneState{Origin: "example.com."}
    rdd := parse.NewRDTypeDispatch()
    soa, err := parse.ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_6(t *testing.T) {
    test_zone := `IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s := &scanner.Scanner{Source: string(test_zone)}
    zs := &parse.ZoneState{Origin: "example.com."}
    rdd := parse.NewRDTypeDispatch()
    soa, err := parse.ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}


func TestTrueName(t *testing.T) {
    zs := &parse.ZoneState{Origin: "example.com."}
    AssertEqual(t, "foo.bar.example.com.", parse.TrueName(zs, "foo.bar"))
    AssertEqual(t, "example.com.", parse.TrueName(zs, "@"))
    AssertEqual(t, "bar.example.com.", parse.TrueName(zs, "bar.example.com."))
}
