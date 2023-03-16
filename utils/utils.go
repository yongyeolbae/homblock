package utils

import "log"

func HandelErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
