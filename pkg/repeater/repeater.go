package repeater

import "github.com/Crow314/im920s-controller/pkg/module"

func Run(im920s *module.Im920s) {
	go runner(im920s)
}

func runner(im920s *module.Im920s) {
	receiver := im920s.DataReceiver()

	for {
		data := <-receiver
		_ = im920s.Broadcast(data.Data())
	}
}
