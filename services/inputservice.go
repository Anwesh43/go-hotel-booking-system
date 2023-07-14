package services

import (
	"bufio"
	"os"
)

type InputService struct {
	inputs []string
}

func (is *InputService) StartProcessing(ch chan []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		if word == "Quit" {
			break
		}
		is.inputs = append(is.inputs, word)
	}
	ch <- is.inputs
}

func NewInputService() InputService {
	return InputService{
		inputs: make([]string, 0),
	}
}
