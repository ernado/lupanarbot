package laws

import (
	"crypto/rand"
	_ "embed"
	"encoding/json"
	"fmt"
	"math/big"
)

var (
	//go:embed articles.json
	articlesJSON []byte

	//go:embed constitution.json
	constitutionJSON []byte
)

func init() {
	// Initialize the law collection from embedded JSON data
	lawCollection, err := NewLawCollection(articlesJSON)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize law collection: %v", err))
	}
	fmt.Printf("Law collection initialized with %d articles.\n", lawCollection.GetArticleCount())
}

func RandomArticle() (*Article, error) {
	// Get a random article from the law collection
	lawCollection, err := NewLawCollection(articlesJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to create law collection: %w", err)
	}
	return lawCollection.GetRandomArticle()
}

func RandomConstitutionArticle() (*Article, error) {
	// Get a random article from the constitution
	lawCollection, err := NewLawCollection(constitutionJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to create constitution collection: %w", err)
	}
	return lawCollection.GetRandomArticle()
}

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type LawCollection struct {
	Articles []Article
}

// NewLawCollection creates a new collection from JSON data
func NewLawCollection(jsonData []byte) (*LawCollection, error) {
	var articles []Article
	if err := json.Unmarshal(jsonData, &articles); err != nil {
		return nil, fmt.Errorf("failed to unmarshal articles: %w", err)
	}

	return &LawCollection{Articles: articles}, nil
}

// GetRandomArticle returns a random article using crypto/rand
func (lc *LawCollection) GetRandomArticle() (*Article, error) {
	if len(lc.Articles) == 0 {
		return nil, fmt.Errorf("no articles available")
	}

	// Generate cryptographically secure random number
	maxNumber := big.NewInt(int64(len(lc.Articles)))
	n, err := rand.Int(rand.Reader, maxNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random number: %w", err)
	}

	index := n.Int64()
	return &lc.Articles[index], nil
}

// GetArticleByTitle finds an article by its title
func (lc *LawCollection) GetArticleByTitle(title string) *Article {
	for _, article := range lc.Articles {
		if article.Title == title {
			return &article
		}
	}
	return nil
}

// GetArticleCount returns the total number of articles
func (lc *LawCollection) GetArticleCount() int {
	return len(lc.Articles)
}
