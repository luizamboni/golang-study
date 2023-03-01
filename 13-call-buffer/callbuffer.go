package main

import (
	"fmt"
	"sync"
	"time"
)

type CallerBuffer struct {
	inputs           []int
	call             func(val int) error
	maxLen           int
	timeoutInSeconds int
}

func (s *CallerBuffer) addItem(value int) {
	s.inputs = append(s.inputs, value)
	if len(s.inputs) >= s.maxLen {
		s.flushData()
	}
}

func (s *CallerBuffer) flushData() {
	fmt.Println("send inputs: ", s.inputs)
	s.inputs = nil
}

func (s *CallerBuffer) start() {
	ticker := time.NewTicker(time.Duration(s.maxLen) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				s.flushData()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func sendData(number int) error {
	return nil
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	inputs := []int{1, 2, 3, 4, 4, 23, 23, 454, 5, 665, 6565, 5, 5, 3, 4, 54, 545, 55, 52, 1, 2, 9, 7}

	buffer := &CallerBuffer{
		call:             sendData,
		maxLen:           10,
		timeoutInSeconds: 10,
	}

	for _, item := range inputs {
		buffer.addItem(item)
		// buffer.sendDataIfInTime()
	}
	go buffer.start()

	wg.Wait()
	// fmt.Println(buffer)
}
