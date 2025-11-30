package protocol

import (
	"fmt"
	"slices"
	"strings"
)

/*
Gets the stimulus values of the active data or memory trace
*/
func (v *VNA) GetStimulus() ([]float64, error) {
	data, err := v.QueryByteSequence("CALCulate1:DATA:STIMulus?")
	if err != nil {
		return nil, err
	}
	return v.ByteToFloatArray(data)
}

/*
Defines how the measure result at any sweep point is post-processed
*/
func (v *VNA) SetTraceFormat(trace, format string) error {
	if err := v.SelectTrace(trace); err != nil {
		return err
	}
	message := fmt.Sprintf("CALCulate1:FORMat %s", format)
	return v.Write(message)
}

/*
Creates a new trace for a desired measurement
parameter
*/
func (v *VNA) CreateTrace(name, parameter string) error {
	valid := []string{"S11", "S12", "S21", "S22"}
	if !slices.Contains(valid, parameter) {
		return fmt.Errorf(
			"parameter must be %s, got '%s'", 
			strings.Join(valid, ", "), parameter,
		)
	}
	message := fmt.Sprintf("CALCulate1:PARameter:SDEFine '%s', '%s'", name, parameter)
	return v.Write(message)
}

/*
Deletes a trace by its name
*/
func (v *VNA) DeleteTrace(trace string) error {
	message := fmt.Sprintf("CALCulate1:PARamater:DELete '%s'", trace)
	return v.Write(message)
}

/*
Select a target trace by its name
*/
func (v *VNA) SelectTrace(trace string) error {
	message := fmt.Sprintf("CALCulate1:PARameter:SELect '%s'", trace)
	return v.Write(message)
}


/*
Sets marker to a trace
*/
func (v *VNA) SetTraceMarker(trace string, id int, status bool) error {
	var statusStr string = "OFF"
	if err := v.SelectTrace(trace); err != nil {
		return err
	}
	if status {
		statusStr = "ON"
	}
	message := fmt.Sprintf("CALCulate1:MARKer%d %s", id, statusStr)
	return v.Write(message)
}