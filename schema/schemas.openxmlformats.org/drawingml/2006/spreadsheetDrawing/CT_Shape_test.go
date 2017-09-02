// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Shape()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}