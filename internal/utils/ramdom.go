package utils

import (
	"math/rand"
	"time"
)

func GenerateSixRandomDigit() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := random.Intn(900000) + 100000 // Generates number between 100000-999999
	return otp
}
