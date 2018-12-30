package interceptor

import (
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ContextUnaryClientInterceptor intercepts unary gRPC calls from client-side and
// forwards necessary metadata from incoming context to outgoing context.
func ContextUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		pairs := make([]string, 0)

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key, values := range md {
				if strings.HasPrefix(strings.ToLower(key), "x-") {
					for _, value := range values {
						pairs = append(pairs, key, value)
					}
				}
			}
		}

		ctx = metadata.AppendToOutgoingContext(ctx, pairs...)
		return invoker(ctx, method, req, resp, cc, opts...)
	}
}
