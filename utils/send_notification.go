package utils

import (
	"log"
)

func SendNotification(message string) bool {
	if len(message) > 0 {
		log.Println(message)
		return true
	}
	return false
}
