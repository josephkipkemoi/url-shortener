package database

import (
	"math/rand"
	"os"
	"time"
)

/*
checkErr function checks for errors and stops the program execution
dumping the error message on the terminal
*/
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

/*
checkFileExists function checks if the database file exists on the server
it will create a new file with the given name if it does not exists
otherwise the existing file will be opened
*/
func checkFileExists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil
}

/*
generateShortKey function generates unique short keys for the original URLS
This function generates a random alphanumeric key of length 6 characters,
ensuring uniqueness of keys
*/
func GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortKey)
}
