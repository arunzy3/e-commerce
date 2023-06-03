package models

type Errors struct {
	Error string `json:"error"`
	Type  string `json:"type"`
	Param string `json:"param,omitempty"`
}
