package repeater

import (
	"github.com/Crow314/im920s-controller/pkg/module"
	"time"
)

func Run(im920s *module.Im920s, storeSize int, sendTimes int, interval int, mirroringDest chan<- module.ReceivedData) {
	store := newIDStore(storeSize)
	receiver := im920s.DataReceiver()

	for {
		data := <-receiver

		// 受信未済チェック
		if store.checkExistence(data.Data()) {
			continue
		}

		store.add(data.Data())
		go resend(data.Data(), im920s, sendTimes, interval)
		if mirroringDest != nil {
			go mirror(mirroringDest, data)
		}
	}
}

func resend(data []byte, im920s *module.Im920s, times int, interval int) {
	for i := 0; i < times; i++ {
		_ = im920s.Broadcast(data) // TODO error handling
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func mirror(dest chan<- module.ReceivedData, data module.ReceivedData) {
	select {
	case dest <- data:
	case <-time.After(10 * time.Second): // Timeout
	}
}
