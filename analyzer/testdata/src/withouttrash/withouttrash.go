package withouttrash

import (
	"log"
)

func Lowercase() {
	log.Print("server started! ❌")                  // want "message must contains only english letters, <space> and digits"
	log.Fatal("connection failed!!! 😀")             // want "message must contains only english letters, <space> and digits"
	log.Print("warning: something went wrong... 🚀") // want "message must contains only english letters, <space> and digits"
	log.Panic("l❌k")                                // want "message must contains only english letters, <space> and digits"
	log.Fatal("server started", "S❌...!!!!;;;")     // want "message must contains only english letters, <space> and digits" "message must start with lowercase letter"
	log.Fatal("S.!?")                               // want "message must contains only english letters, <space> and digits" "message must start with lowercase letter"
}
