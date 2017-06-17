package example

import (
	"fmt"
	"sync"
	"time"
)

// declare wait group
var wg sync.WaitGroup

////////////////////////////////////////////
//  With go routine
////////////////////////////////////////////
func WithGoRoutine() time.Duration {

	// adding work group
	// for my example, I added 3 work groups
	wg.Add(3)

	// tracking current time
	now := time.Now()

	// declare quite channel
	quit := make(chan bool)

	// calling message with go routine function
	// for the three different variables
	msg1 := MessageWithGoRoutine("message one")
	msg2 := MessageWithGoRoutine("message two")
	msg3 := MessageWithGoRoutine("message three")

	// This go routine wait until 3 work groups done
	// when 3 work groups done, it'll assign true to the quit channel
	go func() {
		wg.Wait()
		quit <- true
	}()

	// This block contains select statement
	// and the work of this select statement
	// is communicating with the 3 difference channel
	// that I defined at the top ( msg1, msg2, and msg3 )
	for {
		select {
		case msg := <-msg1:
			fmt.Println(msg)
			wg.Done()
		case msg := <-msg2:
			fmt.Println(msg)
			wg.Done()
		case msg := <-msg3:
			fmt.Println(msg)
			wg.Done()
		case <-quit:

			// calculating total required time
			timeRequired := time.Now().Sub(now)

			// then printout the required time
			fmt.Print("Total time required with Go routine: ")
			fmt.Println(timeRequired)

			return timeRequired
		}
	}
}

////////////////////////////////////////////
//  Messaging function
// ( Implements go routine feature )
////////////////////////////////////////////
func MessageWithGoRoutine(msg string) chan string {

	// define mssage channel
	message := make(chan string)

	// implementing another go routine
	// this go routine wait for 1 sec
	// after that writing message into the message channel
	go func() {
		time.Sleep(1 * time.Second)

		message <- "Message: " + msg
	}()

	// returning message channel
	return message
}
