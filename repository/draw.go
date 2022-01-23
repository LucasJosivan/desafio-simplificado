package repository

import (
	"math/rand"
	"time"
)

func Flip() string {
	rand.Seed(time.Now().UnixNano())
	return getResult(rand.Intn(2))

}

func getResult(flip int) string {
	switch flip {
	case 0:
		return "cara"
	case 1:
		return "coroa"
	default:
		return Flip()
	}
}
