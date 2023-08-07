package main

import (
	"net/http"
	"google.golang.org/grpc"
	"strings"
)

//
/////
			// https://github.com/grpc/grpc-go/blob/master/server.go#L992
			// https://github.com/philips/grpc-gateway-example/blob/master/cmd/serve.go#L51
			// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
			// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(rpcServer *grpc.Server, swgHandler http.Handler, othHandler http.Handler) http.Handler {
	return http.HandlerFunc(
	func(
			w http.ResponseWriter, r *http.Request,
		) {
	// TODO(tamird): point to merged gRPC code rather than a PR.
	// This is a partial recreation of gRPC's internal checks \
	// https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
	if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") && true { rpcServer.ServeHTTP(w, r) } else {
			if strings.Contains(r.RequestURI,/*swg->,*/"/openapi") {
					if r.RequestURI=="/openapi/v3" { http.ServeFile(w, r, "api/vipex.cc/wx/openapiv3.yml") } else {
					swgHandler.ServeHTTP(w, r)
					}
			} else {
							othHandler.ServeHTTP(w, r)
			}
	}
	})
}

//

func init() {
	//
}