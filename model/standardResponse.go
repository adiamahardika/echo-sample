package model

type StandardResponse struct {
	HttpStatus  int         `json:"http_status"`
	Description []string    `json:"description"`
	Result      interface{} `json:"result,omitempty"`
}
