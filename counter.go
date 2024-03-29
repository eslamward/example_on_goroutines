package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loopThrowLine(line string) {
	wg.Add(1)
	for _, word := range strings.Fields(line) {
		mute.Lock()
		_, ok := counterMap[word]
		if ok {
			counterMap[word] += 1
		} else {
			counterMap[word] = 1
		}
		mute.Unlock()
	}
	wg.Done()
}

func printCounterMap(counterMap map[string]int) {
	for key, value := range counterMap {
		fmt.Printf("%s %d\n", key, value)
	}
}

func handelFile() {
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 004)
	if err != nil {
		fmt.Println("Error Reading The File ...!", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		go loopThrowLine(line)

	}
}
