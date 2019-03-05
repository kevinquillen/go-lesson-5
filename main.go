package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const distanceToMars = 62100000

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Spaceline    Days    Trip-Type    Price (in millions)")
	fmt.Println("=====================================================")

	for total := 0; total < 10; total++ {
		ticket := RandomTicket()

		// logic here to determine price based on distance, trip type, carrier speed

		fmt.Printf("%v    %v    %v    $100\n", ticket.carrier, ticket.tripLength, ticket.tripType)
	}
}

// RandomTicket - returns a randomized ticket.
func RandomTicket() *Ticket {
	const (
		departureDate = "2020-10-21"
		layoutISO     = "2006-01-02"
		layoutUS      = "January 2, 2006"
	)

	// is there a better / more normalized way to do this?
	carriers := [3]string{"Space Adventures", "Virgin Galactic", "Space X"}
	tripTypes := [2]string{"Round Trip", "One Way"}
	departs, _ := time.Parse(layoutISO, departureDate)

	return &Ticket{
		carrier:       carriers[rand.Intn(len(carriers))],
		tripType:      tripTypes[rand.Intn(len(tripTypes))],
		tripLength:    random(20, 50),
		departureDate: departs.Format(layoutUS),
	}
}

// This does not exist in Go?
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Ticket holds information about a ticket.
type Ticket struct {
	carrier       string
	tripType      string
	tripLength    int
	departureDate string
}
