package scanner

import  "regexp"
//import  "fmt"

type Scanner struct {
    Source string
    CurIdx int
}

type Token struct {
    End bool
    Token string
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

func (s *Scanner) ValidIndex(CurIdx int) bool {
    return validIndex(CurIdx, s.Source)
}

func (s *Scanner) consume(CurIdx *int, r *regexp.Regexp) bool {
    return s.consumen(CurIdx, r, 1)
}

func (s *Scanner) consumen(CurIdx *int, r *regexp.Regexp, n int) bool {
    // Consume as much r as possible
    // If at least one r was consumed return true
    ret := false
    for {
        if !s.ValidIndex(*CurIdx) {
            break
        }
        if r.MatchString(string(s.Source[*CurIdx])) {
            *CurIdx = *CurIdx + n
            ret = true
        } else {
            break
        }
    }
    return ret
}

func (s *Scanner) Peekn(n int) chan Token {
    fake_CurIdx := s.CurIdx
    return s.next(`\n`, &fake_CurIdx, n)
}

func (s *Scanner) PeekUntil(pattern string) chan Token {
    fake_CurIdx := s.CurIdx
    return s.next(pattern, &fake_CurIdx, 1)
}

func (s *Scanner) Peek() chan Token {
    fake_CurIdx := s.CurIdx
    return s.next(`\n`, &fake_CurIdx, 1)
}

func (s *Scanner) Next() chan Token {
    return s.next(`\n`, &s.CurIdx, 1)
}

func (s *Scanner) NextUntil(pattern string) chan Token {
    return s.next(pattern, &s.CurIdx, 1)
}

func (s *Scanner) next(until string, CurIdx *int, n int) chan Token {
    yield := make (chan Token, n)
    go func () {
        ws := regexp.MustCompile(`[^\S\n]`)
        delim := regexp.MustCompile(until)
        for i := 0; i < n; i++ {
            if !s.ValidIndex(*CurIdx) {
                yield <- Token{true, ""}
                continue
            }
            s.consume(CurIdx, ws)
            start := *CurIdx
            end := *CurIdx
            // invariant: end is valid. curidx might not be valid
            for {
                if !s.ValidIndex(*CurIdx) {
                    break
                }
                if delim.MatchString(string(s.Source[*CurIdx])) {
                    break
                }
                if ws.MatchString(string(s.Source[*CurIdx])){
                    break
                }
                *CurIdx++
                end = *CurIdx
            }
            if start == end {
                *CurIdx++
                yield <- Token{Token:s.Source[start:end + 1]}
            } else {
                yield <- Token{Token:s.Source[start:end]}
            }
        }
    } ()
    return yield
}
