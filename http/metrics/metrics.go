package metrics

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/stdyum/api-common/hc"
)

func metrics(ctx *hc.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"RAM": getRAMUsage(),
	})
}

func getRAMUsage() string {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return formatBytes(memStats.HeapSys)
}

func formatBytes(bytes uint64) string {
	const (
		KB = 1 << 10
		MB = 1 << 20
		GB = 1 << 30
	)

	switch {
	case bytes < KB:
		return fmt.Sprintf("%d B", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	case bytes < GB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	default:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	}
}
