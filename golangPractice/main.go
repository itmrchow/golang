package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink from its water bowl",
		"The dark bird of prey lands on a small tree after hunting for fish",
	}

	histogram := make(map[string]int)
	wordsCh := make(chan string)

	go func() {
		defer close(wordsCh)

		for _, line := range data {
			//取值出來切割
			words := strings.Split(line, " ")

			for _, word := range words {
				//轉小寫
				word = strings.ToLower(word)
				//放入chennel
				wordsCh <- word
			}
		}
	}()

	<-wordsCh

	for {
		word, opened := <-wordsCh
		if !opened {
			break
		}
		histogram[word]++
	}

	for k, v := range histogram {
		fmt.Println(fmt.Sprintf("%s\t(%d)", k, v))
	}

}
