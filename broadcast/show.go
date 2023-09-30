package broadcast

import (
	"fmt"
	"strings"
)

type Show struct {
	OBS    OBS
	Name   string
	Scenes Scenes

	// Id       int
	// Season   []*Season
	// Episodes []*Episode

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

// NOTE
// Since items required to have a unique name even across scenes
func (sh *Show) Item(name string) *Item {
	for _, scene := range sh.Scenes {
		if item := scene.Items.Name(name); item != nil {
			return item
		}
	}
	return nil
}

func (sh *Show) ParseScene(index int, name string) *Scene {
	// Validate name & index
	var err error
	if !(0 < len(name) && len(name) < 255) &&
		!(0 <= index && index < 999) {
		fmt.Printf("err(%v)", err)
		return nil
	}

	// NOTE
	// Prevent Duplicates
	if scene := sh.Scene(name); scene != nil {
		return scene
	}

	parsedScene := &Scene{
		Show:  sh,
		Index: index,
		Name:  name,
	}

	sh.Scenes = append(sh.Scenes, parsedScene)

	return parsedScene
}
