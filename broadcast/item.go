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

func (is Items) First() *Item {
	if is.IsNotEmpty() {
		return is[0]
	}
	return nil
}

func (is Items) Last() *Item {
	if is.IsNotEmpty() {
		return is[len(is)-1]
	}
	return nil
}

func (is Items) Index(index float64) *Item {
	for _, i := range is {
		if i.Index == index {
			return i
		}
	}
	return nil
}

func (is Items) Id(id float64) *Item {
	for _, i := range is {
		if i.Id == id {
			return i
		}
	}
	return nil
}

func (is Items) Name(name string) *Item {
	for _, item := range is {
		if len(item.Name) == len(name) &&
			item.Name == name {
			return item
		}
		if item.TypeIs(GroupType) {
			for _, groupItem := range item.Group {
				if len(groupItem.Name) == len(name) &&
					groupItem.Name == name {
					return groupItem
				}
			}
		}
	}
	return nil
}

func (is Items) Groups() (groups Items) {
	for _, i := range is {
		if i.Type == GroupType {
			groups = append(groups, i)
		}
	}
	return groups
}

func (is Items) Parent(item *Item) (children Items) {
	for _, i := range is {
		if i.Parent.Id == item.Id {
			children = append(children, i)
		}
	}
	return children
}

func (is Items) YAML(spaces int) {
	prefix := strings.Repeat(" ", spaces)
	fmt.Printf("%sitems:\n", prefix)
	for _, item := range is {
		item.YAML(spaces + 2)
	}
}

// TODO: I get we have to pass it to obs as float64 but its a fucking integer
//
//	so why are we fighting with making it a float64 when we can do that
//	conversion later?
type Item struct {
	Scene  *Scene
	Parent *Item

	Id    float64
	Index float64
	Name  string
	Type  SourceType
	Group Items
}

func EmptyItem() *Item {
	return &Item{
		Scene:  nil,
		Parent: nil,
		Id:     -1,
		Index:  -1,
		Name:   "",
		Type:   UndefinedType,
		Group:  EmptyItems(),
	}
}

func (i *Item) ParseGroupItem(id, index float64, iType, name string) *Item {
	if !(0 < len(name) && len(name) < 255) &&
		!(0 < len(iType) && len(iType) < 128) &&
		!(0 <= index && index < 999) &&
		!(0 <= id && id < 999) ||
		i.Type != GroupType {
		// TODO
		// If we are failing to parse an item we have big problems;
		// especially after all this validation
		return i
	}

	parsedItem := &Item{
		Scene:  i.Scene,
		Parent: i,
		Id:     id,
		Index:  index,
		Name:   name,
		Type:   MarshalSourceType(iType),
	}

	i.Group = append(i.Group, parsedItem)

	return i
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
