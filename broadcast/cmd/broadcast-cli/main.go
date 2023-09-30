package main

import (
	"fmt"
	"os"

	broadcast "github.com/shehackedyou/streamkit/broadcast"

	cli "github.com/multiverse-os/cli"
)

func main() {
	show := &broadcast.Show{
		OBS:  broadcast.Connect(broadcast.DefaultConfig()["host"]),
		Name: broadcast.DefaultConfig()["name"],
	}

	// TODO: Need to create a disconnect function
	//defer client.Disconnect()

	cmd, initErrors := cli.New(cli.App{
		Name:        "broadcast-cli",
		Description: "broadcast command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Commands: cli.Commands(
			cli.Command{
				Name:        "scenes",
				Alias:       "s",
				Description: "list all obs scenes",
				Subcommands: cli.Commands(
					cli.Command{
						Name:        "preview",
						Alias:       "q",
						Description: "list obs preview scene",
						Action: func(c *cli.Context) error {
							show.GetSceneList()

							previewScene := show.GetPreviewScene()
							fmt.Printf("obs.Show.Scenes.GetPreviewScene(%v)\n", previewScene)

							return nil
						},
					},
					cli.Command{
						Name:        "program",
						Alias:       "p",
						Description: "list obs program scene",
						Action: func(c *cli.Context) error {
							show.GetSceneList()

							programScene := show.GetProgramScene()
							fmt.Printf("obs.Show.Scenes.GetProgramScene(%v)\n", programScene)

							return nil
						},
					},
					cli.Command{
						Name:        "transition",
						Alias:       "t",
						Description: "transition obs program scene",
						Action: func(c *cli.Context) error {
							fmt.Printf("obs.ListScenes() caching scenes locally...\n")

							endScene := show.Scenes.First()
							fmt.Printf("obs.Show.Scenes.First() => endScene(%v)\n", endScene)

							primaryScene := show.Scenes.Last()
							fmt.Printf("obs.Show.Scenes.Last() => primaryScene(%v)\n", primaryScene)

							ok, err := show.SceneTransition(primaryScene)
							if err != nil {
								panic(err)
							}
							fmt.Printf("did we transition?\n")
							fmt.Printf("we should return FALSE if primaryScene is already current!(%v)\n", ok)

							return nil
						},
					},
				),
				Action: func(c *cli.Context) error {
					scenes := show.GetSceneList()

					fmt.Printf("xxxxxxxxxxxxxxxxxx\n")
					scenes.YAML(0)
					fmt.Printf("xxxxxxxxxxxxxxxxxx\n")

					return nil
				},
			},
			cli.Command{
				Name:        "scene",
				Alias:       "c",
				Description: "scene from name",
				Flags: cli.Flags(
					cli.Flag{
						Name:  "name",
						Alias: "n",
						// TODO CLI FRAMEWORK - action fallthrough?
						Description: "Select the name of the scene",
					},
				),
				Subcommands: cli.Commands(
					cli.Command{
						Name:        "index",
						Alias:       "i",
						Description: "obtain item index from scene",
						Action: func(c *cli.Context) error {
							sceneName := c.Flag("name").String()
							if len(sceneName) == 0 {
								return fmt.Errorf("failed to provide scene name")
							}

							scene := show.Scene(sceneName)
							fmt.Printf("scene(%v) with index(%v)\n", scene, scene.Index)
							return nil
						},
					},
					cli.Command{
						Name:        "transition",
						Alias:       "t",
						Description: "transition to scene",
						Action: func(c *cli.Context) error {
							sceneName := c.Flag("name").String()
							if len(sceneName) == 0 {
								return fmt.Errorf("failed to provide scene name")
							}

							scene := show.Scene(sceneName)
							fmt.Printf("scene(%v) with index(%v)\n", scene, scene.Index)
							return nil
						},
					},
					cli.Command{
						Name:        "items",
						Alias:       "i",
						Description: "list all items of scene",
						Action: func(c *cli.Context) error {
							sceneName := c.Flag("name").String()
							if len(sceneName) == 0 {
								return fmt.Errorf("failed to provide scene name")
							}

							scene := show.Scene(sceneName)

							items := show.ListSceneItems(scene)

							items.YAML(2)

							fmt.Printf("scene(%v) with index(%v)\n", scene, scene.Index)
							return nil
						},
					},
				),
				Action: func(c *cli.Context) error {
					c.CLI.Log(" action of scene")
					sceneName := c.Flag("name").String()
					if len(sceneName) == 0 {
						return fmt.Errorf("failed to provide scene name")
					}

					scene := show.Scene(sceneName)
					// TODO
					// Our problem is now that this will work but we need a before action
					// to grab our scene on our subcommands
					if scene != nil {
						fmt.Printf("xxxxxxxxxxxxxxxxxx\n")
						fmt.Printf(" scene(%v) with name(%v)\n", scene, scene.Name)
						fmt.Printf("xxxxxxxxxxxxxxxxxx\n")
					}

					return nil
				},
			},
		),
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				c.CLI.Log("[onStart] action")
				show.GetSceneList()
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			OnExit: func(c *cli.Context) error {
				c.CLI.Log("[onExit] action")
				fmt.Printf("scenes parsed?(%v)\n", len(show.Scenes))
				show.YAML(0)

				return nil
			},
		},
	})
	fmt.Printf("initErrors(%v)\n", initErrors)
	fmt.Printf("len(initErrors)(%v)\n", len(initErrors))
	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
