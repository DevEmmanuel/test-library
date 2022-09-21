package mylogger

import "log"

func LogInfo(msg string)  {
	log.Printf("Info - %v", msg)
}