package samplescraper

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTimeoutSecondsNaN(t *testing.T) {
	cfg := newDefaultTargetsItem()

	require.NoError(t, cfg.Validate())

	cfg.TimeoutSeconds = math.NaN()

	err := cfg.Validate()

	require.ErrorContains(t, err,
		"timeout_seconds value must be greater than 0")

	require.ErrorContains(t, err,
		"timeout_seconds value must be less than 30")
}
