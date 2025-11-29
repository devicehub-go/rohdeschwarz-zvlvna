package protocol

import "fmt"

func (v *VNA) Reset() error {
	return v.Write("*RST")
}

/*
Stops further command processing until all commands
sent before this have been executed
*/
func (v *VNA) Wait() error {
	return v.Write("*WAI")
}

/*
Stops command processing until '1' is returned, i.e.
until de Operation Complete bit has been set int the ESR
*/
func (v *VNA) WaitOperationComplete() error {
	response, err := v.Query("*OPC?")
	if err != nil {
		return err
	}
	if string(response) != "1" {
		return fmt.Errorf("invalid OPC response, got %s", string(response))
	}
	return nil
}

/*
Triggers a sweep or starts continuos sweeps
*/
func (v *VNA) TriggerSweep() error {
	if err := v.Write("INITiate:IMMediate"); err != nil {
		return err
	}
	return v.WaitOperationComplete()
}
