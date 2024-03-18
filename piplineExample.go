package main

// Create function that take string as input
// then in first stage convert into slice of string
// then in second stage capatlize every item in the slice
// then in third stage loop throw the list and print it

func ConvertStringToChanel(st string) <-chan rune {

	ch := make(chan rune) // unbufferd chan

	for _, char := range st {
		ch <- char
	}
	close(ch)

	return ch
}

func CapatlizeCharcter(ch <-chan rune) <-chan string {

	ch2 := make(chan string)

	for item := range ch {

		ch2 <- string(item)
	}
	close(ch2)
	return ch2

}
