package objects

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"myexamples.org/tree-with-details-form/widgets"
)

type InfoList struct {
	// private
	treeId   string
	treeRoot *Root
	// public
	Name  string
	Nodes []Info
}

func (o *InfoList) NodeCount() int {
	return len(o.Nodes)
}

func (o *InfoList) NewInfo(name string) *Info {
	a := NewInfo(name, o.treeRoot)
	o.Nodes = append(o.Nodes, *a)
	return &o.Nodes[len(o.Nodes)-1]
}

var errInfoExists = fmt.Errorf("info already exists")
var errInfoNameRequired = fmt.Errorf("provide info name")

func (o *InfoList) GetDetailsForm(id string) fyne.CanvasObject {

	infoValidator := func(s string) (err error) {
		if s == "" {
			return errInfoNameRequired
		}
		for _, info := range o.Nodes {
			if info.Name == s {
				return errInfoExists
			}
		}
		return nil
	}

	name := widget.NewEntry()
	name.SetPlaceHolder("My new info name")
	name.Validator = infoValidator
	name.SetText("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "Info name"},
		},
	}
	form.SubmitText = "Add"
	form.OnSubmit = func() {
		if len(name.Text) > 0 {
			o.NewInfo(name.Text)
			o.treeRoot.RefreshTree(id)
			name.SetText("")
		}
	}

	addInfoCard := widget.NewCard(
		"Add info",
		"",
		container.NewVBox(
			form,
		),
	)

	manageContainer := container.NewHBox(
		container.NewVBox(
			widgets.NewContentRect(350),
			addInfoCard,
		),
	)

	output := container.NewAppTabs(
		container.NewTabItem("Manage List", manageContainer),
		container.NewTabItem("Params", widget.NewLabel("No option listeds")),
		container.NewTabItem("Options", widget.NewLabel("No option listeds")),
	)

	return output
}

func (o *InfoList) get_child_ids() []string {
	var result []string
	for _, obj := range o.Nodes {
		result = append(result, obj.treeId)
	}
	return result
}

func (o *InfoList) setup_node() {
	o.treeId = fmt.Sprintf("infos:%d_%s", o.treeRoot.nextInfoListId(), strings.ReplaceAll(o.Name, "_", ""))
}

func NewInfoList(name string, root *Root) *InfoList {
	o := &InfoList{
		Name:     name,
		treeRoot: root,
	}
	o.setup_node()
	return o
}
