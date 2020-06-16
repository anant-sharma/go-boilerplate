package clock

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/structs"
)

// GetTimeStamp Method
func GetTimeStamp(c context.Context) structs.Clock {
	return structs.Clock{
		Timestamp: time.Now(),
	}
}
