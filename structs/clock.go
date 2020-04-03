package structs

import "time"

// Clock Structure
type Clock struct {
	Timestamp time.Time `json:"timestamp"`
}
