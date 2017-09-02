// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/purl.org/dc/terms"
)

func TestDCMITypeConstructor(t *testing.T) {
	v := terms.NewDCMIType()
	if v == nil {
		t.Errorf("terms.NewDCMIType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.DCMIType should validate: %s", err)
	}
}

func TestDCMITypeMarshalUnmarshal(t *testing.T) {
	v := terms.NewDCMIType()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewDCMIType()
	xml.Unmarshal(buf, v2)
}