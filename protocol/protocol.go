package protocol

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/devicehub-go/unicomm"
)

type VNA struct {
	Communication unicomm.Unicomm
	mutex         sync.Mutex
}

/*
Establishes a connection with the device
*/
func (v *VNA) Connect() error {
	if err := v.Communication.Connect(); err != nil {
		return err
	} else if err := v.Communication.Write([]byte("*IDN?\n")); err != nil {
		return err
	}
	response, err := v.Communication.ReadUntil("\n")
	if err != nil {
		return err
	}
	fmt.Println("Connected to ", string(response))
	return nil
}

/*
Closes the connection with the device
*/
func (v *VNA) Disconnect() error {
	return v.Communication.Disconnect()
}

/*
Returns true if device is connected
*/
func (v *VNA) IsConnected() bool {
	return v.Communication.IsConnected()
}

/*
Writes a message to device
*/
func (v *VNA) Write(message string) error {
	if !v.IsConnected() {
		return fmt.Errorf("no device connected")
	}
	if !strings.Contains(message, "\n") {
		message = message + "\n"
	}
	return v.Communication.Write([]byte(message))
}

/*
Writes a sequence of messages to device
*/
func (v *VNA) WriteSequence(messages []string) error {
	for _, message := range messages {
		if err := v.Write(message); err != nil {
			return err
		}
	}
	return nil
}

/*
Reads byte response from device
*/
func (v *VNA) Query(message string) ([]byte, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if err := v.Write(message); err != nil {
		return nil, err
	}
	response, err := v.Communication.ReadUntil("\n")
	if err != nil {
		return nil, err
	}
	return response[:len(response)-1], nil
}

/*
Reads a sequence of bytes from device
*/
func (v *VNA) QueryByteSequence(message string) ([]byte, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	if err := v.Write(message); err != nil {
		return nil, err
	}
	header, err := v.Communication.Read(2)
	if err != nil {
		return nil, err
	}
	if len(header) != 2 || header[0] != '#' {
		return nil, fmt.Errorf("invalid header, got %s", string(header))
	}
	numDigits := int(header[1] - '0')
	numBytes, err := v.Communication.Read(uint(numDigits))
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(string(numBytes))
	if err != nil {
		return nil, err
	}

	payload := make([]byte, 0)
	for len(payload) != int(n) {
		response, err := v.Communication.Read(uint(n - len(payload)))
		if err != nil {
			return nil, err
		}
		payload = append(payload, response...)
	}
	return payload, nil
}
