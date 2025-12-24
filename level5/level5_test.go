package level5

import (
	"testing"
)

func TestCalculateRating(t *testing.T) {
	tests := []struct {
		name           string
		player         Player
		expectedRating float64
	}{
		{
			name:           "No misses",
			player:         NewPlayer("Player1", 10, 0, 4),
			expectedRating: 12.0,
		},
		{
			name:           "Normal case",
			player:         NewPlayer("Player2", 8, 2, 4),
			expectedRating: 5.0,
		},
		{
			name:           "Zero assists",
			player:         NewPlayer("Player3", 5, 1, 0),
			expectedRating: 5.0,
		},
		{
			name:           "All stats",
			player:         NewPlayer("Player4", 12, 3, 6),
			expectedRating: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.player.Rating != tt.expectedRating {
				t.Errorf("Expected rating %.2f, got %.2f", tt.expectedRating, tt.player.Rating)
			}
		})
	}
}
