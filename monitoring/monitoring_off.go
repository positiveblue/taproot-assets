//go:build !monitoring
// +build !monitoring

package monitoring

import (
	"fmt"

	"google.golang.org/grpc"
)

// GetPromInterceptors returns the set of interceptors for Prometheus
// monitoring if monitoring is enabled, else empty slices. Monitoring is
// currently disabled.
func GetPromInterceptors() ([]grpc.UnaryServerInterceptor,
	[]grpc.StreamServerInterceptor) {

	return []grpc.UnaryServerInterceptor{}, []grpc.StreamServerInterceptor{}
}

// ExportPrometheusMetrics is required for tapd to compile so that Prometheus
// metric exporting can be hidden behind a build tag.
func ExportPrometheusMetrics(_ *grpc.Server, _ PrometheusCfg) error {
	return fmt.Errorf("tapd must be built with the monitoring tag to " +
		"enable exporting Prometheus metrics")
}
