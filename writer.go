package main

import (
	"fmt"
	"log"
	"socialNetwork/basicModule"
)

func main() {
	log.SetPrefix("writer-main:")
	log.SetFlags(0)
	names := []string{"Gladys", "Ramon", "Jose"}
	message, err := basicModule.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\r\n", message)
}
