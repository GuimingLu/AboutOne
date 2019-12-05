// Copyright (c) 2019-present AboutOne, Inc. All Rights Reserved.


package model

import (
	"encoding/json"
	"io"
)

type MfaSecret struct {
	Secret string `json:"secret"`
	QRCode string `json:"qr_code"`
}

func (me *MfaSecret) ToJson() string {
	b, _ := json.Marshal(me)
	return string(b)
}

func MfaSecretFromJson(data io.Reader) *MfaSecret {
	var me *MfaSecret
	json.NewDecoder(data).Decode(&me)
	return me
}