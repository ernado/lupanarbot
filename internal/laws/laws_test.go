package laws

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomArticle(t *testing.T) {
	article, err := RandomArticle()
	require.NoError(t, err)
	require.NotEmpty(t, article.Title)
	t.Logf("Random article: %s - %s", article.Title, article.Text)
}
