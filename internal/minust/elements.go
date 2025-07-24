package minust

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"

	"github.com/klauspost/compress/gzip"
)

type Element struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//go:embed "elements.json.gz"
var data []byte

var Elements []Element

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
