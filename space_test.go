package confluence

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	goconfluence "github.com/virtomize/confluence-go-api"
	"os"
	"testing"
)

var cf = NewConfluence("https://co-jesh.atlassian.net/wiki/rest/api", "cho911115@gmail.com", os.Getenv("TOKEN"))

func TestSpaces(t *testing.T) {
	spaces, err := cf.client.GetAllSpaces(goconfluence.AllSpacesQuery{})
	errHandler(err)

	for i, space := range spaces.Results {
		fmt.Printf("%d번째 : %s\n", i, space.Name)
	}

	assert.NotEqual(t, spaces, nil)
}
