package checkup

import "time"

type Config struct {
	Interval  time.Duration
	UniqueIDs bool
}
