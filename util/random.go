package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	number   = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func randomNumber(n int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomId() int64 {
	return randomInt(1, 1000)
}

func RandomName() string {
	return RandomString(6)
}

func RandomUsername() string {
	var sb strings.Builder

	sb.WriteString("@")
	sb.WriteString(RandomString(8))

	return sb.String()
}

func RandomPhone() string {
	var sb strings.Builder

	sb.WriteString("08")
	sb.WriteString(randomNumber(10))

	return sb.String()
}

func RandomGender() int64 {
	return randomInt(1, 2)
}

func RandomAge() int64 {
	return randomInt(10, 80)
}

func RandomDate() time.Time {
	year := randomInt(1900, 2022)
	month := randomInt(1, 12)
	date := randomInt(1, 28)

	randomDate := fmt.Sprintf("%d-%d-%d", year, month, date)

	time, err := time.Parse("2006-1-2", randomDate)
	if err != nil {
		panic(err)
	}

	return time
}
