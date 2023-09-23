package broadcast

import (
	"fmt"

	obs "github.com/shehackedyou/streamkit/broadcast/obs"
)

// TODO: obs folder is primarily legacy working obs interaction
type OBS struct {
	*obs.Broadcast
	Show *Show
}

func DefaultConfig() map[string]string {
	return obs.DefaultConfig()
}

func Connect(host string) *obs.Broadcast {
	return &obs.Broadcast{
		Client: obs.Connect(host),
	}
}

// TODO
// So when we need to access the client to load our show objects
// we want to do it at this level to eventually obsolete the obs
// folder

// TODO
// So this works, next we want to load the items...
// cleanup and fix it to be better later just get functionality
// needed:
//  1. transition scenes
//  2. unhide/hide items within scenes
//
// Thats what is all remaining needed for most basic producerbot
func (o *OBS) ListScenes() {
	resp, _ := o.Client.Scenes.GetSceneList()
	for _, s := range resp.Scenes {
		fmt.Printf("%2d %s\n", s.SceneIndex, s.SceneName)
		o.Show.ParseScene(s.SceneIndex, s.SceneName)
	}
}
