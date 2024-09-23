package utilities

import "math/rand"

func GenerateUserID(n int) string {
	numsAndLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var uid []byte

	for i := 0; i < n; i++ {
		uid = append(uid, numsAndLetters[rand.Intn(len(numsAndLetters))])
	}

	return string(uid)
}
