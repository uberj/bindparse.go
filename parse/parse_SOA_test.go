package parse

import (
    "testing"
    "github.com/uberj/bindparse/rdtype"
)

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
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.SOA); !ok {
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
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_3(t *testing.T) {
    test_zone := `example.com.  3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_4(t *testing.T) {
    test_zone := `@ 3600  IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_5(t *testing.T) {
    test_zone := `@ IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}

func TestParseSOA_6(t *testing.T) {
    test_zone := `IN   SOA   ns.example.org. sysadmins.example.org.(2014021804     180 180 1209600 60)`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.SOA); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
}


func TestTrueName(t *testing.T) {
    zs := &ZoneState{Origin: "example.com."}
    AssertEqual(t, "foo.bar.example.com.", trueName(zs, "foo.bar"))
    AssertEqual(t, "example.com.", trueName(zs, "@"))
    AssertEqual(t, "bar.example.com.", trueName(zs, "bar.example.com."))
}
