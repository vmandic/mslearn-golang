package main

import (
	"log"
	"time"
)

func main() {
	log.SetPrefix("main(): ")

	log.Println("Hi, this logged")
	time.Sleep(time.Second)

	log.Println("Hi, this logged")
	time.Sleep(time.Second)

	log.Println("Hi, this logged")
	time.Sleep(time.Second)

	defer log.Fatalln("Fatal log 1")
	defer log.Fatalln("Fatal log 2")
	panic("Panicing!")

	log.Println("after fatal log which never happens")
}
