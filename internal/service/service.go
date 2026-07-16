package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(s string) (string, error) {
	s = strings.TrimSpace(s)

	if s == "" {
		return "", errors.New("empty string")
	}

	if strings.Trim(s, ".- ") == "" {
		return morse.ToText(s), nil
	}

	return morse.ToMorse(s), nil
}
