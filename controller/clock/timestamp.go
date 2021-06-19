package clock

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/structs"
	newrelictracing "github.com/anant-sharma/go-utils/new-relic/tracing"
)

// GetTimeStamp Method
func GetTimeStamp(ctx context.Context) structs.Clock {
	segment, _ := newrelictracing.NewSegment(ctx, "controller.clock.GetTimestamp")
	defer segment.End()

	return structs.Clock{
		Timestamp: time.Now(),
	}
}
