package confluence

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	data := NewTemplate()
	fmt.Println(data)
	assert.NotEqual(t, data, nil)
}
