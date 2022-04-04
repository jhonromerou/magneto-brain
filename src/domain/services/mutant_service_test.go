package services

import (
	"reflect"
	"testing"
)

func Test_Mutant_Service(t *testing.T) {

	tests := []struct {
		name               string
		dnaRequestSequence string
		want               bool
	}{
		{
			name:               "dna not is of mutant",
			dnaRequestSequence: `{"dna":["CGACAA", "CTGTGC", "TAGTTT", "AGAGGT", "CCGCAA", "TCACTA"]}`,
			want:               false,
		},
		{
			name:               "is a mutant - horizontal sequence",
			dnaRequestSequence: `{"dna":["AAAATT", "TTCCCC", "AAAACC", "AGAAGT", "CCGCAA", "TCACTA"]}`,
			want:               true,
		},
		{
			name:               "is a mutant - vertical sequence",
			dnaRequestSequence: `{"dna":["AATTGG", "AATTGG", "AATTGG", "AATTGG", "AATTGG", "AATTGG"]}`,
			want:               true,
		},
		{
			name:               "is a mutant - oblicuos sequence",
			dnaRequestSequence: `{"dna":["CGTCAA", "CTGTGC", "TAATTT", "AGAAGT", "CCGCAA", "TCACTA"]}`,
			want:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dna, _ := NewDnaFromBody(tt.dnaRequestSequence)
			dnaTableSequence := dna.GetTable()
			mutantService := NewMutantService(dnaTableSequence)

			got := mutantService.IsMutant()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.Handle() \nactual = %v \nexpect = %v", got, tt.want)
			}
		})
	}
}
