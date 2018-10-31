package utils
import (
	"log"
)

func Recover(err interface{}) {
	if r := recover(); r != nil {
		log.Printf("%v", err)
	}
}
