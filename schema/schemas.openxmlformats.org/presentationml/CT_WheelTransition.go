// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_WheelTransition struct {
	// Spokes
	SpokesAttr *uint32
}

func NewCT_WheelTransition() *CT_WheelTransition {
	ret := &CT_WheelTransition{}
	return ret
}
func (m *CT_WheelTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.SpokesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spokes"},
			Value: fmt.Sprintf("%v", *m.SpokesAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WheelTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spokes" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpokesAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WheelTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_WheelTransition) Validate() error {
	return m.ValidateWithPath("CT_WheelTransition")
}
func (m *CT_WheelTransition) ValidateWithPath(path string) error {
	return nil
}
