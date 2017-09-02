// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"
)

func TestCT_MapConstructor(t *testing.T) {
	v := spreadsheetml.NewCT_Map()
	if v == nil {
		t.Errorf("spreadsheetml.NewCT_Map must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetml.CT_Map should validate: %s", err)
	}
}

func TestCT_MapMarshalUnmarshal(t *testing.T) {
	v := spreadsheetml.NewCT_Map()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetml.NewCT_Map()
	xml.Unmarshal(buf, v2)
}