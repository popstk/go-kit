package system

import (
	"log"
	"time"
)

func TimeThis(f func() error) error {
	from := time.Now()
	err := f()
	log.Printf("%v used %s", &f, time.Since(from))
	return err
}

