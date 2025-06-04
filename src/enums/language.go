package enums

import "encoding/json"

type Language int

const (
	English Language = iota + 1
	Lithuanian
)

func (lang Language) String() string {
	return [...]string{"", "English", "Lithuanian"}[lang]
}

func (lang Language) MarshalJSON() ([]byte, error) {
	return json.Marshal(lang.String())
}
