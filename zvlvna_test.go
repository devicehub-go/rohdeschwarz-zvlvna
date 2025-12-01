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
			Host:         "169.254.21.159",
			Port:         5025,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		},
	})
	if err := vna.Connect(); err != nil {
		panic(err)
	}
	defer vna.Disconnect()

	vna.Reset()
	vna.Write("*CLS")
	vna.SetSweep(101)
	vna.SetContinuosSweep(true)

	start := time.Now()
	_, err := vna.GetSData("Trc1")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))
	fmt.Println(vna.GetErrors())
}
