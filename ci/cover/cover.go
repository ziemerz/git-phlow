package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Open("cover/percentage")
	if err != nil {
		log.Panicln(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	//	percentage, _ := strconv.Atoi(args[1])
	//	fmt.Println(percentage)

	if 30 >= 40 {
		fmt.Println("above threshold")
		os.Exit(0)
	} else {
		fmt.Println("below threshold")
		os.Exit(1)
	}

}
