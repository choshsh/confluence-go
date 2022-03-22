package confluence

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	goconfluence "github.com/virtomize/confluence-go-api"
	"testing"
)

func TestCreateContent(t *testing.T) {
	contentSearch, err := cf.client.GetContent(goconfluence.ContentQuery{
		Title:  "api-test-02",
		Type:   "page",
		Expand: []string{"version"},
	})
	errHandler(err)

	content := &goconfluence.Content{}

	if contentSearch.Size > 0 {
		content = cf.UpdateContent(contentSearch.Results[0])
	} else {
		content = cf.CreateContent()
	}

	fmt.Println(content.Links.Base + content.Links.TinyUI)
	assert.NotEqual(t, content, nil)
}
