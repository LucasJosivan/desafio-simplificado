package domains

import (
	"errors"
	"strings"
)

var (
	ErrRespInvalid = errors.New("você deve escolher 'cara' ou 'coroa'")
)

type Game struct {
	Wins  int
	Loses int
}

func Check(resp string) error {
	if resp != "s" && resp != "n" {
		return ErrRespInvalid
	}
	return nil
}

func Verify(resp string) error {
	if resp != "cara" && resp != "coroa" {
		return ErrRespInvalid
	}
	return nil
}

func ToLower(s string) string {
	return strings.ToLower(s)
}
