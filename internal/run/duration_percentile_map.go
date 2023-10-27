package run

import (
	"fmt"
	"strings"
	"time"
)

type DurationPercentileMap map[float64]time.Duration

func (m *DurationPercentileMap) String(indent int) string {
	var s string
	for _, pc := range []float64{0.5, 0.75, 0.9, 0.95, 0.99, 1.0} {
		p := fmt.Sprintf("p%.0f", pc * 100)
		s += strings.Repeat(" ", indent) + fmt.Sprintf("%4s: %s\n", p, m.Get(pc))
	}

	return s
}

func (m *DurationPercentileMap) Get(pc float64) string {
	return fmt.Sprintf("%.1fms", (*m)[pc].Seconds()*1000)
}
