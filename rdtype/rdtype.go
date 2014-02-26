package rdtype

import "fmt"

type Rdtyper interface {
    String() string
}

type SOA struct {
    Name string
    Ttl int32
    Primary string
    Contact string
    Serial int32
    Retry int32
    Refresh int32
    Expire int32
    Minimum int32
}

func (soa SOA) String() string {
    return fmt.Sprintf("%s %d SOA IN %s %s (%d %d %d %d %d)",
        soa.Name,
        soa.Ttl,
        soa.Primary,
        soa.Contact,
        soa.Serial,
        soa.Retry,
        soa.Refresh,
        soa.Expire,
        soa.Minimum)
}

type NS struct {
    Name string
    Ttl int32
    TargetName string
}

func (soa NS) String() string {
    return fmt.Sprintf("%s %d NS IN %s",
        soa.Name,
        soa.Ttl,
        soa.TargetName)
}
