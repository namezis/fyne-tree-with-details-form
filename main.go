package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"myexamples.org/tree-with-details-form/objects"
)

type Program struct {
	App              fyne.App
	MainWin          fyne.Window
	RootConfig       *objects.Root
	MainTree         *widget.Tree
	DetailsContainer *fyne.Container
}

var program Program

func main() {
	program.App = app.New()

	program.MainWin = program.App.NewWindow("Tree Example")
	program.RootConfig = objects.NewRoot()

	for i := 0; i < 220; i++ {
		info := program.RootConfig.Infos.NewInfo(fmt.Sprintf("sample%d", i))
		info.Param1 = fmt.Sprintf("text of sample%d", i)
		info.Param2 = fmt.Sprintf("The plates will still shift sample%d", i)
		info.Param3 = fmt.Sprintf("and the clouds will still spew. sample%d", i)
		info.Param4 = fmt.Sprintf("The sun will slowly rise sample%d", i)
		info.Param5 = fmt.Sprintf("and the moon will follow too. sample%d", i)
		info.Param6 = fmt.Sprintf("Nature's first green is gold, sample%d", i)
		info.Param7 = fmt.Sprintf("Her hardest hue to hold. sample%d", i)
		info.Param8 = fmt.Sprintf("Her early leaf's a flower; sample%d", i)
		info.Param9 = fmt.Sprintf("But only so an hour. sample%d", i)
		info.Param10 = fmt.Sprintf("Time doesn’t heal wounds sample%d", i)
		info.Param11 = fmt.Sprintf("to make you forget. sample%d", i)
		info.Param12 = fmt.Sprintf("It doesn’t heal wounds too. sample%d", i)
		info.Param13 = fmt.Sprintf("The plates will still shift. sample%d", i)
		info.Param14 = true
		info.Param15 = false
	}

	// Set up tree
	program.MainTree = &widget.Tree{}

	// Tree callbacks
	program.MainTree.ChildUIDs = func(uid string) (c []string) {
		c = program.RootConfig.GetChildIds(uid)
		return
	}
	program.MainTree.IsBranch = func(uid string) (b bool) {
		b = program.RootConfig.IsBranchBranch(uid)
		return
	}
	program.MainTree.CreateNode = func(branch bool) fyne.CanvasObject {
		return widget.NewLabel("Template")
	}
	program.MainTree.UpdateNode = func(uid string, branch bool, node fyne.CanvasObject) {
		caption := program.RootConfig.GetNodeName(uid)
		label := node.(*widget.Label)
		label.SetText(caption)
	}
	program.MainTree.ExtendBaseWidget(program.MainTree)

	program.MainTree.OnSelected = func(id string) {
		var selectedObjects []fyne.CanvasObject
		switch program.RootConfig.GetType(id) {
		case objects.ObjectType_Info:
			node := program.RootConfig.GetInfoNode(id)
			if node != nil {
				canvas_object := node.GetDetailsForm(id)
				selectedObjects = []fyne.CanvasObject{canvas_object}
			}
		case objects.ObjectType_Infos:
			node := program.RootConfig.GetInfosNode(id)
			if node != nil {
				canvas_object := node.GetDetailsForm(id)
				selectedObjects = []fyne.CanvasObject{canvas_object}
			}
		case objects.ObjectType_Setting:
			canvas_object := program.RootConfig.Settings.Get_Settings_Form(id)
			selectedObjects = []fyne.CanvasObject{canvas_object}
		default:
			selectedObjects = []fyne.CanvasObject{widget.NewLabel("No selection")}
		}
		program.DetailsContainer.Objects = selectedObjects
		program.DetailsContainer.Refresh()
	}
	program.MainTree.OnUnselected = func(id string) {
		program.DetailsContainer.Objects = []fyne.CanvasObject{widget.NewLabel("No selection")}
		program.DetailsContainer.Refresh()
	}

	// Root node callbakcs
	program.RootConfig.OnRefreshTree = func(id string) {
		program.MainTree.Refresh()
		if id != "" {
			program.MainTree.OpenBranch(id)
		}
	}
	program.RootConfig.OnHideDetailsForm = func() {
		program.DetailsContainer.Objects = []fyne.CanvasObject{widget.NewLabel("No selection")}
		program.DetailsContainer.Refresh()
	}

	// Main form setup
	program.DetailsContainer = container.NewMax(widget.NewLabel("No selection"))
	middle := container.NewHSplit(
		program.MainTree,
		program.DetailsContainer,
	)
	middle.SetOffset(0.4)

	status := widget.NewLabel("Application loaded.")

	mainContent := container.NewBorder(nil, status, nil, nil, middle)

	program.MainWin.SetContent(mainContent)
	program.MainWin.SetFixedSize(false)
	program.MainWin.Resize(fyne.NewSize(1000, 800))
	program.MainWin.SetMaster()
	program.MainWin.ShowAndRun()
}
