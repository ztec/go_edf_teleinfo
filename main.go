package go_edf_teleinfo

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// Teleinfo (all) data returnable by EDF teleinfo. Remember that I cannot test all cases, meaning some labels are missing. Please contribute to add them
type Teleinfo struct {
	OPTARIF string `json:"OPTARIF"` //abonement
	ISOUSC  int64  `json:"ISOUSC"`  //abonement_puissance
	HCHC    int64  `json:"HCHC"`    //index_heures_creuses
	HCHP    int64  `json:"HCHP"`    //index_heures_pleines
	IINST   int64  `json:"IINST"`   //intensitee_instantanee
	IMAX    int64  `json:"IMAX"`    //intensitee_max
	PAPP    int64  `json:"PAPP"`    //puissance_apparente
	HHPHC   string `json:"HHPHC"`   //groupe_horaire
	PTEC    string `json:"PTEC"`    //Période Tarifaire en cours
	RAW     []byte `json:"RAW"`     //Raw edf teleinfo payload
}

// ScannerSplitter try to identify start and end of EDF teleinfo payload
func ScannerSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	//Start + End datagram markers
	if bytes.Contains(data, []byte{13, 3, 2, 10}) {
		i := bytes.Index(data, []byte{13, 3, 2, 10})
		return i + 4, data[0:i], nil
	}
	// Request more data.
	return 0, nil, nil
}

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
		err = errors.New("Line is not the right format. Ligne shoud look like 'PAPP 00290 ,' (Etiquette / Donnée / Checksum)")
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
		err = errors.New("Checksum verification failed")
		return
	}

	return
}
