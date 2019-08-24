package system

import (
	"log"
	"time"
)

// TimeThis time function and log
func TimeThis(f func() error) error {
	from := time.Now()
	err := f()
	log.Printf("%v used %s", &f, time.Since(from))
	return err
}
