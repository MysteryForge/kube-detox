package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

type ColorHandler struct{}

func (h *ColorHandler) Handle(_ context.Context, r slog.Record) error {
	var msg string
	var attrs []string
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a.Key+"="+a.Value.String())
		return true
	})

	switch r.Message {
	case "[PASS]":
		msg = "\033[32m[PASS]\033[0m"
	case "[FAIL]":
		msg = "\033[31m[FAIL]\033[0m"
	case "[WARN]":
		msg = "\033[33m[WARN]\033[0m"
	default:
		msg = r.Message
	}

	if len(attrs) > 0 {
		_, _ = os.Stdout.WriteString(msg + " " + strings.Join(attrs, " ") + "\n")
	} else {
		_, _ = os.Stdout.WriteString(msg + "\n")
	}
	return nil
}

func (h *ColorHandler) Enabled(_ context.Context, _ slog.Level) bool { return true }
func (h *ColorHandler) WithAttrs(_ []slog.Attr) slog.Handler         { return h }
func (h *ColorHandler) WithGroup(_ string) slog.Handler              { return h }
