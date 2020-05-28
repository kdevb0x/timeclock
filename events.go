package main

import (
	"encoding/json"
	"time"
)

type timer struct {
	// internal ticker
	ticker time.Ticker

	// is this timer currently ticking/tracking time?
	active bool

	// NOTE: probably not needed here
	total time.Duration

	// event propogation
	trigger chan Event
}

func newTimer() *timer {
	return &timer{
		active:  false,
		ticker:  *time.NewTicker(time.Second),
		trigger: make(chan Event),
	}
}

type Event struct {
	// [kdv] TODO: maybe add "Employee" field to make it multi-user capable
	// incase this is released as a lib.

	// time the event occurred
	Timestamp string    `json:"timestamp,omitempty"`
	Code      EventType `json:"type"`
}

type EventType int

const (
	None EventType = iota
	ClockIn
	ClockOut

	// prefer the more specific types below

	LunchStart
	LunchEnd
	ShiftStart
	ShiftEnd
)

func (e *Event) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, e)

}

func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(e)
}
