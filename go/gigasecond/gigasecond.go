package gigasecond

import "time"

const gigaSecond = 1e9 * time.Second

// AddGigasecond adds 1e9 seconds to the time passed in parameter, returning the result.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
