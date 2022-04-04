package services

const (
	MUTANT_SECUENCE_MATCH_SIZE = 2
	MUTANT_SECUENCE_SIZE       = 4
	MUTANT_PATTERN             = "^(A|T|C|G){6}$"
)

type MutantDetectorMatch struct {
	lastLetter string
	matches    uint8
}

// Validates if a sequence has mutant matches.
func (c *MutantDetectorMatch) HasConsecutiveSequences(letter string) bool {
	if letter != c.lastLetter {
		c.matches = 1
	} else {
		c.matches++
	}
	if c.matches == MUTANT_SECUENCE_SIZE {
		return true
	}

	c.lastLetter = letter
	return false
}
