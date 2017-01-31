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

	var hourInc = 0;

	if (minute != 0) {
		hourInc = (minute / 60);
		minute = minute % 60;
	}

	if (minute < 0) {
		minute = 60 + minute;
		hourInc--;
	}

	hour += hourInc;

	if (hour > 23) {
		hour = hour % 24;
	} else if (hour < 0) {
		hour = (-hour) % 24;
		hour = 24 - hour;
	}

	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (c Clock) Add(minutes int) Clock {
	
	var hourInc = 0;

	c.minutes += minutes;
	hourInc = (c.minutes / 60);
	c.minutes = c.minutes % 60;

	if (c.minutes < 0) {
		c.minutes = 60 + c.minutes;
		c.hours--;
	}
	
	c.hours += hourInc;

	if (c.hours > 23) {
		c.hours = c.hours % 24;
	} else if (c.hours < 0) {
		c.hours = (-c.hours) % 24;
		c.hours = 24 - c.hours;
	}

	if (c.hours > 23) {
		c.hours = c.hours % 24;
	}
	
	return c;
}

func abs(i int) int {
	if (i < 0) {
		return -i;
	}
	return i;
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
