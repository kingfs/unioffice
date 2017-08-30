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

type CT_Scenario struct {
	// Scenario Name
	NameAttr string
	// Scenario Locked
	LockedAttr *bool
	// Hidden Scenario
	HiddenAttr *bool
	// Changing Cell Count
	CountAttr *uint32
	// User Name
	UserAttr *string
	// Scenario Comment
	CommentAttr *string
	// Input Cells
	InputCells []*CT_InputCells
}

func NewCT_Scenario() *CT_Scenario {
	ret := &CT_Scenario{}
	return ret
}
func (m *CT_Scenario) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"},
			Value: fmt.Sprintf("%v", *m.LockedAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%v", *m.HiddenAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.UserAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "user"},
			Value: fmt.Sprintf("%v", *m.UserAttr)})
	}
	if m.CommentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "comment"},
			Value: fmt.Sprintf("%v", *m.CommentAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	seinputCells := xml.StartElement{Name: xml.Name{Local: "x:inputCells"}}
	e.EncodeElement(m.InputCells, seinputCells)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Scenario) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = &parsed
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
		if attr.Name.Local == "user" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UserAttr = &parsed
		}
		if attr.Name.Local == "comment" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CommentAttr = &parsed
		}
	}
lCT_Scenario:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "inputCells":
				tmp := NewCT_InputCells()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.InputCells = append(m.InputCells, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Scenario
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Scenario) Validate() error {
	return m.ValidateWithPath("CT_Scenario")
}
func (m *CT_Scenario) ValidateWithPath(path string) error {
	for i, v := range m.InputCells {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/InputCells[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
