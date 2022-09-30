package config

import "time"

const (
	FPS = 50.0
	// FPS               = 50.0
	DT                = 1 / FPS
	SendTickerSeconds = 0.01
	EVICTION_INTERVAL = 5 * time.Second
)
