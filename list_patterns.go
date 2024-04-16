package locales

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type (
	JoinType      uint8
	ListJoinToken uint8
)

var (
	DefaultAndListJoinPattern = NewListPattern("{0} and {1}", "{0}, {1}", "{0}, {1}", "{0} and {1}")
	DefaultOrListJoinPattern  = NewListPattern("{0} or {1}", "{0}, {1}", "{0}, {1}", "{0} or {1}")
)

const (
	And JoinType = iota
	Or
)

const (
	ListJoinA ListJoinToken = iota
	ListJoinSep
	ListJoinB
)

func (t ListJoinToken) String() string {
	switch t {
	case ListJoinA:
		return "A"
	case ListJoinB:
		return "B"
	case ListJoinSep:
		return "SEP"
	}
	return ""
}

type ListPatternSpec struct {
	Pattern string
	Sep     string
	Tokens  [3]ListJoinToken
}

func (s *ListPatternSpec) String() string {
	return fmt.Sprintf("%q [sep=%q, tokens=%v]", s.Pattern, s.Sep, s.Tokens)
}

func NewListPatternSpec(pattern string) *ListPatternSpec {
	if pattern == "" {
		return nil
	}
	l := len(pattern)
	s := &ListPatternSpec{Pattern: pattern}
	ai, bi := strings.Index(pattern, "{0}"), strings.Index(pattern, "{1}")

	if ai == 0 {
		s.Tokens[0] = ListJoinA
		if bi == l-3 {
			s.Tokens[1] = ListJoinSep
			s.Tokens[2] = ListJoinB
		} else {
			s.Tokens[1] = ListJoinB
			s.Tokens[2] = ListJoinSep
		}
	} else if bi == 0 {
		s.Tokens[0] = ListJoinB
		if ai == l-3 {
			s.Tokens[1] = ListJoinSep
			s.Tokens[2] = ListJoinA
		} else {
			s.Tokens[1] = ListJoinA
			s.Tokens[2] = ListJoinSep
		}
	} else {
		panic("bad pattern: " + pattern)
	}

	s.Sep = pattern[3 : l-3]
	return s
}

func (s *ListPatternSpec) Format(itens ...string) string {
	var w bytes.Buffer
	if l := len(itens); l > 0 {
		w.Write([]byte(itens[0]))
		for i := 1; i < l; i++ {
			w.Write([]byte(s.Sep))
			w.Write([]byte(itens[i]))
		}
	}
	return w.String()
}

type ListPattern struct {
	Two,
	Start,
	Middle,
	End *ListPatternSpec
}

func NewListPattern(two, start, middle, end string) *ListPattern {
	return &ListPattern{
		NewListPatternSpec(two),
		NewListPatternSpec(start),
		NewListPatternSpec(middle),
		NewListPatternSpec(end),
	}
}

func NewListPatternFromSlice(def []string) *ListPattern {
	if len(def) == 0 {
		return nil
	}
	return NewListPattern(def[0], def[1], def[2], def[3])
}

func (p *ListPattern) ToSlice() []string {
	if p == nil {
		return nil
	}
	ret := make([]string, 4)
	if p.Two != nil {
		ret[0] = p.Two.Pattern
	}
	if p.Start != nil {
		ret[1] = p.Start.Pattern
	}
	if p.Middle != nil {
		ret[2] = p.Middle.Pattern
	}
	if p.End != nil {
		ret[3] = p.End.Pattern
	}
	return ret
}

func (p *ListPattern) FormatTo(w io.Writer, itens ...string) {
	if p == nil {
		w.Write([]byte(strings.Join(itens, ", ")))
		return
	}
	switch len(itens) {
	case 0:
		return
	case 1:
		w.Write([]byte(itens[0]))
	case 2:
		w.Write([]byte(p.Two.Format(itens...)))
	case 3:
		w.Write([]byte(p.Start.Format(itens[0], p.End.Format(itens[len(itens)-2:]...))))
	default:
		end := p.End.Format(itens[len(itens)-2:]...)
		middle := p.Middle.Format(append(itens[1:len(itens)-2], end)...)
		w.Write([]byte(p.Start.Format(itens[0], middle)))
	}
}

type ListPatterns struct {
	AndP, OrP *ListPattern
}

func NewListPatternsFromSlice(s [][]string) *ListPatterns {
	if len(s) == 0 {
		return nil
	}
	p := &ListPatterns{
		NewListPatternFromSlice(s[0]),
		NewListPatternFromSlice(s[1]),
	}
	if p.AndP == nil {
		p.AndP = DefaultAndListJoinPattern
	}
	if p.OrP == nil {
		p.OrP = DefaultOrListJoinPattern
	}
	return p
}

func (f *ListPatterns) ToSlice() [][]string {
	if f == nil {
		return nil
	}
	return [][]string{f.AndP.ToSlice(), f.OrP.ToSlice()}
}

func (p *ListPatterns) And(w io.Writer, itens ...string) {
	p.Do(w, And, itens...)
}

func (p *ListPatterns) AndS(itens ...string) string {
	return p.DoS(And, itens...)
}

func (p *ListPatterns) Or(w io.Writer, itens ...string) {
	p.Do(w, Or, itens...)
}

func (p *ListPatterns) OrS(itens ...string) string {
	return p.DoS(Or, itens...)
}

func (p *ListPatterns) Do(w io.Writer, typ JoinType, itens ...string) {
	if p == nil {
		w.Write([]byte(strings.Join(itens, ", ")))
		return
	}
	switch typ {
	case Or:
		p.OrP.FormatTo(w, itens...)
	default:
		p.AndP.FormatTo(w, itens...)
	}
}

func (p *ListPatterns) DoS(typ JoinType, itens ...string) string {
	var w bytes.Buffer
	p.Do(&w, typ, itens...)
	return w.String()
}
