package objects

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"myexamples.org/tree-with-details-form/widgets"
)

const settings_list_id = "settings:program"

type SettingsGeneral struct {
	// private
	treeId string
	// public
	DebugLevel string
	Param1     string
	Param2     bool
}

const setting_general_id = "setting:general"

type SettingsList struct {
	// private
	treeId   string
	treeRoot *Root
	// public
	Name    string
	General SettingsGeneral
}

func (o *SettingsList) String() string {
	return o.Name
}

func (o *SettingsList) get_child_ids() []string {
	return []string{o.General.treeId}
}

func (o *SettingsList) get_name(id string) string {
	switch id {
	case setting_general_id:
		return "General"
	default:
		return "Unknown"
	}
}

func (o *SettingsList) Get_Settings_Form(id string) fyne.CanvasObject {
	if id == setting_general_id {
		return o.get_general_form()
	}
	return widget.NewLabel("No selection")
}

func (o *SettingsList) get_general_form() fyne.CanvasObject {

	boundParam1 := binding.BindString(&o.General.Param1)
	boundParam2 := binding.BindBool(&o.General.Param2)

	paramsCard := widget.NewCard(
		"Params",
		"",
		container.NewVBox(
			container.New(layout.NewFormLayout(),
				widget.NewLabel("Param1:"), widget.NewEntryWithData(boundParam1),
				widget.NewLabel("Param2:"), widget.NewCheckWithData("", boundParam2),
			),
		),
	)

	container := container.NewHBox(
		container.NewVBox(
			widgets.NewContentRect(400),
			paramsCard,
			layout.NewSpacer(),
		),
	)

	return container
}

func (o *SettingsList) set_child_ids() {
	o.treeId = settings_list_id
	o.General.treeId = setting_general_id
}

func NewSettingsList(root *Root) *SettingsList {

	o := &SettingsList{
		Name:     "Application Settings",
		treeRoot: root,
	}
	o.set_child_ids()
	return o
}
