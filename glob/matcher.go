package glob

import (
	"strings"

	"github.com/gobwas/glob"
)

type Matcher struct {
	matchers []m
}

type m struct {
	g glob.Glob
	b bool
}

func PreprocessGlobSpec(globSpec string) *Matcher {
	globSpecs := strings.Split(globSpec, ",")

	matchers := make([]m, len(globSpecs))
	for k, v := range globSpecs {
		v := strings.TrimSpace(v)
		b := strings.HasPrefix(v, "!")
		if b {
			v = v[1:]
		}
		g, _ := glob.Compile(v)
		matchers[k] = m{g, !b}
	}

	return &Matcher{matchers}
}

func (m *Matcher) Matches(path string) bool {
	for _, v := range m.matchers {
		if v.g.Match(path) != v.b {
			return false
		}
	}
	return true
}
