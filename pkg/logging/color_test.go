package logging

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestColorHandler_Handle(t *testing.T) {
	// Simple test to ensure Handler runs without error
	h := &ColorHandler{}
	
	t.Run("Pass message", func(t *testing.T) {
		r := slog.NewRecord(time.Now(), slog.LevelInfo, "[PASS]", 0)
		if err := h.Handle(context.Background(), r); err != nil {
			t.Errorf("Handle() failed: %v", err)
		}
	})

	t.Run("Fail message", func(t *testing.T) {
		r := slog.NewRecord(time.Now(), slog.LevelInfo, "[FAIL]", 0)
		if err := h.Handle(context.Background(), r); err != nil {
			t.Errorf("Handle() failed: %v", err)
		}
	})
}
