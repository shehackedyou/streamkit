package obs

import (
	goobs "github.com/andreykaipov/goobs"
)

type Broadcast struct {
	Client *goobs.Client
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"name": "she hacked you",
		"host": "10.100.100.1:4444",
	}
}
