package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Texttomorse_and_revers(str string) string {

	s1 := morse.ToMorse(str)
	s2 := morse.ToText(s1)
	if s2 == str {
		return s1
	} else {
		return morse.ToText(str)
	}
}
