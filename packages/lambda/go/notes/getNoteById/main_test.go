package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllNotes(t *testing.T) {
	note := GetNoteById("id1")
	require.NotEmpty(t, note)
	require.Equal(t, "id1", note["noteId"])
	note2 := GetNoteById("id2")
	require.NotEmpty(t, note2)
	require.Equal(t, "id2", note2["noteId"])
	note3 := GetNoteById("id3")
	require.Empty(t, note3)

}
