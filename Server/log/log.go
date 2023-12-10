package log

import "log"

func Print(message ...string) {
	log.Default().Println(message)
}
