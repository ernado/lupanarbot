package minust

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/binary"
	"encoding/json"
	"io"

	"github.com/klauspost/compress/gzip"
)

type Element struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

//go:embed "elements.json.gz"
var data []byte

var Elements []Element

func Random() Element {
	if len(Elements) == 0 {
		return Element{
			ID:    228,
			Title: "БД Не инициализирована",
		}
	}

	// Use crypto/rand to generate random id.
	totalElements := len(Elements)

	// Generate random index using crypto/rand
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Fallback to first element if crypto/rand fails
		return Elements[0]
	}

	// Convert bytes to uint32 and get index within bounds
	randomUint32 := binary.BigEndian.Uint32(randomBytes)
	// Use uint64 to avoid overflow on 32-bit systems
	randomIndex := int(uint64(randomUint32) % uint64(totalElements))

	return Elements[randomIndex]
}

func init() {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		panic("failed to unzip elements data: " + err.Error())
	}

	defer func() {
		_ = reader.Close()
	}()

	unzippedData, err := io.ReadAll(reader)
	if err != nil {
		panic("failed to read unzipped data: " + err.Error())
	}

	if len(unzippedData) == 0 {
		panic("unzipped data is empty")
	}

	if err := json.Unmarshal(unzippedData, &Elements); err != nil {
		panic("failed to unmarshal elements data: " + err.Error())
	}
}
