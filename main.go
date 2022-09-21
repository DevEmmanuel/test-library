package mylogger

import "log"

func LogInfo(msg string)  {
	log.Printf("Info - %v", msg)
}

func LogError(msg string)  {
	log.Printf("Error - %v", msg)
}

func LogDebug(msg string)  {
	log.Printf("Debug - %v", msg)
}