package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTemplatesList(t *testing.T) {
	templatesList = []string{"template_test"}
	expected := []string{"template_test"}

	assert.Equal(t, expected, GetTemplatesList())
}
