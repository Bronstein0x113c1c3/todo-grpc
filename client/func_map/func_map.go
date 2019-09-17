package func_map

import (
	"time"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func TimeElapsed(inputTime time.Time) time.Duration {
	return time.Since(inputTime)
}

func TimeElapsedTimestamp(inputTime *timestamp.Timestamp) time.Duration {
	t, err := ptypes.Timestamp(inputTime)
	
	if err != nil {
		return time.Duration(0)
	}

	trunc := time.Duration(time.Second)

	return TimeElapsed(t).Truncate(trunc)
}

func Mod(i, j int) bool {
	return i%j == 0
}