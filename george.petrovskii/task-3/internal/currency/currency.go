package currency

type Currency struct {
	NumCode  int        `xml:"NumCode"  json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    CommaFloat `xml:"Value"    json:"value"`
}
