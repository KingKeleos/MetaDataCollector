package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "testing correct port",
			want: 42000,
		},
	}
	for _, tt := range tests {
		fmt.Print("Starting the server on Port 4200?")
		assert.Equal(t, port, tt.want)
	}
}
