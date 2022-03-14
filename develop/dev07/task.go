package main

import (
	"fmt"
	"reflect"
	"time"
)

// OR1 Каждый канал обрабатывается в своей горутине
func OR1(channels ...<-chan interface{}) <-chan interface{} {
	start := make(chan struct{})
	done := make(chan interface{})
	if len(channels) == 0 {
		close(done)
	}
	for _, ch := range channels {
		go func(c <-chan interface{}) {
			defer close(done)
			<-start // маркер ожидание, пока не развернутся все горутины
			<-c
		}(ch)
	}
	close(start)
	return done // single канал
}

// OR2 Итеративно проверяем каждый канал в одной горутине
func OR2(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		for {
			for _, ch := range channels {
				select {
				case <-ch:
					close(done)
					return
				default:

				}
			}
		}
	}()

	return done
}

// OR3 Реализация на reflect
func OR3(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})
	defer close(done)
	var refChannels []reflect.SelectCase // слайс каналов
	for _, ch := range channels {
		refCh := reflect.SelectCase{
			Dir:  reflect.SelectRecv,  // <- chan
			Chan: reflect.ValueOf(ch), // value
		}
		refChannels = append(refChannels, refCh)
	}
	reflect.Select(refChannels)
	return done // канал вернется после работы Select
}

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = OR1

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(6*time.Second),
		sig(1*time.Hour),
		sig(5*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
