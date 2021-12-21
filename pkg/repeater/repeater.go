package repeater

import "github.com/Crow314/im920s-controller/pkg/module"

func Run(im920s *module.Im920s, storeSize int) {
	go runner(im920s, storeSize)
}

func runner(im920s *module.Im920s, storeSize int) {
	receiver := im920s.DataReceiver()
	store := newIDStore(storeSize)

	for {
		data := <-receiver

		// 受信未済チェック
		if store.checkExistence(data.Data()) {
			continue
		}

		store.add(data.Data())
		_ = im920s.Broadcast(data.Data())
	}
}
