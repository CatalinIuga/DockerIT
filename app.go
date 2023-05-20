package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Ping() string {
	// ping docker
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "error"
	}
	ping, err := cli.Ping(a.ctx)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%s!", ping.APIVersion)
}

// get container list
func (a *App) GetContainerList() string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "error"
	}
	ping, err := cli.Ping(a.ctx)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("Hello %s, It's show time!", ping.APIVersion)

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
