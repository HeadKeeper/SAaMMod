package histogram

import (
	"github.com/aybabtme/uniplot/histogram"
	"os"
)

func DrawHistogram(values []float64) error {
	bins := 20
	maxWidth := 10
	hist := histogram.Hist(bins, values)
	err := histogram.Fprint(os.Stdout, hist, histogram.Linear(maxWidth))
	return err
}
