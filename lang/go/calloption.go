package gohelper

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func WithUID(ctx context.Context, uid string) context.Context {
	if uid == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, "X_LZCAPI_UID", uid)
}
