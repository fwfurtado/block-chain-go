package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"strconv"
)

func From(value interface{}) string {
	hasher := sha256.New()
	bytes, _ := json.Marshal(value)

	io.WriteString(hasher, string(bytes[:]))

	return hex.EncodeToString(hasher.Sum(nil))
}

func Apply(value int) string {
	hasher := sha256.New()
	strValue := strconv.Itoa(value)
	io.WriteString(hasher, strValue)

	return hex.EncodeToString(hasher.Sum(nil))
}
