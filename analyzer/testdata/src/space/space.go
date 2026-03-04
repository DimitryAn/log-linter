package space

import "log"

func Space() {
	log.Print(" lk") // want "message must start with lowercase letter"

	log.Fatal(" сава") // want "message must start with lowercase letter" "message must contains only english letters, <space> and digits"

}
