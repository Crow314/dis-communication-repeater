package main

import (
	"flag"
	"github.com/Crow314/dis-communication-repeater/pkg/repeater"
	"github.com/Crow314/im920s-controller/pkg/connector"
	"github.com/Crow314/im920s-controller/pkg/module"
)

func main() {
	storeSize := flag.Int("s", 256, "remember packet identifications specified number of times")
	flag.Parse()
	tty := flag.Args()[0]

	conn := connector.NewConnector(tty)
	im920s := module.NewIm920s(conn.TransmitChannel(), conn.ReceiveChannel())

	repeater.Run(im920s, *storeSize)
}
