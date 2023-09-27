package main

import (
	"fmt"
	"os"

	broadcast "github.com/shehackedyou/streamkit/broadcast"

	cli "github.com/multiverse-os/cli"
)

func main() {
	obs := &broadcast.OBS{
		Broadcast: broadcast.Connect(broadcast.DefaultConfig()["host"]),
		Show:      broadcast.OpenShow("she hacked you"),
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
							obs.ListScenes()

							previewScene := obs.GetPreviewScene()
							fmt.Printf("obs.Show.Scenes.GetPreviewScene(%v)\n", previewScene)

							return nil
						},
					},
					cli.Command{
						Name:        "program",
						Alias:       "p",
						Description: "list obs program scene",
						Action: func(c *cli.Context) error {
							obs.ListScenes()

							programScene := obs.GetProgramScene()
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
							obs.ListScenes()

							endScene := obs.Show.Scenes.First()
							fmt.Printf("obs.Show.Scenes.First() => endScene(%v)\n", endScene)

							primaryScene := obs.Show.Scenes.Last()
							fmt.Printf("obs.Show.Scenes.Last() => primaryScene(%v)\n", primaryScene)

							ok, err := obs.SceneTransition(primaryScene)
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
					obs.ListScenes()

					fmt.Printf("xxxxxxxxxxxxxxxxxx\n")
					fmt.Printf(" this is what runs scenes and gives us a list\n")
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
						Name:        "id",
						Alias:       "i",
						Description: "obtain item id from scene",
						Action: func(c *cli.Context) error {
							c.CLI.Log("test")
							fmt.Printf("how many flags(%v) and object flags(%v)\n", len(c.Flags), c.Flags)
							sceneName := c.Flag("name").String()
							if len(sceneName) == 0 {
								fmt.Printf("error: failed to provide scene name\n")
								return fmt.Errorf("failed to provide scene name")
							}

							scene := obs.Show.Scene(sceneName)
							fmt.Printf("scene(%v) with name(%v)\n", scene, sceneName)

							fmt.Printf("xxxxxxxxxxxxxxxxxx\n")

							return nil
						},
					},
				),
				Action: func(c *cli.Context) error {
					c.CLI.Log("action of scene")
					sceneName := c.Flag("name").String()
					if len(sceneName) == 0 {
						return fmt.Errorf("failed to provide scene name")
					}

					scene := obs.Show.Scene(sceneName)
					// TODO
					// Our problem is now that this will work but we need a before action
					// to grab our scene on our subcommands
					fmt.Printf("xxxxxxxxxxxxxxxxxx\n")
					fmt.Printf("scene(%v) with name(%v)\n", scene, sceneName)
					fmt.Printf("xxxxxxxxxxxxxxxxxx\n")

					return nil
				},
			},
		),
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				c.CLI.Log("onStart action")

				obs.ListScenes()

				fmt.Printf("how many flags (%v)\n", len(c.Flags))

				fmt.Printf("first flag? (%v)\n", c.Flags.First())
				fmt.Printf("last flag? (%v)\n", c.Flags.Last())

				for index, fl := range c.Flags {
					fmt.Printf("flag@index(%v)=value(%v)\n", index, fl)
				}

				sceneName := c.Flag("name").String()

				fmt.Printf("is the flag available this early?(%v)\n", sceneName)
				fmt.Printf("ran obs.ListScenes() to cache them\n")

				c.CLI.Log("broadcast ", fmt.Sprintf("%v", obs))
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			OnExit: func(c *cli.Context) error {
				c.CLI.Log("onExit action")

				fmt.Printf("scenes parsed?(%v)\n", len(obs.Show.Scenes))
				fmt.Printf("now lets iterate over OUR type of scene...\n\n")
				//obs.Show.YAML(0)

				return nil
			},
		},
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
