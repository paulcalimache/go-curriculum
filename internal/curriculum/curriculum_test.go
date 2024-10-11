package curriculum

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

var (
	data        = `firstname: John`
	invalidData = `invalidField: invalidValue`
)

func TestParse(t *testing.T) {
	fs := fstest.MapFS{
		"data.yaml": {Data: []byte(data)},
	}

	cv, err := Parse(fs, "data.yaml")

	assert.Equal(t, "John", cv.Firstname)
	assert.NoError(t, err)
}

func TestParseNotYamlFile(t *testing.T) {
	fs := fstest.MapFS{
		"data.json": {Data: []byte(data)},
	}

	cv, err := Parse(fs, "data.json")

	assert.Nil(t, cv)
	assert.Error(t, err)
}

func TestParseNonexistentFile(t *testing.T) {
	fs := fstest.MapFS{}

	cv, err := Parse(fs, "data.yaml")

	assert.Nil(t, cv)
	assert.Error(t, err)
}

func TestParseInvalidYaml(t *testing.T) {
	fs := fstest.MapFS{
		"data.yaml": {Data: []byte(invalidData)},
	}

	cv, err := Parse(fs, "data.yaml")

	assert.Empty(t, cv)
	assert.NoError(t, err)
}

func TestTempletize(t *testing.T) {
	var cv CV
	output, err := cv.Templetize("classic")

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
}

func TestTempletizeInvalidTemplate(t *testing.T) {
	var cv CV
	output, err := cv.Templetize("invalid_template")

	assert.Error(t, err)
	assert.Empty(t, output)
}
