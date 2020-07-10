package events

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

// IDGeneratorFunc function generating ids
type IDGeneratorFunc func(event Event) (string, error)

// DefaultIDGenerator default id generator
func DefaultIDGenerator(event Event) (string, error) {
	values := make([]string, len(event.Attributes))
	i := 0
	for k, v := range event.Attributes {
		values[i] = k + ":" + v
		i++
	}
	sort.Strings(values)
	plain := strings.Join(values, ",")
	hash := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(hash[:]), nil
}
