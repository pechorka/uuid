package uuid

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
)

func IntToUUID(i int64) string {
	arr := int64toByteArray(i)
	rawUUID := [16]byte{
		arr[7], arr[6], arr[5], arr[4],
		arr[3], arr[2],
		arr[1], arr[0],
		arr[0], arr[1], arr[2], arr[3],
		arr[4], arr[5], arr[6], arr[7],
	}
	uuidBuffer := [36]byte{}
	encodeHex(uuidBuffer[:], rawUUID)
	return string(uuidBuffer[:])
}

func UUIDToInt(strUUID string) int64 {
	if len(strUUID) != 36 {
		return -1
	}
	uuid := []byte(strUUID)
	var intBytes [8]byte
	hex.Decode(intBytes[0:2], uuid[21:23])
	hex.Decode(intBytes[2:], uuid[24:])
	return byteArrayToInt64(intBytes)
}

func int64toByteArray(i int64) (arr [8]byte) {
	binary.BigEndian.PutUint64(arr[:], uint64(i))
	return arr
}

func byteArrayToInt64(arr [8]byte) int64 {
	return int64(binary.BigEndian.Uint64(arr[:]))
}

func New() string {
	return IntToUUID(rand.Int63())
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
