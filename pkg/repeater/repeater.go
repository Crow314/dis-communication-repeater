package repeater

import "github.com/Crow314/im920s-controller/pkg/module"

func Run(im920s *module.Im920s, detach bool, storeSize int) {
	store := newIDStore(storeSize)

	if detach { // background
		go runner(im920s, store)
	} else { // foreground
		runner(im920s, store)
	}
}

func runner(im920s *module.Im920s, store *idStore) {
	receiver := im920s.DataReceiver()

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
