package parse

import (
    "testing"
    "fmt"
    "github.com/uberj/bindparse/rdtype"
)

func TestParseNS_1(t *testing.T) {
    test_zone := `  IN  NS  ns1.example.com.`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.NS); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
    AssertEqual(t, rec.(rdtype.NS).Name, "example.com.")
    AssertEqual(t, rec.(rdtype.NS).TargetName, "ns1.example.com.")
}

func TestParseNS_2(t *testing.T) {
    test_zone := `@  IN  NS  ns1.example.com.`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.NS); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
    AssertEqual(t, rec.(rdtype.NS).Name, "example.com.")
    AssertEqual(t, rec.(rdtype.NS).TargetName, "ns1.example.com.")
}

func TestParseNS_3(t *testing.T) {
    test_zone := `example.com.  IN  NS  ns1.example.com.`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.NS); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
    AssertEqual(t, rec.(rdtype.NS).Name, "example.com.")
    AssertEqual(t, rec.(rdtype.NS).TargetName, "ns1.example.com.")
}

func TestParseNS_4(t *testing.T) {
    test_zone := `example.com. 999 IN  NS  ns1.example.com.`
    s, zs := setUp(test_zone)
    rdd := NewRDTypeDispatch()
    rec, err := ParseRecord(zs, s, rdd)
    if err != nil {
        t.Fatalf("%s\n", err)
    }
    if type_, ok := rec.(rdtype.NS); !ok {
        t.Fatalf("Expected type rdtype.SOA but unexpected type %T\n", type_)
    }
    AssertEqual(t, rec.(rdtype.NS).Name, "example.com.")
    AssertEqual(t, rec.(rdtype.NS).TargetName, "ns1.example.com.")
    fmt.Println(rec.(rdtype.NS).Ttl)
}
