package protocol

import "fmt"

func (v *VNA) SetSweep(numPoints int) error {
	message := fmt.Sprintf("SENSe1:SWEep:POINts %d", numPoints)
	return v.Write(message)
}

/*
Set the device to operate in continuos or single sweep mode
*/
func (v *VNA) SetContinuosSweep(status bool) error {
	statusStr := "OFF"
	if status {
		statusStr = "ON"
	}
	if err := v.WriteSequence([]string{
		"INITiate1:SCOPe SINGle",
		fmt.Sprintf("INITiate:CONTinuous %s", statusStr),
	}); err != nil {
		return err
	}
	return v.WaitOperationComplete()
}

/*
Select a target trace by its name
*/
func (v *VNA) SelectTrace(trace string) error {
	message := fmt.Sprintf("CALCulate1:PARameter:SELect '%s'", trace)
	return v.Write(message)
}
