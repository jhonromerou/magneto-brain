package services

import (
	"reflect"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
)

func Test_Dna_Sequence_constructor(t *testing.T) {
	tests := []struct {
		name        string
		dnaSequence string
		wantError   bool
		want        func() DnaSequence
	}{
		{
			name:        "invalid format sequence",
			dnaSequence: "",
			wantError:   true,
			want: func() DnaSequence {
				got, _ := NewDnaFromBody("")
				return got
			},
		},
		{
			name:        "invalidad pattern",
			dnaSequence: `{"dna":["XXXX"]}`,
			wantError:   false,
			want: func() DnaSequence {
				got, _ := NewDnaFromBody(`{"dna":["XXXX"]}`)
				return got
			},
		},
		{
			name:        "success create",
			dnaSequence: `{"dna":["CGACAA", "CTGTGC", "TAGTTT", "AGAGGT", "CCGCAA", "TCACTA"]}`,
			wantError:   false,
			want: func() DnaSequence {
				got, _ := NewDnaFromBody(`{"dna":["CGACAA", "CTGTGC", "TAGTTT", "AGAGGT", "CCGCAA", "TCACTA"]}`)
				return got
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDnaFromBody(tt.dnaSequence)
			want := tt.want()
			if (err != nil) != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, got, want))
			}

			if !reflect.DeepEqual(got, want) {
				t.Fatal(domain.TestingErrorNotMatched(2, got, want))
			}

			gotString := got.ToString()
			wantString := want.ToString()
			if !reflect.DeepEqual(gotString, wantString) {
				t.Error(domain.TestingErrorNotMatched(3, gotString, wantString))
			}

		})
	}
}
