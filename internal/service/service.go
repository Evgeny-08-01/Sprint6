package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)
func TexttoMorseAndRevers(str string) string {
	s1 := morse.ToText(str)
	s2 := morse.ToMorse(s1)
	if s2 == str {
    return morse.ToText(str)
	} else {	
		return morse.ToMorse(str)
	}
}
