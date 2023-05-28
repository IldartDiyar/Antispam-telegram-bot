package utils

import (
	"log"
	"os"
)

// InitLogger - функция для настройки логирования
func InitLogger() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Println("Error opening file: ", err)
	}
	log.SetOutput(f)
}
