package protocol

import (
	"fmt"
	"math"
	"strconv"
)

type Preamble struct {
	StartFreq float64
	StopFreq  float64
	NumPoints int
}

type WaveForm struct {
	Real      []float64
	Imaginary []float64
	Magnitude []float64
	Phase     []float64
	Frequency []float64
}

/*
Get systems errors
*/
func (v *VNA) GetErrors() (string, error) {
	response, err := v.Query("SYSTem:ERRor?")
	if err != nil {
		return "", err
	}
	return string(response), nil
}

/*
Gets the preamble for the desired channel
*/
func (v *VNA) GetPreamble() (Preamble, error) {
	var preamble Preamble
	var response []byte
	var err error

	response, err = v.Query("SENS1:FREQ:STAR?")
	if err != nil {
		return preamble, err
	}
	fmt.Println("start", response)
	preamble.StartFreq, err = strconv.ParseFloat(string(response), 64)
	if err != nil {
		return preamble, err
	}
	response, err = v.Query("SENS1:FREQ:STOP?")
	if err != nil {
		return preamble, err
	}
	preamble.StopFreq, err = strconv.ParseFloat(string(response), 64)
	if err != nil {
		return preamble, err
	}
	response, err = v.Query("SENS1:SWE:POIN?")
	if err != nil {
		return preamble, err
	}
	preamble.NumPoints, err = strconv.Atoi(string(response))
	if err != nil {
		return preamble, err
	}

	return preamble, nil
}

/*
Starts a measurement in single sweep mode and returns
the acquired data from the device
*/
func (v *VNA) GetSingleWave(trace string) (WaveForm, error) {
	if err := v.SetContinuosSweep(false); err != nil {
		return WaveForm{}, err
	} else if err := v.SelectTrace(trace); err != nil {
		return WaveForm{}, err
	} else if err := v.TriggerSweep(); err != nil {
		return WaveForm{}, err
	}
	return v.GetSData()
}

func (v *VNA) GetSData() (WaveForm, error) {
	var waveform WaveForm

	if err := v.WriteSequence([]string{
		"FORMat:DATA REAL,32",
		"FORMat:BORDer SWAPped",
	}); err != nil {
		return waveform, err
	}

	data, err := v.QueryByteSequence("CALC:DATA? SDATA")
	if err != nil {
		return waveform, err
	}
	values, err := v.ByteToFloatArray(data)
	if err != nil {
		return waveform, err
	} else if len(values)%2 != 0 {
		return waveform, fmt.Errorf("values must be aligned to 2, got %d", len(values))
	}

	n := len(values) / 2
	waveform.Frequency = make([]float64, n)
	waveform.Real = make([]float64, n)
	waveform.Imaginary = make([]float64, n)
	waveform.Magnitude = make([]float64, n)
	waveform.Phase = make([]float64, n)

	preamble, err := v.GetPreamble()
	if err != nil {
		return waveform, err
	}
	factor := (preamble.StopFreq - preamble.StartFreq) / float64(preamble.NumPoints)

	for i := range n {
		real, img := values[i*2], values[i*2+1]
		waveform.Real[i] = real
		waveform.Imaginary[i] = img
		waveform.Magnitude[i] = 20 * math.Log10(math.Sqrt(math.Pow(real, 2)+math.Pow(img, 2)))
		waveform.Phase[i] = math.Atan2(img, real) * (180 / math.Pi)
		if i%2 == 0 {
			waveform.Frequency[i] = float64(i/2) * factor
		}
	}

	return waveform, nil
}