package main

import (
	"log/slog"
)

func logLevel(val string) (slog.Level, error) {
	var (
		level = slog.LevelInfo
	)

	if err := level.UnmarshalText([]byte(val)); err != nil {
		return level, err
	}

	return level, nil
}
