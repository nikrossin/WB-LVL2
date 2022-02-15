package main

import "fmt"

type TrafficLight struct {
	green       state
	greenBlink  state
	yellow      state
	red         state
	yellowBlink state

	current state
	id      int
}

func NewTrafficLight(id int) *TrafficLight {
	t := TrafficLight{}
	t.green = &greenState{&t}
	t.greenBlink = &greenBlinkState{&t}
	t.yellow = &yellowState{&t}
	t.red = &redState{&t}
	t.yellowBlink = &yellowBlinkState{&t}
	t.current = t.green
	t.id = id
	return &t
}

type state interface {
	callGreen() error
	callGreenBlink() error
	callYellow() error
	callRed() error
	callYellowBlink() error
}

type greenState struct {
	trafficLight *TrafficLight
}

func (s *greenState) callGreen() error {
	return fmt.Errorf("light green already turn on!")
}

func (s *greenState) callGreenBlink() error {
	s.trafficLight.current = s.trafficLight.greenBlink
	fmt.Println("Turn On Blink Green!")
	return nil
}
func (s *greenState) callYellow() error {
	return fmt.Errorf("light Yellow turn on after Blink Green!")
}

func (s *greenState) callRed() error {
	return fmt.Errorf("light Red turn on after Yellow light!")
}

func (s *greenState) callYellowBlink() error {
	s.trafficLight.current = s.trafficLight.yellowBlink
	fmt.Println("light YellowBlink turn on!")
	return nil
}

type greenBlinkState struct {
	trafficLight *TrafficLight
}

func (s *greenBlinkState) callGreen() error {
	return fmt.Errorf("light Green turn on after Red!")
}

func (s *greenBlinkState) callGreenBlink() error {
	return fmt.Errorf("light green already turn on!")
}
func (s *greenBlinkState) callYellow() error {
	s.trafficLight.current = s.trafficLight.yellow
	fmt.Println("light Yellow turn on!")
	return nil
}

func (s *greenBlinkState) callRed() error {
	return fmt.Errorf("light Red turn on after Yellow light!")
}

func (s *greenBlinkState) callYellowBlink() error {
	s.trafficLight.current = s.trafficLight.yellowBlink
	fmt.Println("light YellowBlink turn on!")
	return nil
}

type yellowState struct {
	trafficLight *TrafficLight
}

func (s *yellowState) callGreen() error {
	return fmt.Errorf("light Green turn on after Red!")
}

func (s *yellowState) callGreenBlink() error {
	return fmt.Errorf("light Green Blink turn on after Green")

}
func (s *yellowState) callYellow() error {
	return fmt.Errorf("light Yellow already turn on!")
}

func (s *yellowState) callRed() error {
	s.trafficLight.current = s.trafficLight.red
	fmt.Println("Red light turn on!")
	return nil
}

func (s *yellowState) callYellowBlink() error {
	s.trafficLight.current = s.trafficLight.yellowBlink
	fmt.Println("light YellowBlink turn on!")
	return nil
}

type redState struct {
	trafficLight *TrafficLight
}

func (s *redState) callGreen() error {
	s.trafficLight.current = s.trafficLight.green
	fmt.Println("light Green turn on!")
	return nil
}

func (s *redState) callGreenBlink() error {
	return fmt.Errorf("light Green Blink turn on after Green")
}
func (s *redState) callYellow() error {
	return fmt.Errorf("light Yellow turn on after Blink Green!")
}

func (s *redState) callRed() error {
	return fmt.Errorf("light Red already turn on!")
}

func (s *redState) callYellowBlink() error {
	s.trafficLight.current = s.trafficLight.yellowBlink
	fmt.Println("light YellowBlink turn on!")
	return nil
}

type yellowBlinkState struct {
	trafficLight *TrafficLight
}

func (s *yellowBlinkState) callGreen() error {
	s.trafficLight.current = s.trafficLight.green
	fmt.Println("light Green turn on!")
	return nil
}
func (s *yellowBlinkState) callGreenBlink() error {
	return fmt.Errorf("light Green Blink turn on after Green")
}
func (s *yellowBlinkState) callYellow() error {
	return fmt.Errorf("light Yellow turn on after Blink Green!")
}

func (s *yellowBlinkState) callRed() error {
	s.trafficLight.current = s.trafficLight.red
	fmt.Println("light Red turn on!")
	return nil
}
func (s *yellowBlinkState) callYellowBlink() error {
	return fmt.Errorf("light YellowBlink already turn on!")
}

func main() {
	trafficLight := NewTrafficLight(1)

	if err := trafficLight.current.callGreen(); err != nil {
		fmt.Println(err)
	}
	if err := trafficLight.current.callGreenBlink(); err != nil {
		fmt.Println(err)
	}
	if err := trafficLight.current.callYellow(); err != nil {
		fmt.Println(err)
	}
	if err := trafficLight.current.callRed(); err != nil {
		fmt.Println(err)
	}
	if err := trafficLight.current.callGreen(); err != nil {
		fmt.Println(err)
	}
	if err := trafficLight.current.callRed(); err != nil {
		fmt.Println(err)
	}
	if err := trafficLight.current.callYellowBlink(); err != nil {
		fmt.Println(err)
	}
}