package zvlvna_test

import (
	"fmt"
	"testing"
	"time"

	zvlvna "github.com/devicehub-go/rohdeschwarz-zvlvna"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestGetMeasurement(t *testing.T) {
	vna := zvlvna.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.4.148",
			Port:         5025,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		},
	})
	if err := vna.Connect(); err != nil {
		panic(err)
	}
	defer vna.Disconnect()

	vna.Write("*CLS")
	vna.SetSweep(201)

	start := time.Now()
	wave, err := vna.GetSingleWave("Trc1")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wave.Real))
	fmt.Println(time.Since(start))
	fmt.Println(vna.GetErrors())
}
