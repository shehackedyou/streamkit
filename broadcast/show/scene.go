package show

import (
	"fmt"
	"strings"

	scene "github.com/shehackedyou/streamkit/broadcast/show/scene"
)

// NOTE
// This is our collection type for creating methods based on a collection
// of scenes within a show
type Scenes []*Scene

func (scs Scenes) IsEmpty() bool    { return len(scs) == 0 }
func (scs Scenes) IsNotEmpty() bool { return !scs.IsEmpty() }

// TODO
// Honestly not even sure if this makes sense since slices don't maintain order
func (scs Scenes) First() *Scene {
	if scs.IsNotEmpty() {
		return scs[0]
	}
	return nil
}

func (scs Scenes) Last() *Scene {
	if scs.IsNotEmpty() {
		return scs[(len(scs) - 1)]
	}
	return nil
}

func (scs Scenes) Scene(name string) *Scene {
	for _, sc := range scs {
		if sc.Name == name {
			return sc
		}
	}
	return nil
}

func (scs Scenes) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sscenes:\n", prefix)
	for _, sc := range scs {
		prefix = strings.Repeat(" ", spaces+2)
		fmt.Printf("%sscene:\n", prefix)
		prefix = strings.Repeat(" ", spaces+4)
		fmt.Printf("%sindex: %v\n", prefix, sc.Index)
		fmt.Printf("%sname: %v\n", prefix, sc.Name)
		if sc.Items.IsNotEmpty() {
			sc.Items.YAML(spaces + 6)
		}
	}
}

type Scene struct {
	Index int
	Name  string
	Items scene.Items
}

// TODO: These make more sense as we build in more validation on any and every
// input because we are writing secure software (sorta; ok not really)
func (sc *Scene) IsNil() bool    { return sc == nil }
func (sc *Scene) IsNotNil() bool { return !sc.IsNil() }

func (sc Scene) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sscene:\n", prefix)
	prefix = strings.Repeat(" ", spaces+2)
	fmt.Printf("%sindex: %v\n", prefix, sc.Index)
	fmt.Printf("%sname: %v\n", prefix, sc.Name)
	if sc.Items.IsNotEmpty() {
		sc.Items.YAML(spaces + 2)
	}
}

func (sc *Scene) HasName(name string) bool {
	return (sc != nil || len(sc.Name) != len(name) || len(name) == 0)
}

// NOTE
// This does work because OBS doesn't allow duplicate names; not
// even across scenes so technically could have all the items
// together too.
func (sc *Scene) Item(name string) *scene.Item {
	// TODO: Need to have a way to iterate over the items in the scene
	for _, item := range sc.Items {
		// TODO: Should we bother strings.ToLower() for each?
		if item.Name == name {
			return item
		}
	}
	return nil
}

func (sc *Scene) ParseItem(id, index int, iType, name string) *scene.Item {
	var err error
	// NOTE
	// Validate for each of the types so we only generate valid objects
	// and 128 is probably too long for iType but works for a placeholder
	if !(0 < len(name) && len(name) < 255) &&
		!(0 < len(iType) && len(iType) < 128) &&
		!(0 <= index && index < 999) &&
		!(0 <= id && id < 999) {
		// TODO: If we are failing to parse an item we have big problems;
		// especially after all this validation
		panic(err)
	}

	parsedItem := &scene.Item{
		Id:    id,
		Index: index,
		Name:  name,
		Type:  scene.MarshalSourceType(iType),
	}

	// TODO
	// Need to prevent duplicates here, so we save ourselves from
	// tedious headache causing problems; simple as searching before
	// doing this append; or switching to a linked-list which may
	// make more sense and most people dont realize is in the stdlib
	sc.Items = append(sc.Items, parsedItem)

	return parsedItem
}
