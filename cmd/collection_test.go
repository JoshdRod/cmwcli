package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
	//"github.com/spf13/afero" TODO: Re-include afero to allow test generated files to be removed automatically
	"os"
)

func TestCreateRecord(t *testing.T) {
	cmd := NewCollectionCmd()

	args := []string{"--data=../testdata/cbor-in-1.cbor", "--output=out.cbor"}
	cmd.SetArgs(args)

	err := cmd.Execute()
	assert.NoError(t, err)

	actual, err := os.ReadFile("out.cbor")
	assert.NoError(t, err)
	expected, err := os.ReadFile("../testdata/expected-record.cbor")
	assert.NoError(t, err)

	assert.Equal(t, expected, actual, "Output record did not have expected contents")
}

func TestCreateCollection(t *testing.T) {
	cmd := NewCollectionCmd()

	args := []string{"--data=../testdata/cbor-in-1.cbor", "--label=cbor1", "--data=../testdata/cbor-in-2.cbor", "--label=cbor2", "--output=out.cbor"}
	cmd.SetArgs(args)


	err := cmd.Execute()
	assert.NoError(t, err)

	actual, err := os.ReadFile("out.cbor")
	assert.NoError(t, err)
	expected, err := os.ReadFile("../testdata/expected-collection.cbor")
	assert.NoError(t, err)

	assert.Equal(t, expected, actual, "Output collection did not have expected contents")
}
