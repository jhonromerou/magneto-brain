package services

import (
	"encoding/json"
	"regexp"
	"strings"
)

const ADN_WIDTH_SIZE = 6
const ADN_HEIGHT_SIZE = 6

type DnaTableSequence [][]string

type DnaSequence struct {
	dna     []string
	table   DnaTableSequence
	isValid bool
}

// SetDna sets value dna sequence.
func (d *DnaSequence) SetDna(dna []string) {
	d.dna = dna
	for _, letters := range dna {
		ok := d.haveValidLetters(letters)
		if !ok {
			d.setIsInvalid()
			break
		}

		letters := d.getLetters(letters)
		d.table = append(d.table, letters)
		d.setIsValid()
	}
}

// GetTable gets table with dna sequence format.
func (d *DnaSequence) GetTable() DnaTableSequence {
	if d.isValid {
		return d.table
	}
	return nil
}

func (d *DnaSequence) ToString() string {
	return strings.Join(d.dna, ",")

}

func (c *DnaSequence) haveValidLetters(letters string) bool {
	ok, _ := regexp.MatchString(MUTANT_PATTERN, letters)
	return ok
}

func (c *DnaSequence) getLetters(letters string) []string {
	return strings.Split(letters, "")
}

func (c *DnaSequence) setIsValid() {
	c.isValid = true
}

func (c *DnaSequence) setIsInvalid() {
	c.table = DnaTableSequence{}
	c.isValid = false
}

type DnaSequenceJSON struct {
	Dna []string `json:"dna"`
}

// NewDnaFromBody gets dna sequence from body.
func NewDnaFromBody(body string) (DnaSequence, error) {
	dnaSequence := DnaSequence{}
	dnaSequenceJSON := DnaSequenceJSON{}

	err := json.Unmarshal([]byte(body), &dnaSequenceJSON)
	if err != nil {
		return DnaSequence{}, err
	}

	dnaSequence.SetDna(dnaSequenceJSON.Dna)
	return dnaSequence, nil
}
