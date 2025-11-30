package protocol

import "fmt"

func (v *VNA) SetSweep(numPoints int) error {
	message := fmt.Sprintf("SENSe1:SWEep:POINts %d", numPoints)
	return v.Write(message)
}

