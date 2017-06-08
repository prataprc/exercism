package gigasecond

import "time"

const testVersion = 4

func AddGigasecond(in time.Time) time.Time {
	return in.Add(1000000000 * time.Second)
}
