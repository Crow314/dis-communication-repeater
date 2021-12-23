package repeater

import (
	"github.com/Crow314/im920s-controller/pkg/module"
	"time"
)

func Run(im920s *module.Im920s, detach bool, storeSize int, sendTimes int, interval int) {
	store := newIDStore(storeSize)

	if detach { // background
		go runner(im920s, store, sendTimes, interval)
	} else { // foreground
		runner(im920s, store, sendTimes, interval)
	}
}

func runner(im920s *module.Im920s, store *idStore, times int, interval int) {
	receiver := im920s.DataReceiver()

	for {
		data := <-receiver

		// 受信未済チェック
		if store.checkExistence(data.Data()) {
			continue
		}

		store.add(data.Data())
		go resend(data.Data(), im920s, times, interval)
	}
}

func resend(data []byte, im920s *module.Im920s, times int, interval int) {
	for i := 0; i < times; i++ {
		_ = im920s.Broadcast(data) // TODO error handling
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
