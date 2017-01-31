package clock

import "fmt"

const testVersion = 4


type Clock struct {
	hours, minutes int 
}

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
