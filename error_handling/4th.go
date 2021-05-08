package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("anything.txt")
	if err != nil {
		fmt.Println("err happend", err)

		log.Println("err happened", err)
		log.Fatalln(err)
		panic(err)

	}

}

/*the fatal functions call os.exit(i)after writing the log message
fatalln is equivalent to println()followed by a call to os .exit*/
