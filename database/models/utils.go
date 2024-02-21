package models

import "log"

/*
checkErr function checks for errors and stops the program execution
dumping the error message on the terminal
*/
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
