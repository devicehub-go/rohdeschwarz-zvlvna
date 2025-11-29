package zvlvna

import (
	"github.com/devicehub-go/rohdeschwarz-zvlvna/protocol"
	"github.com/devicehub-go/unicomm"
)

/*
Creates a new instance of Rohde&Schwarz ZVL Network
Analyzer to communicate and control the device
*/
func New(options unicomm.Options) *protocol.VNA {
	options.Delimiter = "\n"
	return &protocol.VNA{
		Communication: unicomm.New(options),
	}
}
