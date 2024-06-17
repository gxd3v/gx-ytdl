package managers

import (
	"context"
	"github.com/gx/gx-ytdl/serverv2/pkg/database"
)

type Ytdl interface {
	List(ctx context.Context, sort, filters map[string]database.Filter)
}
