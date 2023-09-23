package scene

import (
	"fmt"
	"strings"
)

// NOTE
// Collection functions
type Items []*Item

func (its Items) IsEmpty() bool    { return len(its) == 0 }
func (its Items) IsNotEmpty() bool { return !its.IsEmpty() }

func EmptyItems() Items { return make([]*Item, 0) }

func (its Items) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sitems:\n", prefix)
	for _, item := range its {
		item.YAML(spaces + 2)
	}
}

type Item struct {
	Id    int
	Index int
	Name  string
	Type  SourceType

	Parent *Item
	Group  Items

	// TODO
	// Going to need a way to move up to show and scene without
	// creating a loop of imports
}

func (i Item) IsUndefined() bool    { return i.Type == UndefinedType }
func (i Item) IsNotUndefined() bool { return !i.IsUndefined() }
func (i *Item) IsNil() bool         { return i == nil }
func (i *Item) IsNotNil() bool      { return !i.IsNil() }

func EmptyItem() *Item {
	return &Item{
		Id:    -1,
		Index: -1,
		Name:  "",
		Type:  UndefinedType,
	}
}

func (i *Item) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sitem:\n", prefix)
	prefix = strings.Repeat(" ", spaces+2)
	fmt.Printf("%sid: %v\n", prefix, i.Id)
	fmt.Printf("%sindex: %v\n", prefix, i.Index)
	fmt.Printf("%ssource_type: %v\n", prefix, i.Type.String())
	fmt.Printf("%ssource_name: %v\n", prefix, i.Name)
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
