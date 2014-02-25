package parse

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/uberj/bindparse/scanner"
    "github.com/uberj/bindparse/rdtype"
)


type ZoneState struct {
    /*
     * Tracks the state of a zone during parsing
     */
    Origin string
    Ttl int32
}

var RDCLASS = map[string] int {
    "IN": 1,
}

type rdTypeDispatch struct {
    RDType map[string]func(
        name string,
        ttl int32,
        rdclass string,
        zstate *ZoneState,
        s *scanner.Scanner) (rdtype.Rdtyper, error)
}

func NewRDTypeDispatch() *rdTypeDispatch {
    return &rdTypeDispatch{
        map[string]func(
            name string,
            ttl int32,
            rdclass string,
            zstate *ZoneState,
            s *scanner.Scanner) (rdtype.Rdtyper, error) {
            "SOA": ParseSOA,
        },
    }
}

// RDCLASS

func TrueName(zstate *ZoneState, name string) string {
    // If name is @ return origin
    // If name doesn't end with root append origin
    if name == "@" {
        return zstate.Origin
    }

    if !strings.HasSuffix(name, ".") {
        new_name := []string{name, (*zstate).Origin}
        return strings.Join(new_name, ".")
    }

    return name
}

func goodToken(t *scanner.Token) error {
    if t.End {
        return fmt.Errorf("Unexpected end of input\n")
    }
    return nil
}

func ParseRecord(zstate *ZoneState, s *scanner.Scanner, dispatch *rdTypeDispatch) (rdtype.Rdtyper, error) {
    // Parse Record beginning of record then dispatch the record's specific parse function

    // Three Cases
    // name ttl  class type         -- case 2
    // 0    1    2     3
    //      name class type         -- case 1
    //      0    1     2
    //      ttl  class type         -- case 1
    //      0    1     2
    //           class type         -- case 0
    //           0     1
    // If 0 & 1 are class and type  -- case 0
    //  $ORIGIN and $TTL
    // If 1 & 2 are class and type  -- case 1
    //  if ttl is int32
    //      $ORIGIN and ttl
    //  else
    //      name and ttl
    // If 2 & 3 are class and type  -- case 2
    //  name and ttl

    peeker := s.Peekn(4)
    tokens := [...]scanner.Token{
        <-peeker,
        <-peeker,
        <-peeker,
        <-peeker,
    }

    var ttl int32
    var name, rdclass, rdtype_ string
    var valid_class bool

    for i := 0; i < 3; i++ {
        cur := tokens[i]
        next := tokens[i + 1]

        _, valid_class = RDCLASS[strings.ToUpper(cur.Value)]
        _, valid_rdtype := dispatch.RDType[strings.ToUpper(next.Value)]

        rdclass = cur.Value
        rdtype_ = next.Value

        if valid_class && valid_rdtype {
            switch i {
            case 0:
                <-s.Next()
                <-s.Next()
                name = TrueName(zstate, "@")
                ttl = zstate.Ttl
                goto DISPATCH
            case 1:
                <-s.Next()
                <-s.Next()
                <-s.Next()
                // ttl class type
                // or
                // name class type
                ttl_, is_ttl := strconv.ParseInt(tokens[0].Value, 0, 64)
                if is_ttl != nil {
                    name = TrueName(zstate, tokens[0].Value)
                    ttl = zstate.Ttl
                } else {
                    name = TrueName(zstate, "@")
                    ttl = int32(ttl_)
                }
                goto DISPATCH
            case 2:
                <-s.Next()
                <-s.Next()
                <-s.Next()
                <-s.Next()
                name = TrueName(zstate, tokens[0].Value)
                ttl_, err := strconv.ParseInt(tokens[1].Value, 0, 64)
                if err != nil {
                    return nil, err
                }
                ttl = int32(ttl_)
                goto DISPATCH
            }
        }
    }

DISPATCH:
    return dispatch.RDType[rdtype_](name, ttl, rdclass, zstate, s)
}

func ParseComment(s *scanner.Scanner) (bool, string, error) {
    // ParseComment: Consume a comment if it exists. Stop leaving \n next
    next := <-s.PeekUntil(`[\n;]`);
    if next.Value != ";" {
        return false, "", nil
    }
    for {
        peek := <-s.Peek()
        if peek.Value == "\n" {
            break
        }
        <-s.Next()
    }
    return true, "", nil
}

