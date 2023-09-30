package broadcast

import (
	"fmt"
	"strings"
)

// NOTE
// Collection functions
type Items []*Item

func EmptyItems() Items { return make([]*Item, 0) }

func (is Items) IsEmpty() bool    { return len(is) == 0 }
func (is Items) IsNotEmpty() bool { return !is.IsEmpty() }

func (is Items) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sitems:\n", prefix)
	for _, item := range is {
		item.YAML(spaces + 2)
	}
}

type Item struct {
	Scene *Scene

	Id    float64
	Index float64
	Name  string
	Type  SourceType

	Parent *Item
	Group  Items

	// TODO
	// Going to need a way to move up to show and scene without
	// creating a loop of imports
}

func EmptyItem() *Item {
	return &Item{
		Scene: nil,
		Id:    -1,
		Index: -1,
		Name:  "",
		Type:  UndefinedType,
	}
}

func (i Item) Hide() bool {
	return i.Scene.Show.HideItem(i.Scene.Name, i.Id)
}

func (i Item) Unhide() bool {
	return i.Scene.Show.UnhideItem(i.Scene.Name, i.Id)
}

func (i Item) Lock() bool {
	return i.Scene.Show.LockItem(i.Scene.Name, i.Id)
}

func (i Item) Unlock() bool {
	return i.Scene.Show.UnlockItem(i.Scene.Name, i.Id)
}

func (i Item) IsGroup() bool        { return i.Type == GroupType }
func (i Item) IsUndefined() bool    { return i.Type == UndefinedType }
func (i Item) IsNotUndefined() bool { return !i.IsUndefined() }
func (i *Item) IsNil() bool         { return i == nil }
func (i *Item) IsNotNil() bool      { return !i.IsNil() }

// Aliasing
func (i Item) IsFolder() bool { return i.IsGroup() }
func (i Item) IsEmpty() bool  { return i.IsUndefined() }

func (i *Item) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sitem:\n", prefix)
	prefix = strings.Repeat(" ", spaces+2)
	fmt.Printf("%sindex: %v\n", prefix, i.Index)
	fmt.Printf("%sid: %v\n", prefix, i.Id)
	fmt.Printf("%ssource_name: %v\n", prefix, i.Name)
	fmt.Printf("%ssource_type: %v\n", prefix, i.Type.String())
	if i.Parent.IsNotNil() {
		fmt.Printf("%sparent_group: %v\n", prefix, i.Parent.Name)
	}
	if i.TypeIs(GroupType) {
		fmt.Printf("%sgrouped_items:\n", prefix)
		prefix = strings.Repeat(" ", spaces+4)
		fmt.Printf("%sgrouped_count: %v\n", prefix, len(i.Group))
		i.Group.YAML(spaces + 6)
	}
}

func (i *Item) TypeIs(st SourceType) bool { return i.Type == st }

type SourceType uint8

const (
	UndefinedType SourceType = iota
	InputType
	SceneType // This is a fucking folder not a scene ffs
)

// Aliasing to make things less stupid
const (
	GroupType = SceneType
)

func (st SourceType) String() string {
	switch st {
	case InputType:
		return "OBS_SOURCE_TYPE_INPUT"
	case SceneType:
		return "OBS_SOURCE_TYPE_SCENE"
	default:
		return ""
	}
}

func MarshalSourceType(st string) SourceType {
	switch st {
	case InputType.String():
		return InputType
	case SceneType.String():
		return SceneType
	default:
		return UndefinedType
	}
}
