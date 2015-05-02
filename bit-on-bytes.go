package bitonbytes

import (
	"errors"
	"log"
)

func ShiftByte(in []byte, distance int) ([]byte, []byte, error) {
	// The shifted out bytes in returned in the second output
	// distance is the number of bytes to be shifted.
	// if positive, shift to the right
	// if negative, shift to the left
	// shift right is zero filled.
	l := len(in)
	if l < distance {
		return nil, nil, errors.New("Shift distance " + string(distance) + " is bigger than the length of byte array: " + string(l))
	}
	out := make([]byte, len(in))
	var holder []byte
	if distance < 0 {
		holder = make([]byte, 0-distance)
	} else {
		holder = make([]byte, distance)
	}
	if distance > 0 {
		for i := distance; i < l; i++ {
			out[i] = in[i-distance]
		}
		for i := 0; i < distance; i++ {
			holder[i] = in[l-distance]
		}
	}
	if distance < 0 {
		for i := 0; i < l+distance; i++ {
			out[i] = in[0-distance+i] // [1,2,3] > [2,3,0], distance = -1
		}
		for i := 0; i < 0-distance; i++ {
			holder[i] = in[i]
		}
	}
	return out, holder, nil
}

func ShiftBit(in byte, distance int) (byte, byte, error) {
	// curr in the current byte after shifting
	// adj is contains the bits after shifting, e.g. 10011110 >> 2 => curr:00100111, adj:10000000
	// distance is the number of bits to be shifted.
	// if positive, shift to the right
	// if negative, shift to the left
	// shift right is zero filled
	var adj byte
	adj = 0
	if distance == 0 {
		return in, adj, nil
	}
	if distance > 8 || distance < 8 {
		return adj, adj, errors.New("There are only 8 bits in 1 byte.")
	}

	return in, in, nil
}

func ShiftLeft(in []byte, distance int) ([]byte, error) {
	// distance is the nubmer of bits to be shifted
	var ll int
	l := len(in)
	div := l / 8
	res := l % 8
	if res != 0 {
		ll = div + 1
	}
	log.Println(string(in))
	log.Println(distance)
	if l < ll {
		return nil, errors.New("Shift distance " + string(distance) + " is bigger than the length of byte array: " + string(l) + " bytes")
	}
	out := make([]byte, l)
	if l == ll {
		return out, nil
	}

	//holder := make([]byte, ll)
	return out, nil
}
