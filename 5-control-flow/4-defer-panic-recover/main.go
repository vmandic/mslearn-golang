package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	for i := 1; i < 5; i++ {
		defer println("Deferred execution:", -i)
		// defer executes in reverese order at the end of current func
		println("Classic execution:", i)
	}

	rand.Seed(time.Now().Unix())

	learningGoMsg := fmt.Sprintf("Learning Go %s", strconv.Itoa(rand.Intn(10000)))
	println(learningGoMsg)

	newFile, err := os.Create("learnGo.txt")

	if err != nil {
		println("Err: Could not create file")
		return
	}

	defer newFile.Close()
	defer newFile.Sync()

	if _, err = io.WriteString(newFile, learningGoMsg); err != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

}
