package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	var letters = []rune("testuserndjdismaluehsmsldjd8a6egenidhnalsoduenkdoshsyabdlieee")
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)
	for i := range b {
		// nolint:gosec
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
