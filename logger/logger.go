package logger

import (
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"log/slog"
	"os"
)

// isJetbrains attempts to detect GoLand from env vars since the IDE won't look
// like a terminal. When you run a Go program (or any program) from within an
// Integrated Development Environment (IDE) like PyCharm, the program's
// input/output is typically not directly connected to a terminal. Instead, it's
// connected to a pipe or some other form of redirection that the IDE uses to
// capture the program's output and display it in its own output pane.
func isJetbrains() bool {
	_, golandPresent := os.LookupEnv("GoLand")
	_, ideaDirPresent := os.LookupEnv("IDEA_INITIAL_DIRECTORY")
	return golandPresent || ideaDirPresent
}

// NewLogger creates a new instance of a slog.Logger depending on the given configuration. For
// GitHub Actions, it will use plain text logging. In a terminal it will use colorized text, and
// on a server it will output using JSON for structured logging.
func NewLogger() *slog.Logger {
	w := os.Stderr
	jetbrains := isJetbrains()
	terminal := isatty.IsTerminal(w.Fd())
	if terminal || jetbrains {
		return slog.New(
			tint.NewHandler(colorable.NewColorable(w), nil),
		)
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
