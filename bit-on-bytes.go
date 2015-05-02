package bitonbytes

import (
	"errors"
	"log"
)

func ShiftByte(in []byte, distance int) ([]byte, error) {
	// distance is the number of bytes to be shifted.
	// if positive, shift to the right
	// if negative, shift to the left
	if l < distance {
		return nil, errors.New("Shift distance " + distance + " is bigger than the length of byte array: " + len(in))
	}
	out := make([]byte, len(in))
	if distance > 0 {
		for i := distance; i < l; i++ {
			out[i] = in[i-distance]
		}
	}
	if distance < 0 {
		for i := 0; i < 0-distance; i++ {
			out[i] = in[0-distance+i] // [1,2,3] > [2,3,0], distance = -1
		}
	}
	return out, nil
}

func ShiftLeft(in []byte, distance int) ([]byte, error) {
	// distance is the nubmer of bits to be shifted
	l := len(in)
	div := l / 8
	res := l % 8
	if res != 0 {
		ll := div + 1
	}
	log.Println(string(in))
	log.Println(distance)
	if l < ll {
		return nil, errors.New("Shift distance " + distance + " is bigger than the length of byte array: " + l + " bytes")
	}
	if l == ll {
		out := make([]byte, l)
		return out, nil
	}

	holder := make([]byte, ll)
	return out, nil
}
