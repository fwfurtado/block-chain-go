package blockchain

import (
	"strings"

	"github.com/fwfurtado/blockchain-go/pkg/hashing"
)

const hashPrefix = "00000"

func generateProofOfWorkBy(previousProof int) int {

	proofOfWork := 1

	for ; ; proofOfWork++ {

		if solveThePuzzle(previousProof, proofOfWork) {
			return proofOfWork
		}
	}

}

func solveThePuzzle(previousProof, proofOfWork int) bool {

	square := func(a int) int {
		return a * a
	}

	hash := hashing.Apply(square(proofOfWork) - square(previousProof))

	return strings.HasPrefix(hash, hashPrefix)
}
