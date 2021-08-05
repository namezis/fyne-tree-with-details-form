package objects

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"myexamples.org/tree-with-details-form/widgets"
)

type Info struct {
	// private
	treeId   string
	treeRoot *Root
	// public
	Name    string
	Param1  string
	Param2  string
	Param3  string
	Param4  string
	Param5  string
	Param6  string
	Param7  string
	Param8  string
	Param9  string
	Param10 string
	Param11 string
	Param12 string
	Param13 string
	Param14 bool
	Param15 bool
}

func (o *Info) String() string {
	return o.Name
}

func (o *Info) GetDetailsForm(id string) fyne.CanvasObject {

	boundParam1 := binding.BindString(&o.Param1)
	boundParam2 := binding.BindString(&o.Param2)
	boundParam3 := binding.BindString(&o.Param3)
	boundParam4 := binding.BindString(&o.Param4)
	boundParam5 := binding.BindString(&o.Param5)
	boundParam6 := binding.BindString(&o.Param6)
	boundParam7 := binding.BindString(&o.Param7)
	boundParam8 := binding.BindString(&o.Param8)
	boundParam9 := binding.BindString(&o.Param9)
	boundParam10 := binding.BindString(&o.Param10)
	boundParam11 := binding.BindString(&o.Param11)
	boundParam12 := binding.BindString(&o.Param12)
	boundParam13 := binding.BindString(&o.Param13)
	boundParam14 := binding.BindBool(&o.Param14)
	boundParam15 := binding.BindBool(&o.Param15)

	paramsCard := widget.NewCard(
		"Params",
		"",
		container.NewVBox(
			container.New(layout.NewFormLayout(),
				widget.NewLabel("Param1:"), widget.NewEntryWithData(boundParam1),
				widget.NewLabel("Param2:"), widget.NewEntryWithData(boundParam2),
				widget.NewLabel("Param3:"), widget.NewEntryWithData(boundParam3),
				widget.NewLabel("Param4:"), widget.NewEntryWithData(boundParam4),
				widget.NewSeparator(), widget.NewSeparator(),
				widget.NewLabel("Param5:"), widget.NewEntryWithData(boundParam5),
				widget.NewLabel("Param6:"), widget.NewEntryWithData(boundParam6),
				widget.NewLabel("Param7:"), widget.NewEntryWithData(boundParam7),
				widget.NewLabel("Param8:"), widget.NewEntryWithData(boundParam8),
				widget.NewLabel("Param15:"), widget.NewCheckWithData("", boundParam15),
				widget.NewSeparator(), widget.NewSeparator(),
				widget.NewLabel("Param9:"), widget.NewEntryWithData(boundParam9),
				widget.NewLabel("Param10:"), widget.NewEntryWithData(boundParam10),
				widget.NewLabel("Param11:"), widget.NewEntryWithData(boundParam11),
				widget.NewLabel("Param12:"), widget.NewEntryWithData(boundParam12),
				widget.NewLabel("Param13:"), widget.NewEntryWithData(boundParam13),
				widget.NewSeparator(), widget.NewSeparator(),
				widget.NewLabel("Param14:"), widget.NewCheckWithData("", boundParam14),
			),
		),
	)

	paramsContainer := container.NewHBox(
		container.NewVBox(
			widgets.NewContentRect(500),
			paramsCard,
		),
	)

	optionsCard := widget.NewCard(
		"Options",
		"",
		container.NewVBox(
			container.New(layout.NewFormLayout(),
				widget.NewLabel("Enable broker status:"), widget.NewLabel(""),
				widget.NewLabel("Broker status interval:"), widget.NewLabel(""),
				widget.NewLabel(""), widget.NewLabel(""),
				widget.NewLabel("Enable broker network info:"), widget.NewLabel(""),
				widget.NewLabel("Broker network info interval:"), widget.NewLabel(""),
				widget.NewLabel(""), widget.NewLabel(""),
			),
		),
	)

	optionsContainer := container.NewHBox(
		container.NewVBox(
			widgets.NewContentRect(400),
			optionsCard,
		),
	)

	output := container.NewAppTabs(
		container.NewTabItem("Params", paramsContainer),
		container.NewTabItem("Options", optionsContainer),
	)

	return output
}

func (o *Info) get_child_ids() []string {
	return []string{}
}

func (o *Info) setup_node() {
	o.treeId = fmt.Sprintf("info:%d_%s", o.treeRoot.nextInfoId(), strings.ReplaceAll(o.Name, "_", ""))
}

func NewInfo(name string, root *Root) *Info {
	o := &Info{
		Name:     name,
		treeRoot: root,
	}
	o.setup_node()
	return o
}
