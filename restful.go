package iotmaker_server_json

// pt-br: monta uma saída de dados no formato restful
// en: make a data out in restful format
type jSonOut struct {
	Meta    metaJSonOut `json:"Meta"`
	Objects interface{} `json:"Objects"`
}

// pt-br: monta a parte de meta-data do restful
// en: mount a mate-data from a restful data out
type metaJSonOut struct {
	Cache      string   `json:"Cache,omitempty"`
	Limit      int64    `json:"Limit,omitempty"`
	Next       string   `json:"Next,omitempty"`
	Offset     int64    `json:"Offset"`
	Previous   string   `json:"Previous,omitempty"`
	TotalCount int64    `json:"TotalCount,omitempty"`
	Success    bool     `json:"Success"`
	Error      []string `json:"Error,omitempty"`
}
