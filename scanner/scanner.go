package scanner

import  "regexp"

type Scanner struct {
    Source string
    curidx int
}

type Token struct {
    End bool
    Value string
}

func validIndex(i int, s string) bool {
    if i < 0 {
        return false
    } else if i >= len(s) {
        return false
    } else{
        return true
    }
}

func (s *Scanner) ValidIndex(curidx int) bool {
    return validIndex(curidx, s.Source)
}

func (s *Scanner) consume(curidx *int, r *regexp.Regexp) bool {
    return s.consumen(curidx, r, 1)
}

func (s *Scanner) consumen(curidx *int, r *regexp.Regexp, n int) bool {
    // Consume as much r as possible
    // If at least one r was consumed return true
    ret := false
    for {
        if !s.ValidIndex(*curidx) {
            break
        }
        if r.MatchString(string(s.Source[*curidx])) {
            *curidx = *curidx + n
            ret = true
        } else {
            break
        }
    }
    return ret
}

func (s *Scanner) Peekn(n int) chan Token {
    fake_curidx := s.curidx
    return s.next(`[\n;]`, &fake_curidx, n)
}

func (s *Scanner) PeekUntil(pattern string) chan Token {
    fake_curidx := s.curidx
    return s.next(pattern, &fake_curidx, 1)
}

func (s *Scanner) Peek() chan Token {
    fake_curidx := s.curidx
    return s.next(`[\n;]`, &fake_curidx, 1)
}

func (s *Scanner) Next() chan Token {
    return s.next(`[\n;]`, &s.curidx, 1)
}

func (s *Scanner) NextUntil(pattern string) chan Token {
    return s.next(pattern, &s.curidx, 1)
}

func (s *Scanner) next(until string, curidx *int, n int) chan Token {
    yield := make (chan Token, n)
    go func () {
        ws := regexp.MustCompile(`[^\S\n]`)
        delim := regexp.MustCompile(until)
        for i := 0; i < n; i++ {
            if !s.ValidIndex(*curidx) {
                yield <- Token{true, ""}
                continue
            }
            s.consume(curidx, ws)
            start := *curidx
            end := *curidx
            // invariant: end is valid. curidx might not be valid
            for {
                if !s.ValidIndex(*curidx) {
                    break
                }
                if delim.MatchString(string(s.Source[*curidx])) {
                    break
                }
                if ws.MatchString(string(s.Source[*curidx])){
                    break
                }
                *curidx++
                end = *curidx
            }
            if start == end {
                *curidx++
                yield <- Token{Value:s.Source[start:end + 1]}
            } else {
                yield <- Token{Value:s.Source[start:end]}
            }
        }
    } ()
    return yield
}
