package go_russian_regions

import (
	"regexp"
)

type Region struct {
	Name     string
	Code     string
	reString string
	re       *regexp.Regexp
}

func (r *Region) Is(text string) bool {
	if r.re == nil {
		return false
	}
	return r.re.MatchString(r.canonicalName(text))
}

func (r *Region) canonicalName(text string) string { // уберем пробелы вокруг дефиса
	return regexp.MustCompile(`\s*[\-—]\s*`).ReplaceAllString(text, "-")
}
