package metrics

import "github.com/gin-gonic/gin"

func ApplyMetrics(group *gin.RouterGroup) {
	group.Any("/ping", ping)
	group.GET("/metrics", metrics)
}
