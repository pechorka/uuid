package uuid

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func IntToUUID(i int) string {
	rawUUID := [16]byte{
		byte(i >> 0), byte(i >> 4), byte(i >> 8), byte(i >> 12),
		byte(i >> 16), byte(i >> 20),
		byte(i >> 24), byte(i >> 28),
		byte(i >> 32), byte(i >> 36), byte(i >> 40), byte(i >> 44),
		byte(i >> 48), byte(i >> 52), byte(i >> 54), byte(i >> 58),
	}
	uuidBuffer := [36]byte{}
	encodeHex(uuidBuffer[:], rawUUID)
	return string(uuidBuffer[:])
}

func New() string {
	return uuid.New().String()
}

func encodeHex(dst []byte, uuid [16]byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
