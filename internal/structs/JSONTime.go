package structs

import (
	"time"
)

// JSONTime encapsulates go's native time type to allow custom
// manipulation specifically JSON marshalling
type JSONTime struct {
	Time time.Time
}

const layout = "2006-01-02T15:04:05Z07:00"

// MarshalJSON conversts JSONTime to bytes representing a JSON string
func (jt JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(jt.Time.Format(layout)), nil
}

// UnmarshalJSON conversts a JSON string to JSONTime stuct
func (jt *JSONTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	jt.Time, err = time.Parse(layout, string(b))
	return
}
