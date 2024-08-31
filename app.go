package main

import (
	"fmt"

	"github.com/rivo/tview"
)

type CliApplication interface {
	Run() error
}

type cliApplication struct {
	app *tview.Application
}

func NewCliApplication() (CliApplication, error) {

	return &cliApplication{
		app: tview.NewApplication(),
	}, nil
}

func (c *cliApplication) Run() error {

	mainMenu, err := c.createMainMenu()
	if err != nil {
		return fmt.Errorf("failed to create main menu: %w", err)
	}

	if err := c.app.SetRoot(mainMenu, true).Run(); err != nil {
		return fmt.Errorf("failed to run application: %w", err)
	}

	return nil
}

func (c *cliApplication) createMainMenu() (*tview.List, error) {

	mainMenu := tview.NewList()
	mainMenu.AddItem("Infra Services", "Manage infrastructure services", '1', func() {
		menu, err := c.createInfraMenu()
		if err != nil {
			return
		}
		c.app.SetRoot(menu, true)
	})

	mainMenu.AddItem("Solidus Services", "Manage application services", '2', func() {
		menu, err := c.createSolidusMenu()
		if err != nil {
			return
		}
		c.app.SetRoot(menu, true)
	})

	mainMenu.AddItem("Exit", "Exit the application", 'q', func() {
		c.app.Stop()
	})

	return mainMenu, nil
}

func (c *cliApplication) createInfraMenu() (*tview.List, error) {

	infraMenu := tview.NewList()

	for service := range 10 {
		infraMenu.AddItem(fmt.Sprintf("service-%d", service), "Status: ", 0, nil)
	}

	infraMenu.AddItem("Back", "Return to main menu", 'b', func() {
		menu, err := c.createMainMenu()
		if err != nil {
			return
		}
		c.app.SetRoot(menu, true)
	})

	return infraMenu, nil
}

func (c *cliApplication) createSolidusMenu() (*tview.List, error) {

	solidusMenu := tview.NewList()

	for service := range 10 {
		solidusMenu.AddItem(fmt.Sprintf("service-%d", service), "Status: ", 0, nil)
	}

	solidusMenu.AddItem("Back", "Return to main menu", 'b', func() {
		menu, err := c.createMainMenu()
		if err != nil {
			return
		}
		c.app.SetRoot(menu, true)
	})

	return solidusMenu, nil
}
