package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllNotes(t *testing.T) {
	notes := GetAllNotes()
	require.Equal(t, "Hello World!", notes["id1"]["content"])
	require.Equal(t, "Hello Old World!", notes["id2"]["content"])

}
