package broadcast

import (
	"fmt"

	goobs "github.com/andreykaipov/goobs"

	sceneitems "github.com/andreykaipov/goobs/api/requests/sceneitems"
	scenes "github.com/andreykaipov/goobs/api/requests/scenes"
)

// TODO
// NOW what needs to be done is simply this in broadcast module:
//   * Transition to a different scene
//   * Look up an item (even a sub-item, so consider maybe storing all item
//   pointers in the show or scene instead of nesting, otherwise the lookup
//   may be a pain. if we use the SAME pointer, changing one should correctly
//   update it in ALL places. failure to get this functionality means failureS
//   to properly implement it.
//   * Then lastly we need ability to hide and unhide items in scenes

// WITH THAT we can finally go back to streamkit and piece the two components
// together from broadcast and xserver and get our producerbot 100

// NOTE
// This is one of the clever ways to add methods to struct from other module
type OBS *goobs.Client

func OpenShow(name string) *Show {
	show := &Show{
		Name:         DefaultConfig()["name"],
		OBS:          Connect(DefaultConfig()["host"]),
		ProgramScene: EmptyScene(),
		PreviewScene: EmptyScene(),
		Scenes:       EmptyScenes(),
	}
	// TODO: Getting close to when we want to make this both return the scenes and
	// possibly call it cache scenes
	show.ProgramScene = show.GetProgramScene()
	show.PreviewScene = show.GetPreviewScene()
	show.Scenes = show.GetSceneList()

	return show
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"name": "she hacked you",
		"host": "10.100.100.1:4444",
	}
}

