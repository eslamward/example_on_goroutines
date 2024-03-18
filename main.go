package main

import (
	"fmt"
	"os"
	"sync"
)

var counterMap = make(map[string]int)
var wg *sync.WaitGroup
var mute *sync.Mutex

func main() {

	if len(os.Args) > 1 {
		handelFile()
		printCounterMap(counterMap)
		wg.Wait()
	} else {
		fmt.Println("Test PipeLine Example ...!")

		ch1 := ConvertStringToChanel("hello world")
		ch2 := CapatlizeCharcter(ch1)

		for char := range ch2 {
			fmt.Println(char)
		}

	}

}
