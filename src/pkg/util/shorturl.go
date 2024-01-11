package util

import (
	"encoding/base64"
)

func (u *Util) EncodeToShortURL(id1, id2 uint) string {
	bytesToEncode := []byte{byte(id1 >> 8), byte(id1), byte(id2 >> 8), byte(id2)}

	encodedString := base64.URLEncoding.EncodeToString(bytesToEncode)

	return "http://localhost:8080/p?id=" + encodedString
}

func (u *Util) DecodeShortURL(encodedString string) (uint, uint, error) {
	decodedBytes, err := base64.URLEncoding.DecodeString(encodedString)
	if err != nil {
		return 0, 0, err
	}

	id1 := uint(decodedBytes[0])<<8 | uint(decodedBytes[1])
	id2 := uint(decodedBytes[2])<<8 | uint(decodedBytes[3])

	return id1, id2, nil
}
