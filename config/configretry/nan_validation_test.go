package configretry

import (
	"math"
	"testing"
)

func TestRandomizationFactorNaN(t *testing.T) {
	cfg := NewDefaultBackOffConfig()
	cfg.RandomizationFactor = math.NaN()

	if err := cfg.Validate(); err == nil {
		t.Fatal("NaN should be rejected")
	}
}
