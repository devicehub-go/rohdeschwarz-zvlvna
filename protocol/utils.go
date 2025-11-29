package protocol

import (
	"encoding/binary"
	"fmt"
	"math"
)

/*
Converts a byte array to an array of float64 values
*/
func (v *VNA) ByteToFloatArray(payload []byte) ([]float64, error) {
	if len(payload)%4 != 0 {
		return nil, fmt.Errorf("payload not aligned to 4 bytes, got %d", len(payload))
	}
	numValues := len(payload) / 4

	values := make([]float64, numValues)
	for i := range numValues {
		bits := binary.LittleEndian.Uint32(payload[i*4 : i*4+4])
		values[i] = float64(math.Float32frombits(bits))
	}

	return values, nil
}
