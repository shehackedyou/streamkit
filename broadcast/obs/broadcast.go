package obs

import (
	goobs "github.com/andreykaipov/goobs"
)

type Client *goobs.Client

type Broadcast struct {
	Client *goobs.Client
	Mode   []Mode
	Scene  *Scene
	//Studio  *Scene
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"name": "she hacked you",
		"host": "10.100.100.1:4444",
	}
}

func (bc *Broadcast) ListScenes() {

}

//func (bc *Broadcast) Cache() bool {
//	// NOTE: For simplicity, for now we will just set scenes to empty and then
//	// populate with API data. So we set show.Scenes to an empty slice of scenes
//	apiScenesResponse, err := bc.Client.Scenes.GetSceneList()
//	if err != nil {
//		fmt.Printf("error(%v)\n", err)
//	}
//	fmt.Printf("apiScenesResponse(%v)", apiScenesResponse)
//	// apiResponse == type GetSceneListResponse struct {
//	// 	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`
//	// 	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`
//	// 	Scenes []*typedefs.Scene `json:"scenes,omitempty"`
//	// }
//
//	// TODO: Instead of iterating over the obs scene names and then using that
//	// to look it up, lets just see if we can iterate over the scenes in the API
//	// and then append the parsed version of those to append
//
//	// TODO: Its a way,... right?? ?? hello ?
//	//obsSceneNames := []string{}
//	//cachedSceneNames := []string{}
//	//for _, cachedScene := range sh.Scenes {
//	//	cachedSceneNames = append(cachedSceneNames, cachedScene.Name)
//	//}
//	fmt.Printf("show:\n")
//	fmt.Printf("  scenes:\n")
//	//
//	//for _, scene := range apiScenesResponse.Scenes {
//	//	apiResponse, err := bc.Client.SceneItems.GetSceneItemList(
//	//		&sceneitems.GetSceneItemListParams{
//	//			SceneName: scene.SceneName,
//	//		})
//	//	if err != nil {
//	//		fmt.Printf("error(%v)\n", err)
//	//	}
//
//	//	fmt.Printf("      scene:\n")
//	//	fmt.Printf("        object: %v\n", scene)
//	//	//fmt.Printf("        item_count: %v\n", len(apiResponse.SceneItems))
//	//	fmt.Printf("        items:\n")
//	//	fmt.Printf("          object: %v\n", apiResponse)
//	//	for _, sceneItem := range apiResponse.SceneItems {
//	//		fmt.Printf("          item: \n")
//	//		fmt.Printf("            object: %v\n", sceneItem)
//	//	}
//	//}
//	return true
//}
