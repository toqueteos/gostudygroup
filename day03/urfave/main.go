package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "fake-docker"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "APP_LANG",
		},
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "Load configuration from `FILE`",
			EnvVar: "APP_CONFIG_FILE",
		},
	}

	app.Commands = []cli.Command{
		cmd("attach", "a", "Attach local standard input, output, and error streams to a running container"),
		cmd("build", "b", "Build an image from a Dockerfile"),
		cmd("commit", "", "Create a new image from a container's changes"),
		cmd("cp", "", "Copy files/folders between a container and the local filesystem"),
		cmd("create", "", "Create a new container"),
		cmd("diff", "d", "Inspect changes to files or directories on a container's filesystem"),
		cmd("events", "", "Get real time events from the server"),
		cmd("exec", "", "Run a command in a running container"),
		cmd("export", "", "Export a container's filesystem as a tar archive"),
		cmd("history", "h", "Show the history of an image"),
		cmd("images", "i", "List images"),
		cmd("import", "", "Import the contents from a tarball to create a filesystem image"),
		cmd("info", "", "Display system-wide information"),
		cmd("inspect", "", "Return low-level information on Docker objects"),
		cmd("kill", "k", "Kill one or more running containers"),
		cmd("load", "", "Load an image from a tar archive or STDIN"),
		cmd("login", "", "Log in to a Docker registry"),
		cmd("logout", "", "Log out from a Docker registry"),
		cmd("logs", "l", "Fetch the logs of a container"),
		cmd("pause", "", "Pause all processes within one or more containers"),
		cmd("port", "", "List port mappings or a specific mapping for the container"),
		cmd("ps", "", "List containers"),
		cmd("pull", "", "Pull an image or a repository from a registry"),
		cmd("push", "", "Push an image or a repository to a registry"),
		cmd("rename", "", "Rename a container"),
		cmd("restart", "", "Restart one or more containers"),
		cmd("rm", "", "Remove one or more containers"),
		cmd("rmi", "", "Remove one or more images"),
		cmd("run", "r", "Run a command in a new container"),
		cmd("save", "", "Save one or more images to a tar archive (streamed to STDOUT by default)"),
		cmd("search", "", "Search the Docker Hub for images"),
		cmd("start", "", "Start one or more stopped containers"),
		cmd("stats", "", "Display a live stream of container(s) resource usage statistics"),
		cmd("stop", "", "Stop one or more running containers"),
		cmd("tag", "", "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
		cmd("top", "", "Display the running processes of a container"),
		cmd("unpause", "", "Unpause all processes within one or more containers"),
		cmd("update", "", "Update configuration of one or more containers"),
		cmd("version", "", "Show the Docker version information"),
		cmd("wait", "", "Block until one or more containers stop, then print their exit codes"),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func cmd(name, alias, usage string) cli.Command {
	return cli.Command{
		Name:    name,
		Aliases: []string{alias},
		Usage:   usage,
		Action: func(c *cli.Context) error {
			fmt.Printf("called %q with args %v\n", name, c.Args())
			return nil
		},
	}
}
