package findupes

import (
	"testing"
)

func TestNewResult(t *testing.T) {
	t.Run("NewResult", func(t *testing.T) {
		src := "123"
		match := "345"
		acc := Accuracy("Test")

		result := NewResult(src, match, acc)
		if result.SourceID != src && result.MatchID != match && result.Accuracy != acc {
			t.Errorf("expected %v %v %v, got %v", src, match, acc, result)
		}
	})

	t.Run("Result.Export", func(t *testing.T) {
		export := Result{
			SourceID: "123",
			MatchID:  "456",
			Accuracy: Accuracy("Test"),
		}.Export()
		expected := []string{"123", "456", "Test"}

		for i, val := range expected {
			if export[i] != val {
				t.Errorf("expected %v, got %vd", val, export[i])
			}
		}
	})
}

func TestResults(t *testing.T) {
	t.Run("Results.Export", func(t *testing.T) {
		results := Results{
			{
				SourceID: "123",
				MatchID:  "456",
				Accuracy: Accuracy("Test"),
			},
			{
				SourceID: "789",
				MatchID:  "012",
				Accuracy: Accuracy("Test2"),
			},
		}.Export()

		if len(results) != 2 {
			t.Fatalf("expected length %d, got %d", 2, len(results))
		}
	})
}
