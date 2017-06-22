package timers

import (
	"log"
	"time"
)

// TimeTrack measures the time passed since calling this function
// Use it preferably in combination with defer
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
