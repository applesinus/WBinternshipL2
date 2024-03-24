package main

import (
	"testing"
	"time"

	"github.com/beevik/ntp"
)

func TestMain(t *testing.T) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Fatalf("Error while getting time from ntp: %v", err)
	}
	now := time.Now().Add(response.ClockOffset)
	if response.Time.UTC().Sub(now.UTC()) > time.Microsecond {
		t.Fatalf("Expected %v, got %v", now, response.Time)
	}
}
