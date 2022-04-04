package services

type MutantService struct {
	DnaTableSequence DnaTableSequence
	isMutant         bool
	counts           uint8
}

// Validates if is dna of mutant.
func (c *MutantService) IsMutant() bool {
	for i := 0; i < ADN_HEIGHT_SIZE; i++ {
		c.searchHorizontal(i)
		c.searchVertical(i)
		c.searchOblicuos(i)
	}

	return c.isMutant
}

func (c *MutantService) getDnaLetterByCoordinate(rowNumber int, columnNumber int) string {
	return c.DnaTableSequence[rowNumber][columnNumber]
}

func (c *MutantService) setSequenceMatch() {
	c.counts++
	if c.counts == MUTANT_SECUENCE_MATCH_SIZE {
		c.isMutant = true
	}
}

func (c *MutantService) searchHorizontal(rowPosition int) {
	if c.isMutant {
		return
	}

	// fmt.Println("→→→ #", rowPosition, "n")
	mutantDetector := MutantDetectorMatch{}

	columnPosition := 0
	for range c.DnaTableSequence[rowPosition] {
		letter := c.getDnaLetterByCoordinate(rowPosition, columnPosition)

		ok := mutantDetector.HasConsecutiveSequences(letter)
		if ok {
			c.setSequenceMatch()
			break
		}
		columnPosition++
	}
}

func (c *MutantService) searchOblicuos(columnPosition int) {
	if c.isMutant {
		return
	}

	mutantDetector := MutantDetectorMatch{}
	nextColumnPosition := columnPosition

	for rowPosition := 0; rowPosition < ADN_WIDTH_SIZE; rowPosition++ {
		if nextColumnPosition == len(c.DnaTableSequence[rowPosition]) {
			continue
		}

		letter := c.getDnaLetterByCoordinate(rowPosition, nextColumnPosition)
		ok := mutantDetector.HasConsecutiveSequences(letter)
		if ok {
			c.setSequenceMatch()
			break
		}

		nextColumnPosition++
	}
}

func (c *MutantService) searchVertical(columnPosition int) {
	if c.isMutant {
		return
	}

	mutantDetector := MutantDetectorMatch{}

	for rowPosition := 0; rowPosition < ADN_HEIGHT_SIZE; rowPosition++ {
		letter := c.getDnaLetterByCoordinate(rowPosition, columnPosition)
		ok := mutantDetector.HasConsecutiveSequences(letter)
		if ok {
			c.setSequenceMatch()
			break
		}
	}
}

func NewMutantService(dnaTableSequence DnaTableSequence) *MutantService {
	return &MutantService{
		DnaTableSequence: dnaTableSequence,
	}
}
