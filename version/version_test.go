package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionDefault(t *testing.T) {
	expected_version := "0.0.0"
	assert.Equal(t, expected_version, Version)
}

func TestCommitDefault(t *testing.T) {
	expected_commit := "HEAD"
	assert.Equal(t, expected_commit, GitCommit)
}
