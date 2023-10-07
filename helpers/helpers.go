package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max - min + 1) + min)
}

func GetStatus(data int) string {
	if data <= 5 {
		return "aman"
	}

	if data >= 6 && data <= 8 {
		return "siaga"
	}

	return "bahaya"
}
