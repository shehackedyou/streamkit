package broadcast

import (
	"fmt"
	"strings"

	show "github.com/shehackedyou/streamkit/broadcast/show"
)

// TODO: We will need to restructure the goobs from our show object since it
// could just as easily be ffmpeg of output of window that is assembled from
// v42l and whatever else

// Doing this disentanglement has a lot of cascading benefits that may almost
// certainly missed if not thought long abuot it

type Show struct {
	//Id   int
	// Season []*Season
	// Episodes []*Episode
	//StudioScene *show.Scene
	Name        string
	ActiveScene *show.Scene
	Scenes      show.Scenes
}

func OpenShow(name string) *Show {
	return &Show{
		Name:        name,
		ActiveScene: EmptyScene(),
		//StudioScene: EmptyScene(),
		Scenes: make([]*show.Scene, 0),
	}
}

func EmptyScene() *show.Scene {
	return &show.Scene{Name: ""}
}

func (sh Show) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("show:\n")
	prefix = strings.Repeat(" ", spaces+2)
	fmt.Printf("%sname: %v\n", prefix, sh.Name)
	fmt.Printf("%sactive_scene:\n", prefix)
	sh.ActiveScene.YAML(spaces + 4)
	sh.Scenes.YAML(spaces + 4)
}

func (sh *Show) Scene(name string) *show.Scene {
	for _, scene := range sh.Scenes {
		if scene.Name == name {
			return scene
		}
	}
	return nil
}

func (sh *Show) ParseScene(index int, name string) *show.Scene {
	// Validate name & index
	var err error
	if !(0 < len(name) && len(name) < 255) &&
		!(0 <= index && index < 999) {
		panic(err)
	}

	parsedScene := &show.Scene{
		Index: index,
		Name:  name,
	}

	// TODO
	// We need to be checking if the scene has already been parsed
	// otherwise we are going to have a ton of duplicates and if
	// we catch it here we can avoid headaches
	sh.Scenes = append(sh.Scenes, parsedScene)

	return parsedScene
}

// GOOBS TYPEDEF
//
//	type Scene struct {
//		SceneIndex int    `json:"sceneIndex"`
//		SceneName  string `json:"sceneName"`
//	}
