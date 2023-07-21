package core

import (
	"fmt"
	"time"
)

func MeasureTime(then time.Time, label string) {
	elapsed := time.Since(then)
	fmt.Printf("[%s] %s\n", label, elapsed)
}
