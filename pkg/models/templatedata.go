package models

// template data 
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]any
	CSRFToken string 
	Error string
}
