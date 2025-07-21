package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Morze(s string) string {
	if Typestr(s) {
		return morse.ToText(s)
	} else {
		return morse.ToMorse(s)
	}

}

func Typestr(s string) bool {
	s = strings.Trim(s, " ")
	for _, v := range s {
		if v != '.' && v != '-' && v != ' ' {
			return false
		}
	}
	return true
}
