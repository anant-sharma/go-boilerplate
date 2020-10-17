package clock

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/structs"
	opentracing "github.com/anant-sharma/go-utils/open-tracing"
)

// GetTimeStamp Method
func GetTimeStamp(ctx context.Context) structs.Clock {
	span := opentracing.CreateChildSpanFromContext(ctx, "controller.clock.GetTimestamp")
	defer span.Finish()

	return structs.Clock{
		Timestamp: time.Now(),
	}
}
