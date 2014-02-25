package parse

import (
    "testing"
    "fmt"
    "github.com/uberj/bindparse/scanner"
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

func setUp(input string) (*scanner.Scanner, *ZoneState) {
    return &scanner.Scanner{Source: input}, &ZoneState{Origin: "example.com."}
}

func TestParseSOA_1(t *testing.T) {
    test_zone := `example.com.  3600  IN   SOA   ns.example.org. sysadmins.example.org. (
        2014021804     ; Serial
            180     ; Refresh
            180     ; Retry
            1209600     ; Expire
            60     ; Minimum
    )`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    soa, err := ParseRecord(zs, s, rdd)
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
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    soa, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_3(t *testing.T) {
    test_zone := `example.com.  3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    soa, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_4(t *testing.T) {
    test_zone := `@ 3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    soa, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_5(t *testing.T) {
    test_zone := `@ IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    soa, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_6(t *testing.T) {
    test_zone := `IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    soa, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := soa.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}


func TestTrueName(t *testing.T) {
    zs := &ZoneState{Origin: "example.com."}
    AssertEqual(t, "foo.bar.example.com.", TrueName(zs, "foo.bar"))
    AssertEqual(t, "example.com.", TrueName(zs, "@"))
    AssertEqual(t, "bar.example.com.", TrueName(zs, "bar.example.com."))
}
