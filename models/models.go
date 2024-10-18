package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FloatMap  map[string]float64
	Data      map[string]interface{} // when are not sure what type is
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