func ParseNewLines(s *scanner.Scanner) (bool, error) {
    // ParseNewLines: Consume newlines
    ret := false
    for {
        peek := <-s.Peek()
        //if err := goodToken(&peek); err != nil {return false, err}
        if peek.Value == "\n" {
            ret = true
            <-s.Next()
            continue
        }
        break
    }
    return ret, nil
}

func ClearCommentNewLine(s *scanner.Scanner) error {
    for {
        parsed_comment, _, err := ParseComment(s)
        if err != nil {return err}
        parsed_newline, err := ParseNewLines(s)
        if err != nil {return err}

        if !(parsed_newline && parsed_comment) {
            break
        }
    }
    return nil
}


func requireInt32(p string) (int32, error){
    i, err := strconv.ParseInt(p, 0, 32)
    if err != nil {
        return 0, fmt.Errorf("Expected int32 but found %s\n", p)
    }
    return int32(i), nil
}

func parseSOAInt32(s *scanner.Scanner, cur scanner.Token) (int32, error) {
    cur = <-s.Next();
    if err := goodToken(&cur); err != nil {return 0, err}
    i, err := requireInt32(cur.Value)
    if err != nil {return 0, err}
    if err := ClearCommentNewLine(s); err != nil {return 0, err}
    return i, nil
}

func ParseSOA(name string, ttl int32, rdclass string, zstate *ZoneState, s *scanner.Scanner) (rdtype.Rdtyper, error) {
    // Assume scanner is currently on rdtype
    // We expect to see:
    // primary contact (serial, refresh, retry, exprire, minimum)
    // Parse out comments and ignore them
    //fmt.Printf("ParseSOA rdclass='%s' rdtype='%s'\n", rdclass, "SOA")
    cur := <-s.Next();
    if err := goodToken(&cur); err != nil {return nil, err}
    primary := TrueName(zstate, cur.Value)
    //fmt.Printf("primary=%s\n", primary)

    cur = <-s.NextUntil(`[\n\(]`);
    if err := goodToken(&cur); err != nil {return nil, err}
    contact := TrueName(zstate, cur.Value)
    //fmt.Printf("contact=%s\n", contact)

    if err := ClearCommentNewLine(s); err != nil {return nil, err}

    //fmt.Printf("peek='%s'\n", (<-s.Peek()).Value)
    cur = <-s.NextUntil(`[\(\n]`);
    //fmt.Printf("lparen='%s'\n", cur.Value)
    //fmt.Printf("peek='%s'\n", (<-s.Peek()).Value)
    if err := goodToken(&cur); err != nil {return nil, err}
    if cur.Value != "(" {
        return nil, fmt.Errorf(
            "Expected '(' in SOA definition but instead found '%s'\n", cur.Value)
    }

    if err := ClearCommentNewLine(s); err != nil {return nil, err}

    serial, err := parseSOAInt32(s, cur)
    if err != nil {return nil, err}
    //fmt.Printf("serial=%d\n", serial)

    refresh, err := parseSOAInt32(s, cur)
    if err != nil {return nil, err}
    //fmt.Printf("refresh=%d\n", refresh)

    retry, err := parseSOAInt32(s, cur)
    if err != nil {return nil, err}
    //fmt.Printf("retry=%d\n", retry)

    expire, err := parseSOAInt32(s, cur)
    if err != nil {return nil, err}
    //fmt.Printf("expire=%d\n", expire)

    cur = <-s.NextUntil(`[\)\n]`);
    if err := goodToken(&cur); err != nil {return nil, err}
    minimum, err := requireInt32(cur.Value)
    if err != nil {return nil, err}
    //fmt.Printf("minimum=%d\n", minimum)

    //fmt.Printf("-------------\n")
    if err := ClearCommentNewLine(s); err != nil {return nil, err}

    cur = <-s.Next();
    if cur.Value != ")" {
        return nil, fmt.Errorf(
            "Expected ')' in SOA definition but instead found '%s'\n", cur.Value)
    }
    soa := rdtype.SOA{name, ttl, primary, contact, serial, retry, refresh, expire, minimum}
    return rdtype.Rdtyper(soa), nil
}
