package objects

import "strings"

type Root struct {
	// private
	treeId            string
	OnRefreshTree     func(id string)
	OnHideDetailsForm func()
	infoId            int
	infoListId        int
	// public
	Name     string
	Settings *SettingsList
	Infos    *InfoList
}

type ObjectType int

const (
	ObjectType_Unknown = iota
	ObjectType_Root
	ObjectType_Setting
	ObjectType_Settings
	ObjectType_Info
	ObjectType_Infos
)

func (o *Root) nextInfoId() int {
	o.infoId++
	return o.infoId
}

func (o *Root) nextInfoListId() int {
	o.infoListId++
	return o.infoListId
}

func (o *Root) get_child_ids() []string {
	return []string{o.Infos.treeId, o.Settings.treeId}
}

func (o *Root) get_settings_child_ids(id string) []string {
	var result []string
	if o.Settings.treeId == id {
		result = append(result, o.Settings.get_child_ids()...)
	}
	return result
}

func (o *Root) get_infos_child_ids(id string) []string {
	if o.Infos.treeId == id {
		return o.Infos.get_child_ids()
	}
	return []string{}
}

func (o *Root) get_info_child_ids(id string) []string {
	for _, info := range o.Infos.Nodes {
		if info.treeId == id {
			return info.get_child_ids()
		}
	}
	return []string{}
}

func (o *Root) get_info_name(id string) string {
	for _, info := range o.Infos.Nodes {
		if info.treeId == id {
			return info.String()
		}
	}
	return "Unknown"
}

func (o *Root) GetType(id string) ObjectType {

	if id == "" {
		return ObjectType_Root
	} else if strings.HasPrefix(id, "info:") {
		return ObjectType_Info
	} else if strings.HasPrefix(id, "infos:") {
		return ObjectType_Infos
	} else if strings.HasPrefix(id, "setting:") {
		return ObjectType_Setting
	} else if strings.HasPrefix(id, "settings:") {
		return ObjectType_Settings
	}
	return ObjectType_Unknown
}

func (o *Root) GetChildIds(id string) []string {
	switch o.GetType(id) {

	case ObjectType_Root:
		return o.get_child_ids()

	case ObjectType_Info:
		return o.get_info_child_ids(id)

	case ObjectType_Infos:
		return o.get_infos_child_ids(id)

	case ObjectType_Settings:
		return o.get_settings_child_ids(id)

	default:
		return []string{}
	}
}

func (o *Root) GetNodeName(id string) string {
	switch o.GetType(id) {

	case ObjectType_Root:
		return o.Name

	case ObjectType_Info:
		return o.get_info_name(id)

	case ObjectType_Infos:
		return "Infos"

	case ObjectType_Settings:
		return "Settings"

	case ObjectType_Setting:
		return o.Settings.get_name(id)

	default:
		return "Unknown"
	}
}

func (o *Root) IsBranchBranch(id string) bool {
	switch o.GetType(id) {

	case ObjectType_Root:
		return true

	case ObjectType_Info:
		return false

	case ObjectType_Infos:
		infos := o.GetInfosNode(id)
		if infos != nil {
			return infos.NodeCount() != 0
		} else {
			return false
		}

	case ObjectType_Settings:
		return true

	default:
		return false
	}
}

func (o *Root) GetInfoNode(id string) *Info {
	for i := range o.Infos.Nodes {
		if o.Infos.Nodes[i].treeId == id {
			return &o.Infos.Nodes[i]
		}
	}
	return nil
}

func (o *Root) GetInfosNode(id string) *InfoList {
	if o.Infos.treeId == id {
		return o.Infos
	}
	return nil
}

func (o *Root) RefreshTree(id string) {
	o.OnRefreshTree(id)
}

func (o *Root) SetupTree() {

	o.infoId = 0
	o.infoListId = 0

	o.treeId = "root"

	o.Settings.set_child_ids()

	o.Infos.treeRoot = o
	o.Infos.setup_node()
	for i := range o.Infos.Nodes {
		info := &o.Infos.Nodes[i]
		info.treeRoot = o
		info.setup_node()
	}
}

func NewRoot() *Root {
	o := &Root{
		Name: "Root",
	}
	o.Settings = NewSettingsList(o)
	o.Infos = NewInfoList("Infos", o)
	return o
}
