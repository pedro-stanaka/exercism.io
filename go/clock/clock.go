// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"


// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	hours, minutes int 
} // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	var hour = c.hours;
	var minutes = c.minutes;
	var hourInc = 0;

	if (minutes != 0)  {
		hourInc = (c.minutes / 60);
		minutes = c.minutes % 60;
	}

	if (minutes < 0) {
		minutes = 60 + minutes;
		hourInc --;
	}

	hour += hourInc;
	
	if (hour > 23) {
		hour = hour % 24;
	} else if (hour < 0) {
		hour = -hour % 24;
		hour = 24-hour;
	}

	return fmt.Sprintf("%02d:%02d", hour, minutes)
}

func (c Clock) Add(minutes int) Clock {
	if (minutes + c.minutes > 59) {
		c.minutes = (minutes + c.minutes) % 60;
		c.hours;
	} else {
		c.minutes += minutes;
	}

	if ()
	return c;
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
