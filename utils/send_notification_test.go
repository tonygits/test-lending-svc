package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendNotification(t *testing.T) {
	// define notification message
	message := "message sent"
	result := SendNotification(message)
	require.Equal(t, true, result)

	// define empty message
	result = SendNotification("")
	require.Equal(t, false, result)
}
