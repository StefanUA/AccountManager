package structs

import (
	"fmt"
	"strings"
	"time"
)

// JSONTime encapsulates go's native time type to allow custom
// manipulation specifically JSON marshalling
type JSONTime struct {
	time time.Time
}

const layout = "2006-01-02T15:04:05Z07:00"

// MarshalJSON implements
func (jt JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", jt.Format("Mon Jan _2"))
	return []byte(stamp), nil
}

// UnmarshalJSON implements
func (jt JSONTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		jt.time = time.Time{}
		return
	}
	jt.Time, err = time.Parse(layout, s)
	return
}