func Connect(host string) OBS {
	client, err := goobs.New(host)
	if err != nil {
		panic(err)
	}
	return client
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
func (show *Show) GetSceneList() Scenes {
	response, err := show.OBS.Scenes.GetSceneList()
	if err != nil {
		panic(err)
	}
	for _, scene := range response.Scenes {
		scene := show.ParseScene(scene.SceneIndex, scene.SceneName)
		scene.Items = show.ListSceneItems(scene)
	}
	return show.Scenes
}

func (show *Show) ListSceneItems(scene *Scene) Items {
	params := &sceneitems.GetSceneItemListParams{
		SceneName: scene.Name,
	}

	response, err := show.OBS.SceneItems.GetSceneItemList(params)
	if err != nil {
		panic(err)
	}

	for _, item := range response.SceneItems {
		parsedItem := scene.ParseItem(
			float64(item.SceneItemID),
			float64(item.SceneItemIndex),
			item.SourceType,
			item.SourceName,
		)

		if parsedItem.TypeIs(GroupType) {
			parsedItem.Group = show.GetGroupedItemList(scene, parsedItem)
		}
	}
	return scene.Items
}

// TODO: hrmm this might need to be on scene or we have to pass the scene object
// through if we want to parse it
// TODO: this one is troublesome because it has to use scene which means we
// should be moving these over to the actual objects
func (show *Show) GetGroupedItemList(scene *Scene, groupedItem *Item) Items {
	params := &sceneitems.GetGroupSceneItemListParams{
		SceneName: groupedItem.Name,
	}

	response, err := show.OBS.SceneItems.GetGroupSceneItemList(params)
	if err != nil {
		panic(err)
	}

	for _, item := range response.SceneItems {
		groupedItem.ParseGroupItem(
			float64(item.SceneItemID),
			float64(item.SceneItemIndex),
			item.SourceType,
			item.SourceName,
		)
	}
	return groupedItem.Group
}

func (show *Show) GetSceneItemId(sceneName string, offset float64, item string) float64 {
	params := &sceneitems.GetSceneItemIdParams{
		SceneName:    sceneName,
		SearchOffset: offset,
		SourceName:   item,
	}

	response, err := show.OBS.SceneItems.GetSceneItemId(params)
	if err != nil {
		return -1
	}

	return response.SceneItemId
}

func (show *Show) IsItemLocked(sceneName string, itemId float64) bool {
	params := &sceneitems.GetSceneItemLockedParams{
		SceneItemId: itemId,
		SceneName:   sceneName,
	}

	response, err := show.OBS.SceneItems.GetSceneItemLocked(params)
	if err != nil {
		fmt.Printf("err(%v)\n", err)
		return false
	}

	return response.SceneItemLocked
}

func (show *Show) IsItemVisible(sceneName string, itemId float64) bool {
	params := &sceneitems.GetSceneItemEnabledParams{
		SceneItemId: itemId,
		SceneName:   sceneName,
	}

	response, err := show.OBS.SceneItems.GetSceneItemEnabled(params)
	if err != nil {
		fmt.Printf("err(%v)\n", err)
		return false
	}

	return response.SceneItemEnabled
}

func (show *Show) SetItemVisibility(sceneName string, itemId float64, visible bool) bool {
	params := &sceneitems.SetSceneItemEnabledParams{
		SceneItemEnabled: &visible,
		SceneItemId:      itemId,
		SceneName:        sceneName,
	}

	// NOTE Response is literally empty, so dumb
	_, err := show.OBS.SceneItems.SetSceneItemEnabled(params)
	return err != nil
}

func (show *Show) HideItem(sceneName string, itemId float64) bool {
	return show.SetItemVisibility(sceneName, itemId, false)
}

func (show *Show) UnhideItem(sceneName string, itemId float64) bool {
	return show.SetItemVisibility(sceneName, itemId, true)
}

func (show *Show) SetItemLocked(sceneName string, itemId float64, locked bool) bool {
	params := &sceneitems.SetSceneItemLockedParams{
		SceneName:       sceneName,
		SceneItemId:     itemId,
		SceneItemLocked: &locked,
	}

	_, err := show.OBS.SceneItems.SetSceneItemLocked(params)
	return err != nil
}

func (show *Show) LockItem(sceneName string, itemId float64) bool {
	return show.SetItemLocked(sceneName, itemId, true)
}

func (show *Show) UnlockItem(sceneName string, itemId float64) bool {
	return show.SetItemLocked(sceneName, itemId, false)
}

func (show *Show) GetSceneItemIndex(sceneName string, itemId float64) float64 {
	params := &sceneitems.GetSceneItemIndexParams{
		SceneName:   sceneName,
		SceneItemId: itemId,
	}

	response, err := show.OBS.SceneItems.GetSceneItemIndex(params)
	if err != nil {
		return -1
	}

	return response.SceneItemIndex
}

func (show *Show) SetSceneItemIndex(
	sceneName string,
	itemId float64,
	itemIndex float64,
) bool {
	params := &sceneitems.SetSceneItemIndexParams{
		SceneName:      sceneName,
		SceneItemId:    itemId,
		SceneItemIndex: itemIndex,
	}

	_, err := show.OBS.SceneItems.SetSceneItemIndex(params)
	return err != nil
}

// NOTE
// There is an ENTIRE section on transitions now completely segregated from
// scenes and the scene setcurrentprogramscene which is a transition
// Transitions let you change duration, style, etc; but obviously we want to
// combine that back together into a scene object; we start here, get our logic
// working and then migrate the majority of the specific logic to the places it
// should go-- first priority producerbot100
// NEXT UP item hiding and unhiding! and I GUESS looking up a specific item,
// going to need to do that I GUESS if we want to control said ITEM
func (show *Show) GetProgramScene() *Scene {
	params := &scenes.GetCurrentProgramSceneParams{}

	response, err := show.OBS.Scenes.GetCurrentProgramScene(params)
	if err != nil {
		return nil
	}

	programSceneName := response.CurrentProgramSceneName
	if len(programSceneName) == 0 {
		return nil
	}

	return show.Scene(programSceneName)
}

func (show *Show) GetPreviewScene() *Scene {
	params := &scenes.GetCurrentPreviewSceneParams{}

	response, err := show.OBS.Scenes.GetCurrentPreviewScene(params)
	if err != nil {
		return nil
	}

	previewSceneName := response.CurrentPreviewSceneName
	if len(previewSceneName) == 0 {
		return nil
	}

	return show.Scene(previewSceneName)
}

func (show *Show) IsStudioMode() bool {
	params := &scenes.GetCurrentPreviewSceneParams{}

	response, err := show.OBS.Scenes.GetCurrentPreviewScene(params)
	if err != nil {
		return false
	}

	previewSceneName := response.CurrentPreviewSceneName
	return len(previewSceneName) != 0
}

func (show *Show) SceneTransition(scene *Scene) (bool, error) {
	if show.ProgramScene.Index == scene.Index {
		return false, fmt.Errorf("scene already program scene")
	}

	params := &scenes.SetCurrentProgramSceneParams{
		SceneName: scene.Name,
	}

	_, err := show.OBS.Scenes.SetCurrentProgramScene(params)
	if err != nil {
		return false, err
	}

	show.ProgramScene = scene
	return true, nil
}
