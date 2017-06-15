package with_and_without_go_routine

import (
	"fmt"
	"time"
)

////////////////////////////////////////////
//  Without go routine
////////////////////////////////////////////
func WithoutGoRoutine() time.Duration {

	// declare quite channel
	now := time.Now()

	// printing out message
	for i := 1; i <= 3; i++ {
		fmt.Println(Message("bar"))
	}

	// calculating total required time
	timeRequired := time.Now().Sub(now)

	// then printout the required time
	fmt.Print("Total time required without Go routine: ")
	fmt.Println(timeRequired)

	return timeRequired
}

////////////////////////////////////////////
//  Messaging function
// ( Doesn't implements go routine feature )
////////////////////////////////////////////
func Message(msg string) string {

	// wait for 1 sec
	time.Sleep(1 * time.Second)

	// then write message
	msg = "Message: " + msg

	// and returns this message
	return msg
}
