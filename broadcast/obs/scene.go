package obs

type Items []*Item

// TODO: Move scene content here
//type Scenes []*Scene
//
//func (scs Scenes) Size() int     { return len(scs) }
//func (scs Scenes) First() *Scene { return scs[0] }
//func (scs Scenes) Last() *Scene  { return scs[scs.Size()-1] }
//
//func (scs Scenes) IsEmpty() bool { return scs.Size() == 0 }
//
//// TODO: Add reverse to get order in the OBS GUI
//
//func (scs Scenes) Name(name string) (*Scene, bool) {
//	fmt.Sprintf("name we are looking for: %v\n", name)
//	for _, scene := range scs {
//		fmt.Sprintf("  scene: %v \n", scene)
//		if scene.HasName(name) {
//			return scene, true
//		}
//	}
//	return nil, false
//}

// TODO: Better track Current scene (its what is transitioned to last, we
//
//	ideally will have a log like system, or at least dates) and get
//	rid of this bool below bc its bad 4 rlly
//

//func (sc Scene) Item(name string) (*Item, bool) {
//	return sc.Items.Name(name)
//}
//
//func (sc Scene) ItemNameContains(searchText string) (*Item, bool) {
//	return sc.Items.NameContains(searchText)
//}
//
//func (sc *Scene) HasName(name string) bool {
//	return sc != nil && sc.Name == name
//}
//
//func (sc *Scene) Transition(sleepDuration ...time.Duration) (*Scene, bool) {
//	if 0 < len(sleepDuration) {
//		fmt.Printf("sleeping \n")
//		time.Sleep(sleepDuration[0])
//	}
//
//	_, err := sc.Broadcast.Client.Scenes.SetCurrentProgramScene(
//		&scenes.SetCurrentProgramSceneParams{
//			SceneName: sc.Name,
//		},
//	)
//
//	if err == nil {
//		sc.IsCurrent = true
//		sc.Broadcast.Scene = sc
//	}
//
//	return sc, err == nil
//}
