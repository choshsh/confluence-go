package confluence

import (
	goconfluence "github.com/virtomize/confluence-go-api"
	"strconv"
)

func (c *confluence) CreateContent() *goconfluence.Content {
	content := &goconfluence.Content{
		Title: "api-test-02",
		Type:  "page",
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          "#api-test\nnew sub\npage", // your page content here
				Representation: "editor2",
			},
		},
		Ancestors: []goconfluence.Ancestor{
			goconfluence.Ancestor{
				ID: "1015809",
			},
		},
		Space: goconfluence.Space{Key: "~166200948"},
	}

	content, err := c.client.CreateContent(content)
	errHandler(err)

	return content
}

func (c *confluence) UpdateContent(contentSearch goconfluence.Content) *goconfluence.Content {
	content := &goconfluence.Content{
		Title: "api-test-02",
		Type:  "page",
		ID:    contentSearch.ID,
		Body: goconfluence.Body{
			Storage: goconfluence.Storage{
				Value:          "#api-test\nnew sub\npage" + strconv.Itoa(contentSearch.Version.Number+1),
				Representation: "editor2",
			},
		},
		Ancestors: []goconfluence.Ancestor{
			goconfluence.Ancestor{
				ID: "1015809",
			},
		},
		Version: &goconfluence.Version{
			Number: contentSearch.Version.Number + 1,
		},
		Space: goconfluence.Space{Key: "~166200948"},
	}

	content, err := c.client.UpdateContent(content)
	errHandler(err)
	return content
}
