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

func TestCT_QueryTableConstructor(t *testing.T) {
	v := spreadsheetml.NewCT_QueryTable()
	if v == nil {
		t.Errorf("spreadsheetml.NewCT_QueryTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetml.CT_QueryTable should validate: %s", err)
	}
}

func TestCT_QueryTableMarshalUnmarshal(t *testing.T) {
	v := spreadsheetml.NewCT_QueryTable()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetml.NewCT_QueryTable()
	xml.Unmarshal(buf, v2)
}