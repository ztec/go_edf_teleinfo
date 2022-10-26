package go_edf_teleinfo

import "bytes"

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
