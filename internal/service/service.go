package service

import (
//	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Texttomorse_and_revers(str string) string {
	s1 := morse.ToText(str)
	s2 := morse.ToMorse(s1)
	if s2 == str {
    return morse.ToText(str)
	} else {	
		return morse.ToMorse(str)
	}
}
