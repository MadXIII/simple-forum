package errorhandle

import (
	"log"
)

//CheckErr - function to handle errors
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
