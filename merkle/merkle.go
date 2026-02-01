package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

func hash(Data string) string {
	h := sha256.Sum256([]byte(Data))
	return hex.EncodeToString(h[:])
}

func BuildMerkleRoot(transaction []string) string {

	if len(transaction) == 0 {
		return ""
	}

	var level []string

	for _, tx := range transaction {
		level = append(level, hash(tx))
	}

	for len(level) > 1 {
		var nextLevel []string
		for i := 0; i < len(level); i += 2 {

			// if odd number of nodes, duplicate last
			if i+1 == len(level) {
				level = append(level, level[i])
			}

			combined := level[i] + level[i+1]
			nextLevel = append(nextLevel, hash(combined))
		}

		level = nextLevel
	}

	// final remaining hash is Merkle Root
	return level[0]
}
