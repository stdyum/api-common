package metrics

import (
	"github.com/stdyum/api-common/hc"
)

func ApplyMetrics(group *hc.RouterGroup) {
	group.Any("/ping", ping)
	group.GET("/metrics", metrics)
}
