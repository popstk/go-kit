package timer

import "github.com/robfig/cron"

var cronGlobal *cron.Cron

func init() {
	cronGlobal = cron.New()
	cronGlobal.Start()
}

// NewCron cron wrapper, return a channel
func NewCron(spec string) (chan struct{}, error) {
	ch := make(chan struct{})
	err := cronGlobal.AddFunc(spec, func() {
		ch <- struct{}{}
	})
	if err != nil {
		return nil, err
	}
	return ch, nil
}
