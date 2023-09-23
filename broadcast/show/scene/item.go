package scene

// NOTE
// Collection functions
type Items []*Item

func EmptyItems() Items { return make([]*Item, 0) }

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

func EmptyItem() *Item {
	return &Item{
		Id:    -1,
		Index: -1,
		Name:  "",
		Type:  UndefinedType,
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
