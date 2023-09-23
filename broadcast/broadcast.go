package broadcast

import (
	"fmt"

	obs "github.com/shehackedyou/streamkit/broadcast/obs"
	show "github.com/shehackedyou/streamkit/broadcast/show"
	scene "github.com/shehackedyou/streamkit/broadcast/show/scene"

	sceneitems "github.com/andreykaipov/goobs/api/requests/sceneitems"
)

// TODO
// NOW what needs to be done is simply this in broadcast module:
//   * Transition to a different scene
//   * Look up an item (even a sub-item, so consider maybe storing all item
//   pointers in the show or scene instead of nesting, otherwise the lookup
//   may be a pain. if we use the SAME pointer, changing one should correctly
//   update it in ALL places. failure to get this functionality means failure
//   to properly implement it.
//   * Then lastly we need ability to hide and unhide items in scenes

// WITH THAT we can finally go back to streamkit and piece the two components
// together from broadcast and xserver and get our producerbot 100

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
		parsedScene := o.Show.ParseScene(s.SceneIndex, s.SceneName)

		o.ListSceneItems(parsedScene)
	}
}

func (o *OBS) ListSceneItems(parsedScene *show.Scene) {
	// TODO
	// This type of shit where we are interacting with sceneitems or
	// typedefs we need to push that into obs but for now lets just
	// get working shit
	params := &sceneitems.GetSceneItemListParams{
		SceneName: parsedScene.Name,
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
		parsedItem := parsedScene.ParseItem(
			item.SceneItemID,
			item.SceneItemIndex,
			item.SourceType,
			item.SourceName,
		)

		// TODO
		// Eventually sort by type, ones that are scene_type
		// are better called "groups" but it doesn't even nest
		// more than 1 level but we should be storing it but
		// start with it
		if parsedItem.TypeIs(scene.GroupType) {
			groupedItems := o.SceneGroupList(parsedScene, parsedItem)

			parsedItem.Group = groupedItems
		}

	}
}

// TODO: hrmm this might need to be on scene or we have to pass the scene object
// through if we want to parse it
func (o *OBS) SceneGroupList(
	parsedScene *show.Scene,
	itemGroup *scene.Item,
) (groupedItems scene.Items) {
	params := &sceneitems.GetGroupSceneItemListParams{
		SceneName: itemGroup.Name,
	}

	response, err := o.Client.SceneItems.GetGroupSceneItemList(params)
	if err != nil {
		panic(err)
	}

	for _, item := range response.SceneItems {
		parsedGroupedItem := parsedScene.ParseItem(
			item.SceneItemID,
			item.SceneItemIndex,
			item.SourceType,
			item.SourceName,
		)

		parsedGroupedItem.Parent = itemGroup
		groupedItems = append(groupedItems, parsedGroupedItem)
	}
	return groupedItems
}
