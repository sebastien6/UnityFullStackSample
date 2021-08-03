package gateway

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// openAPIServer returns OpenAPI specification files located under "/openapiv2/"
func openAPIServer(dir string) http.Handler {
	fs := http.FileServer(http.Dir(dir))
	return http.StripPrefix("/openapiv2/", fs)
}

// healthzServer returns a simple health handler which returns ok.
func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}
