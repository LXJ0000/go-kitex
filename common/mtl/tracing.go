package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false), // 禁用 openTelemetry 默认的 metrics 因为使用了 prometheus

	)
	return p
}
