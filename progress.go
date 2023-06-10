package progressbar

import (
	"fmt"
	"strings"
	"time"

	"github.com/philhanna/commas"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Progress is a custom writer that tracks the progress of a
// long-running operation.
type Progress struct {
	total       int
	soFar       int
	lastPercent int
	sTime       time.Time
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewProgress creates a new Progress structure and returns a pointer to
// it.
func NewProgress(total int) *Progress {
	p := new(Progress)
	p.total = total
	p.soFar = 0
	p.lastPercent = 0
	p.sTime = time.Now()
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Add adds n to the soFar counter, and triggers a display if the
// percent complete is greater than the last percentage complete.
func (p *Progress) Add(n int) {
	p.soFar += n
	percent := int(float64(p.soFar * 100) / float64(p.total))
	if percent != p.lastPercent {
		s := strings.Repeat("*", percent/2)
		for len(s) < 50 {
			s += "."
		}
		if percent > p.lastPercent {
			countString := commas.Format(p.soFar)
			fmt.Printf("Percent complete: %d%%, [%-s] %s records in %v\r",
				percent, s, countString, time.Since(p.sTime))
		}
		p.lastPercent = percent
	}
}
