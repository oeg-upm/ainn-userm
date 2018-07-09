package main

import (
	"math/rand"
	"time"
)

func encrypt(raw_data string) string {
	//  h := md5.New()
	//io.WriteString(h, "The fog is getting thicker!")
	//io.WriteString(h, "And Leon's getting laaarger!")
	//fmt.Printf("%x", h.Sum(nil))
	return raw_data
}

/*source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang*/
func getRandomSequence(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
