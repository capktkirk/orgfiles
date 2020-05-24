package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	var abrev string
	s = strings.Replace(s, "-", " ", -1)
	reg, _ := regexp.Compile("[^a-zA-Z ]*")
	p := reg.ReplaceAllString(s, "")
	test := strings.Fields(p)
	for i := 0; i < len(test); i++ {
		addon := strings.ToUpper(string([]rune(string(test[i][0]))))
		abrev += string(addon)
	}
	return abrev
}
