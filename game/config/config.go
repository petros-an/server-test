package config

import "time"

const (
	FPS = 50.0
	// FPS               = 50.0
	DT                = 1 / FPS
	SendTickerSeconds = 0.01
	EVICTION_INTERVAL = 20 * time.Second
)
