package protocol

import (
	"fmt"
)

/*
Creates or delete a diagram area
*/
func (v *VNA) SetWindowState(window int, state bool) error {
	var stateStr string = "OFF"
	if window <= 0 {
		return fmt.Errorf("window must be greater than 0, got %d", window)
	}
	if state {
		stateStr = "ON"
	}
	message := fmt.Sprintf("DISPlay:WINDow%d:STATe %s", window, stateStr)
	return v.Write(message)
}

/*
Display trace in a diagram area
*/
func (v *VNA) DisplayTrace(window, traceId int, trace string) error {
	message := fmt.Sprintf("DISPlay:WINDow%d:TRACe%d:FEED '%s'", window, traceId, trace)
	return v.Write(message)
}