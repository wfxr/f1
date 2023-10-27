package run

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type DurationPercentileMap map[float64]time.Duration

func (m *DurationPercentileMap) String(indent int) string {
	s := ""
	keys := make([]float64, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for _, percentile := range keys {
		p := fmt.Sprintf("p%.0f", percentile*100)
		s = fmt.Sprintf("%s\n%s%4s: %.1fms", s, strings.Repeat(" ", indent), p, (*m)[percentile].Seconds()*1000)
	}
	return s
}

func (m *DurationPercentileMap) Get(pc float64) string {
	return fmt.Sprintf("%.1fms", (*m)[pc].Seconds()*1000)
}
