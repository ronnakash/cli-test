package main

import (
	"fmt"

	// "github.com/gdamore/tcell/v2"
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

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu, _ := c.createSolidusMenu()
	main, _ := c.createMainMenu()
	sideBar, _ := c.createInfraMenu()

	// menu.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorPink))
	// main.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorPink))
	// sideBar.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorPink))

	grid := tview.NewGrid().
		SetRows(1, -1, -1, 1).
		SetColumns(-1, -2).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 4, 0, 0, false).
		AddItem(newPrimitive("Footer"), 3, 0, 1, 4, 0, 0, false)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(sideBar, 2, 0, 1, 1, 0, 0, false).
		AddItem(main, 1, 1, 2, 3, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(menu).Run(); err != nil {
		panic(err)
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
