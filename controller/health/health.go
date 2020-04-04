package health

import (
	"context"

	"github.com/anant-sharma/go-boilerplate/structs"
)

// Check function to return application health
func Check(c context.Context) structs.Health {
	return structs.Health{
		IsHealthy: true,
	}
}
