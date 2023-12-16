package time_track

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

func Track(logger *log.Helper, funcName string, start time.Time) {
	elapsed := time.Since(start).Seconds()
	logger.Infof("%s took %.2fs", funcName, elapsed)
}
