package generator

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	letterAlphanumeric = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func RandAlphanumeric(n int) string {
	r := make([]rune, n)

	rl := len(letterAlphanumeric)

	for i := range r {
		r[i] = letterAlphanumeric[rand.Intn(rl)]
	}

	return string(r)
}

func RandSecondaryId() string {
	plain1 := RandAlphanumeric(5)
	plain2 := RandAlphanumeric(5)
	time := time.Now().UTC().Nanosecond()

	return fmt.Sprintf("%s%d%s", plain1, time, plain2)
}
