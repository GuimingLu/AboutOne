// Copyright (c) 2019-present AboutOne, Inc. All Rights Reserved.


package app

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
	"za-white-screen/model"
)

type OpenProvider struct {
	JoinProvider
}

const (
	CMD_OPEN = "open"
)

func init() {
	RegisterCommandProvider(&OpenProvider{})
}

func (open *OpenProvider) GetTrigger() string {
	return CMD_OPEN
}

func (open *OpenProvider) GetCommand(a *App, T goi18n.TranslateFunc) *model.Command {
	cmd := open.JoinProvider.GetCommand(a, T)
	cmd.Trigger = CMD_OPEN
	cmd.DisplayName = T("api.command_open.name")
	return cmd
}
