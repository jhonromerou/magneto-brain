package domain

type AnalysisValidatorResponse struct {
	Name   string            `json:"name"`
	Format string            `json:"format"`
	Values map[string]string `json:"values,omitempty"`
}
