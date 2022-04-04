package domain

type AnalysisModel struct {
	DnaType     string `dynamodbav:"dna_type" json:"dna_type"`
	DnaSequence string `dynamodbav:"dna_sequence" json:"dna_sequence"`
}

type StatsModel struct {
	Group    string `dynamodbav:"group"`
	Name     string `dynamodbav:"name"`
	Quantity int    `dynamodbav:"quantity"`
}

const (
	STATS_GROUP_DNA_ANALIZE = "dnaAnalize"
	STATS_NAME_DNA_MUTANT   = "dnaMutant"
	STATS_NAME_DNA_HUMAN    = "dnaHuman"
)
