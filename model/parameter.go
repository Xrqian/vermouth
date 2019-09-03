package model

type Parameter struct {
	Page     int
	PageSize int
	Search   string
	Sort     string
	Desc     bool
	Filter   map[string]interface{}
}
