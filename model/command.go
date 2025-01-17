


package model

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	COMMAND_METHOD_POST = "P"
	COMMAND_METHOD_GET  = "G"
	COMMAND_METHOD_NON  = "N"
	MIN_TRIGGER_LENGTH  = 1
	MAX_TRIGGER_LENGTH  = 128
)

type Command struct {
	Id               string `json:"id"`
	Token            string `json:"token"`
	CreateAt         int64  `json:"create_at"`
	UpdateAt         int64  `json:"update_at"`
	DeleteAt         int64  `json:"delete_at"`
	CreatorId        string `json:"creator_id"`
	TeamId           string `json:"team_id"`
	Trigger          string `json:"trigger"`
	Method           string `json:"method"`
	Username         string `json:"username"`
	IconURL          string `json:"icon_url"`
	AutoComplete     bool   `json:"auto_complete"`
	AutoCompleteDesc string `json:"auto_complete_desc"`
	AutoCompleteHint string `json:"auto_complete_hint"`
	DisplayName      string `json:"display_name"`
	Description      string `json:"description"`
	URL              string `json:"url"`
	CardType         string  `json:"card_type"`
	CommandType      string  `json:"command_type"`
}

//type Scaffold struct {
//	Account      string `json:"account"`
//	ArtifactId   string `json:"artifact_id"`
//	GroupId      string `json:"group_id"`
//	PackageName  string `json:"package_name"`
//	ProjectName  string `json:"project_name"`
//	ResoponseUrl string `json:"resoponse_url"`
//}
//
//type ScaffoldComponentBasicDTO struct {
//	ComponentName string `json:"component_name"`
//	ComponentValue string `json:"component_value"`
//}
//type ScaffoldComponent struct {
//	Account                   string                       `json:"account"`
//	ProjectId                 uint64                       `json:"project_id"`
//	ResoponseUrl              string                       `json:"resoponse_url"`
//	ScaffoldComponentBasicDTO []*ScaffoldComponentBasicDTO `json:"scaffold_component_basic_dto"`
//}

func (o *Command) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}
//func ScaffoldFromJson(data io.Reader) *Scaffold  {
//	var o *Scaffold
//	json.NewDecoder(data).Decode(&o)
//	return o
//}
//func ScaffoldComponentFromJson(data io.Reader) *ScaffoldComponent  {
//	var o *ScaffoldComponent
//	json.NewDecoder(data).Decode(&o)
//	return o
//}

func CommandFromJson(data io.Reader) *Command {
	var o *Command
	json.NewDecoder(data).Decode(&o)
	return o
}

func CommandListToJson(l []*Command) string {
	b, _ := json.Marshal(l)
	return string(b)
}

func CommandListFromJson(data io.Reader) []*Command {
	var o []*Command
	json.NewDecoder(data).Decode(&o)
	return o
}

func (o *Command) IsValid() *AppError {

	if len(o.Id) != 26 {
		return NewAppError("Command.IsValid", "model.command.is_valid.id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.Token) != 26 {
		return NewAppError("Command.IsValid", "model.command.is_valid.token.app_error", nil, "", http.StatusBadRequest)
	}

	if o.CreateAt == 0 {
		return NewAppError("Command.IsValid", "model.command.is_valid.create_at.app_error", nil, "", http.StatusBadRequest)
	}

	if o.UpdateAt == 0 {
		return NewAppError("Command.IsValid", "model.command.is_valid.update_at.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.CreatorId) != 26 {
		return NewAppError("Command.IsValid", "model.command.is_valid.user_id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.TeamId) != 26 {
		return NewAppError("Command.IsValid", "model.command.is_valid.team_id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.Trigger) < MIN_TRIGGER_LENGTH || len(o.Trigger) > MAX_TRIGGER_LENGTH || strings.Index(o.Trigger, "/") == 0 || strings.Contains(o.Trigger, " ") {
		return NewAppError("Command.IsValid", "model.command.is_valid.trigger.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.URL) == 0 || len(o.URL) > 1024 {
		return NewAppError("Command.IsValid", "model.command.is_valid.url.app_error", nil, "", http.StatusBadRequest)
	}

	if !IsValidHttpUrl(o.URL) {
		return NewAppError("Command.IsValid", "model.command.is_valid.url_http.app_error", nil, "", http.StatusBadRequest)
	}

	if !(o.Method == COMMAND_METHOD_GET || o.Method == COMMAND_METHOD_POST || o.Method == COMMAND_METHOD_NON) {
		return NewAppError("Command.IsValid", "model.command.is_valid.method.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.DisplayName) > 64 {
		return NewAppError("Command.IsValid", "model.command.is_valid.display_name.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.Description) > 128 {
		return NewAppError("Command.IsValid", "model.command.is_valid.description.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (o *Command) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	if o.Token == "" {
		o.Token = NewId()
	}

	o.CreateAt = GetMillis()
	o.UpdateAt = o.CreateAt
}

func (o *Command) PreUpdate() {
	o.UpdateAt = GetMillis()
}

func (o *Command) Sanitize() {
	o.Token = ""
	o.CreatorId = ""
	o.Method = ""
	o.URL = ""
	o.Username = ""
	o.IconURL = ""
}
