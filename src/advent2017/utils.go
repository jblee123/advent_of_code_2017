package advent2017

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
)

func ParseInt(s string, label string) int {
	if label == "" {
		label = "<unknown>"
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("couldn't parse %s int: %v\n", label, err)
	}

	return val
}

func GetScanOnByteFunc(b byte) bufio.SplitFunc {
	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, b); i >= 0 {
			// We have a full byte-terminated token.
			return i + 1, data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated token. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	}

	return splitFunc
}
