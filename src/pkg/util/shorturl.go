package util

import (
	"fmt"
	"strings"
)

const (
	Base62Chars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Base            = 62
	ShortLinkLength = 8
)

func (u *Util) CreateShortLink(postID, userID uint) string {
	userIDEncoded := u.EncodeBase62(userID, 2)

	postIDEncoded := u.EncodeBase62(postID, ShortLinkLength-len(userIDEncoded))

	encoded := userIDEncoded + postIDEncoded
	return "localhost:8080/p?id=" + encoded
}

func (u *Util) DecodeShortLink(shortLink string) (postID, userID uint, err error) {
	encoded := strings.TrimPrefix(shortLink, "localhost:8080/p?id=")
	userIDLength := 2

	userIDEncoded := encoded[:userIDLength]
	encoded = encoded[userIDLength:]

	postIDEncoded := encoded

	userID, err = u.DecodeBase62(userIDEncoded)
	if err != nil {
		return 0, 0, err
	}

	postID, err = u.DecodeBase62(postIDEncoded)
	if err != nil {
		return 0, 0, err
	}

	return postID, userID, nil
}

func (u *Util) EncodeBase62(n uint, length int) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(length)

	for n > 0 {
		remainder := n % Base
		encodedBuilder.WriteByte(Base62Chars[remainder])
		n /= Base
	}

	encoded := encodedBuilder.String()

	for len(encoded) < length {
		encoded = "0" + encoded
	}

	return u.ReverseString(encoded)
}

func (u *Util) DecodeBase62(encoded string) (uint, error) {
	decoded := 0
	multiplier := 1

	for i := len(encoded) - 1; i >= 0; i-- {
		char := encoded[i]
		index := strings.IndexByte(Base62Chars, char)
		if index == -1 {
			return 0, fmt.Errorf("invalid base62 character: %c", char)
		}
		decoded += index * multiplier
		multiplier *= Base
	}

	return uint(decoded), nil
}

func (u *Util) ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
