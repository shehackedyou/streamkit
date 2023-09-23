package broadcast

import (
	obs "github.com/shehackedyou/streamkit/broadcast/obs"
)

type OBS obs.Broadcast

func DefaultConfig() map[string]string {
	return map[string]string{
		"name": "she hacked you",
		"host": "10.100.100.1:4444",
	}
}

func Connect(host string) obs.Client {
	return obs.Connect(host)
}
