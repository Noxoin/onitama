package server

import (
	"math/rand"
	"strconv"
)

func GenerateRandomId() string {
	return strconv.FormatUint(rand.Uint64(), 36)
}
