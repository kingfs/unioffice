// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_GroupLevels struct {
	// Grouping Level Count
	CountAttr *uint32
	// OLAP Grouping Levels
	GroupLevel []*CT_GroupLevel
}

func NewCT_GroupLevels() *CT_GroupLevels {
	ret := &CT_GroupLevels{}
	return ret
}
func (m *CT_GroupLevels) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	segroupLevel := xml.StartElement{Name: xml.Name{Local: "x:groupLevel"}}
	e.EncodeElement(m.GroupLevel, segroupLevel)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_GroupLevels) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_GroupLevels:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "groupLevel":
				tmp := NewCT_GroupLevel()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GroupLevel = append(m.GroupLevel, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupLevels
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_GroupLevels) Validate() error {
	return m.ValidateWithPath("CT_GroupLevels")
}
func (m *CT_GroupLevels) ValidateWithPath(path string) error {
	for i, v := range m.GroupLevel {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GroupLevel[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
