package cmd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateRecord(t *testing.T) {
	tempFilePath := filepath.Join(t.TempDir(), "out.cbor")
	cmd := NewCreateCmd()

	args := []string{"--data=../testdata/cbor-in-1.cbor", "--output=" + tempFilePath}
	cmd.SetArgs(args)

	err := cmd.Execute()
	assert.NoError(t, err)

	actual, err := os.ReadFile(tempFilePath)
	assert.NoError(t, err)
	expected, err := os.ReadFile("../testdata/expected-record.cbor")
	assert.NoError(t, err)

	assert.Equal(t, expected, actual, "Output record did not have expected contents")
}

func TestCreateCollection(t *testing.T) {
	tempFilePath := filepath.Join(t.TempDir(), "out.cbor")
	cmd := NewCreateCmd()

	args := []string{"--data=../testdata/cbor-in-1.cbor", "--label=cbor1", "--data=../testdata/cbor-in-2.cbor", "--label=cbor2", "--output=" + tempFilePath}
	cmd.SetArgs(args)

	err := cmd.Execute()
	assert.NoError(t, err)

	actual, err := os.ReadFile(tempFilePath)
	assert.NoError(t, err)
	expected, err := os.ReadFile("../testdata/expected-collection.cbor")
	assert.NoError(t, err)

	assert.Equal(t, expected, actual, "Output collection did not have expected contents")
}
