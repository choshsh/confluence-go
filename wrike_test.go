package confluence

import (
	"fmt"
	"github.com/choshsh/wrike-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWrike(t *testing.T) {
	wrikeAPI := wrike.NewWrike(os.Getenv("BEARER"), nil)
	sprints := wrikeAPI.Sprints("2022.03.SP1")

	fmt.Printf("%+v\n", sprints)
	assert.NotEqual(t, sprints, nil)
}
