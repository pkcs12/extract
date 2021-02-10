package typless

// ExtractResponse represents extract api response
type ExtractResponse struct {
	ObjectID      string             `json:"object_id"`
	FileName      string             `json:"file_name"`
	ExtractFields []ExtractedField   `json:"extracted_fields"`
	LineItems     [][]ExtractedField `json:"line_items"`
}

// ExtractedField represents extracted field details returned from extract api
type ExtractedField struct {
	Name   string           `json:"name"`
	Values []ExtractedValue `json:"values"`
}

// ExtractedValue represents value hash, returned by extract api
type ExtractedValue struct {
	ConfidenceScore float32 `json:"confidence_score"`
	Value           *string `json:"value"`
}
