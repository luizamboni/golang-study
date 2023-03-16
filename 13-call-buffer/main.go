package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Callback func(x []interface{}) error

type CallerBuffer struct {
	inputs    []interface{}
	callback  Callback
	maxLen    int
	timeLimit int
	ticker    *time.Ticker
	flushed   chan bool
}

func (s *CallerBuffer) addItem(value interface{}) {

	s.inputs = append(s.inputs, value)
	if len(s.inputs) >= s.maxLen {
		fmt.Printf("MAX ITEMS flushData: %v \n", time.Now().Format("15:04:05"))
		s.flushed <- true
	}
}

func (s *CallerBuffer) flushData() {
	if len(s.inputs) > 0 {
		s.callback(s.inputs)
		s.inputs = nil
	}
}

func (s *CallerBuffer) start() {
	s.ticker = time.NewTicker(tickerDuration(s.timeLimit))

	for {
		select {

		case <-s.ticker.C:
			fmt.Printf("TIMEOUT REACHED flushData: %v \n", time.Now().Format("15:04:05"))
			s.flushData()

		case <-s.flushed:
			fmt.Println("RESETING ticker")
			s.flushData()
			s.ticker.Stop()
			s.ticker = nil
			s.ticker = time.NewTicker(tickerDuration(s.timeLimit))
		}
	}
}

func tickerDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

func NewCallerBuffer(timeLimit int, maxLen int, callback Callback) *CallerBuffer {

	buffer := &CallerBuffer{
		callback:  sendData,
		maxLen:    maxLen,
		timeLimit: timeLimit,
		flushed:   make(chan bool),
	}

	go buffer.start()
	return buffer
}

func sendData(items []interface{}) error {
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

	var messages []ApiMessage

	for i := 1; i <= 14; i++ {
		messages = append(messages, ApiMessage{i, "mtc"})
	}
	fmt.Printf("INITED %v \n", time.Now().Format("15:04:05"))

	buffer := NewCallerBuffer(5, 5, sendData)

	for _, msg := range messages[0:10] {
		time.Sleep(500 * time.Millisecond)
		buffer.addItem(msg)
	}

	for _, msg := range messages[10:14] {
		time.Sleep(500 * time.Millisecond)
		buffer.addItem(msg)
	}

	wg.Wait()
}
