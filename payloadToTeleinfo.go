package go_edf_teleinfo

import (
	"errors"
	"strconv"
	"strings"
)

// PayloadToTeleinfo Convert text from EDF teleinfo to a proper structure
func PayloadToTeleinfo(edfPayload []byte) (Teleinfo, error) {
	teleinfo := Teleinfo{
		RAW: edfPayload,
	}
	for _, line := range strings.Split(string(edfPayload), "\n") {
		line = strings.Replace(line, "\r", "", -1)
		name, data, err := LineDecoder(line)
		if err != nil {
			return teleinfo, err
		}
		number, _ := strconv.Atoi(data)

		switch name {
		case "OPTARIF":
			teleinfo.OPTARIF = data
		case "HHPHC":
			teleinfo.HHPHC = data
		case "PTEC":
			teleinfo.PTEC = data
		case "ISOUSC":
			teleinfo.ISOUSC = int64(number)
		case "HCHC":
			teleinfo.HCHC = int64(number)
		case "HCHP":
			teleinfo.HCHP = int64(number)
		case "IINST":
			teleinfo.IINST = int64(number)
		case "IMAX":
			teleinfo.IMAX = int64(number)
		case "PAPP":
			teleinfo.PAPP = int64(number)
		}
	}
	return teleinfo, nil
}

// LineDecoder parse one line of EDF teleinfo and check data validity
func LineDecoder(rawLine string) (name string, data string, err error) {

	parts := strings.Split(rawLine, " ")
	if len(parts) < 3 {
		err = errors.New(
			"line is not the right format. Lines should look like 'PAPP 00290 ,' (Etiquette / Donnée / Checksum)",
		)
		return
	}

	sum := 0
	rawLength := len(rawLine)
	for index, char := range []byte(rawLine) {
		if index >= rawLength-2 {
			continue
		}
		sum = sum + int(char)
	}
	sum = (sum & 63) + 32
	if sum == int([]byte(rawLine)[rawLength-1]) {
		name = parts[0]
		data = parts[1]
	} else {
		err = errors.New("checksum verification failed")
		return
	}

	return
}
