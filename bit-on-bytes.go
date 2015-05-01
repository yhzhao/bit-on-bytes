package bitonbytes

import (
	"errors"
	"log"
)

func ShiftLeft(in []byte, distance int) ([]byte, error) {
	// distance is the nubmer of bits to be shifted
	l := len(in)
	log.Println(string(in))
	log.Println(distance)
	if l < distance {
		return nil, errors.New("Shift distance " + distance + " is bigger than the length of byte array: " + l)
	}
	if l == distance {
		out := make([]byte, l)
		return out, nil
	}

	div := l / 8
	res := l % 8
	if res != 0 {
		ll := div + 1
	}
	holder := make([]byte, ll)
	return out, nil
}
