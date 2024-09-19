package lfs

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

// ObfStrings obfuscates the input strings into human readable hashes.
func ObfStrings(inputs []string) []string {
	result := make([]string, len(inputs))
	for i, input := range inputs {
		result[i] = Obf(input)
	}
	return result
}

// Obf obfuscates the input string into a human readable hash.
func Obf(input string) string {
	return Obfp(input, 0)
}

// Obfp obfuscates the input string into a human readable hash with 0-8 padding_bytes at the end.
func Obfp(input string, padding_bytes int) string {
	if padding_bytes < 0 {
		panic("padding_bytes cannot be negative.")
	} else if padding_bytes > 8 {
		panic("padding_bytes cannot exceed 8.")
	}
	h := sha1.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)

	adverb_index := binary.LittleEndian.Uint32(bs[0:4])
	adjective_index := binary.LittleEndian.Uint32(bs[4:8])
	noun_index := binary.LittleEndian.Uint32(bs[8:12])
	bytes := bs[12 : 12+padding_bytes]

	adverb := adverbs[adverb_index%uint32(len(adverbs))]
	adjective := adjectives[adjective_index%uint32(len(adjectives))]
	noun := nouns[noun_index%uint32(len(nouns))]

	return fmt.Sprintf("%s%s%s%X", adverb, adjective, noun, bytes)
}
