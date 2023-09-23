package broadcast

import (
	"fmt"

	"github.com/andreykaipov/goobs/api/requests/sceneitems"
	obs "github.com/shehackedyou/streamkit/broadcast/obs"
	show "github.com/shehackedyou/streamkit/broadcast/show"
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
	response, err := o.Client.Scenes.GetSceneList()
	if err != nil {
		panic(err)
	}
	for _, s := range response.Scenes {
		fmt.Printf("%2d %s\n", s.SceneIndex, s.SceneName)
		o.Show.ParseScene(s.SceneIndex, s.SceneName)
	}
}

func (o *OBS) ListSceneItems(scene *show.Scene) {
	// TODO
	// This type of shit where we are interacting with sceneitems or
	// typedefs we need to push that into obs but for now lets just
	// get working shit
	params := &sceneitems.GetSceneItemListParams{
		SceneName: scene.Name,
	}

	response, err := o.Client.SceneItems.GetSceneItemList(params)
	if err != nil {
		panic(err)
	}

	// TODO
	// We are absolutely dropping this shit, we are going to
	// just have them all be items underneath the scene, then
	// we are going to group the items in a group by using an
	// attribute. OR we just recursively nest them but itd be
	// most likely easier to work with if we just have them all
	// be in scene.Items then have Group("name").Items and just
	// hide the group item
	fmt.Printf("items for scene len(%v)\n ", len(response.SceneItems))
	for _, item := range response.SceneItems {
		fmt.Printf("item:\n")
		fmt.Printf("  id: %v\n", item.SceneItemID)
		fmt.Printf("  index: %v\n", item.SceneItemIndex)
		fmt.Printf("  source_type: %v\n", item.SourceType)

		// TODO
		// Eventually sort by type, ones that are scene_type
		// are better called "groups" but it doesn't even nest
		// more than 1 level but we should be storing it but
		// start with it
		if scene.MarshalSourceType(item.SourceType) == scene.GroupType {
			o.SceneGroupList(item.SourceName)
		}

		fmt.Printf("  source_name: %v\n", item.SourceName)
	}
}

func (o *OBS) SceneGroupList(groupName string) {
	params := &sceneitems.GetGroupSceneItemListParams{
		SceneName: groupName,
	}

	response, err := o.Client.SceneItems.GetGroupSceneItemList(params)
	if err != nil {
		panic(err)
	}

	for _, item := range response.SceneItems {
		fmt.Printf("  item_group:\n")
		fmt.Printf("    id: %v\n", item.SceneItemID)
		fmt.Printf("    index: %v\n", item.SceneItemIndex)
		fmt.Printf("    source_type: %v\n", item.SourceType)
		fmt.Printf("    source_name: %v\n", item.SourceName)
	}
}
