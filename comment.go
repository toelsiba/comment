package comment

import (
	"bufio"
	"bytes"
	"unicode"
	"unicode/utf8"
)

const marker = '#'

func Trim(data []byte) []byte {
	var res bytes.Buffer
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Bytes()
		bs, commented := trimLine(line)
		if len(bs) > 0 {
			res.Write(bs)
			res.WriteByte('\n')
		} else if !commented {
			res.WriteByte('\n')
		}
	}
	return res.Bytes()
}

func trimLine(line []byte) (bs []byte, commented bool) {
	n := 0
	for {
		r, size := utf8.DecodeRune(line)
		if size == 0 {
			break
		}
		line = line[size:]

		if r == marker {
			// read next rune
			r, size = utf8.DecodeRune(line)
			if (size == 0) || (r != marker) {
				commented = true
				break
			}
			line = line[size:]
			// r equal marker
		}

		bs = appendRune(bs, r)

		if !unicode.IsSpace(r) {
			n = len(bs)
		}
	}

	return bs[:n], commented
}

func appendRune(data []byte, r rune) []byte {
	var (
		n    = len(data)
		nmax = n + utf8.UTFMax
	)
	if cap(data) < nmax {
		temp := make([]byte, nmax)
		copy(temp, data)
		data = temp
	}
	size := utf8.EncodeRune(data[n:nmax], r)
	data = data[:n+size]
	return data
}

// For protect comment marker use two markers (# -> ##)
func Shield(data []byte) []byte {
	res := make([]byte, 0, len(data))
	for _, b := range data {
		res = append(res, b)
		if b == marker {
			res = append(res, marker)
		}
	}
	return res
}
