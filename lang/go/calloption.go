package gohelper

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func WithUID(ctx context.Context, uid string) context.Context {
	if uid == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, "X_LZCAPI_UID", uid)
}

func WithCopyAuth(r *http.Request) context.Context {
	h := r.Header
	md := metadata.New(map[string]string{
		"X_LZCAPI_UID": h.Get("X-Hc-User-Id"),
		"Traceparent":  h.Get("Traceparent"),
	})
	return metadata.NewOutgoingContext(r.Context(), md)
}
