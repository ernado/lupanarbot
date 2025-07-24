package minust

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElements(t *testing.T) {
	require.NotEmpty(t, Elements)
}

func TestRandom(t *testing.T) {
	r := Random()
	require.NotEmpty(t, r.Title)
	t.Logf("Random element: %d - %s", r.ID, r.Title)
}
