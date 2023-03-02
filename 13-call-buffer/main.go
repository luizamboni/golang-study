package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Callback func(x []interface{}) error

type CallerBuffer struct {
	inputs           []interface{}
	callback         Callback
	maxLen           int
	timeoutInSeconds int
}

func (s *CallerBuffer) addItem(value interface{}) {
	s.inputs = append(s.inputs, value)
	if len(s.inputs) >= s.maxLen {
		s.flushData()
	}
}

func (s *CallerBuffer) flushData() {
	if len(s.inputs) > 0 {
		var inputs []interface{}
		for _, item := range s.inputs {
			inputs = append(inputs, item)
		}
		s.callback(inputs)
		s.inputs = nil
	}
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

func sendData(items []interface{}) error {
	fmt.Println("flushing ........")
	for _, item := range items {
		concreteItem := item.(ApiMessage)
		fmt.Println(reflect.TypeOf(concreteItem), item)
	}
	return nil
}

var wg sync.WaitGroup

type ApiMessage struct {
	value int
	Type  string
}

func main() {
	wg.Add(1)

	inputs := []int{
		1, 2, 3, 4, 4, 23, 23, 454, 5, 665, 6565,
		5, 5, 3, 4, 54, 545, 55, 52, 1, 2, 9, 7,
	}
	var messages []ApiMessage

	for _, number := range inputs {
		messages = append(messages, ApiMessage{number, "mtc"})
	}

	buffer := &CallerBuffer{
		callback:         sendData,
		maxLen:           10,
		timeoutInSeconds: 10,
	}

	for _, msg := range messages {
		buffer.addItem(msg)
	}

	go buffer.start()

	wg.Wait()
}
