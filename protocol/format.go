package protocol

import (
	"fmt"
	"slices"
	"strings"
)

/*
Controls wether binary data is transferred in normal
or swapped byte order. Formats can be:
	- SWAPped (little endian), default value
	- NORMal (big endian)
*/
func (v *VNA) SetFormatBorder(format string) error {
	valid := []string{"SWAPped", "NORMal"}
	if !slices.Contains(valid, format) {
		return fmt.Errorf(
			"border format must be %s, got %s",
			strings.Join(valid, ", "), format,
		)
	}
	message := fmt.Sprintf("FORMat:BORDer %s", format)
	return v.Write(message)
}

/*
Selects the format for numeric data transferred to
and from the analyzer
*/
func (v *VNA) SetFormatData(format string) error {
	valid := []string{"ASCII", "REAL,32", "REAL,64"}
	if !slices.Contains(valid, format) {
		return fmt.Errorf(
			"data format must be %s, got %s",
			strings.Join(valid, ", "), format,
		)
	}
	message := fmt.Sprintf("FORMat:DATA %s", format)
	return v.Write(message)
}