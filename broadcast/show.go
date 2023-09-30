package broadcast

import (
	"fmt"
	"strings"
)

type Show struct {
	OBS OBS
	//Id   int
	// Season []*Season
	// Episodes []*Episode
	Name   string
	Scenes Scenes
	// TODO
	// While the concept of creating managing or even automating scenes makes
	// sense here but we have to decide if the show stores the OBS concept of the
	// newly named ProgramScene (Active Scene) and Preview Scene (previously
	// poorly named Studio Scene). Honestly it could fit in ~~both~~.

	// But keep in mind we wanted the Show to be segregated from OBS. But the
	// concept of the scene especially OUR abstraction and datatype could easily
	// apply to a 2D engine if done correctly.
	ProgramScene *Scene
	PreviewScene *Scene
}

func (sh Show) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("show:\n")
	prefix = strings.Repeat(" ", spaces+2)
	fmt.Printf("%sname: %v\n", prefix, sh.Name)
	fmt.Printf("%sactive_scene:\n", prefix)
	if sh.ProgramScene.IsNotNil() {
		sh.ProgramScene.YAML(spaces + 4)
	}
	if sh.PreviewScene.IsNotNil() {
		sh.PreviewScene.YAML(spaces + 4)
	}
	sh.Scenes.YAML(spaces + 4)
}

func (sh *Show) Scene(name string) *Scene {
	for _, scene := range sh.Scenes {
		if scene.Name == name {
			return scene
		}
	}
	return nil
}

func (sh *Show) ParseScene(index int, name string) *Scene {
	// Validate name & index
	var err error
	if !(0 < len(name) && len(name) < 255) &&
		!(0 <= index && index < 999) {
		panic(err)
	}

	parsedScene := &Scene{
		Show:  sh,
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
