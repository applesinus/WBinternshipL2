package main

import (
	"testing"
	"time"

	"github.com/beevik/ntp"
)

// Не совсем понимаю, какие тут можно в принципе сделать тесты
func TestMain(t *testing.T) {
	// getting time query from ntp
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		// if there's an error, print it
		t.Fatalf("Error while getting time from ntp: %v", err)
	}

	// comparing current time on the local machine with time from ntp
	now := time.Now().Add(response.ClockOffset)
	if response.Time.UTC().Sub(now.UTC()) > time.Microsecond {
		// if there's a gap more than 1 microsecond, print an error
		t.Fatalf("Expected %v, got %v", now, response.Time)
	}
}
