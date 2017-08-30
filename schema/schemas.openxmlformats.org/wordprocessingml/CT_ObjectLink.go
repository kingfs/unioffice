// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_ObjectLink struct {
	UpdateModeAttr  ST_ObjectUpdateMode
	LockedFieldAttr *sharedTypes.ST_OnOff
	// Object Representation
	DrawAspectAttr ST_ObjectDrawAspect
	IdAttr         string
	// Object Application
	ProgIdAttr *string
	// Object Shape
	ShapeIdAttr *string
	// Field Switches
	FieldCodesAttr *string
}

func NewCT_ObjectLink() *CT_ObjectLink {
	ret := &CT_ObjectLink{}
	ret.UpdateModeAttr = ST_ObjectUpdateMode(1)
	return ret
}
func (m *CT_ObjectLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.UpdateModeAttr.MarshalXMLAttr(xml.Name{Local: "w:updateMode"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.LockedFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lockedField"},
			Value: fmt.Sprintf("%v", *m.LockedFieldAttr)})
	}
	if m.DrawAspectAttr != ST_ObjectDrawAspectUnset {
		attr, err := m.DrawAspectAttr.MarshalXMLAttr(xml.Name{Local: "w:drawAspect"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.ProgIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:progId"},
			Value: fmt.Sprintf("%v", *m.ProgIdAttr)})
	}
	if m.ShapeIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:shapeId"},
			Value: fmt.Sprintf("%v", *m.ShapeIdAttr)})
	}
	if m.FieldCodesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fieldCodes"},
			Value: fmt.Sprintf("%v", *m.FieldCodesAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ObjectLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.UpdateModeAttr = ST_ObjectUpdateMode(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "updateMode" {
			m.UpdateModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "lockedField" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LockedFieldAttr = &parsed
		}
		if attr.Name.Local == "drawAspect" {
			m.DrawAspectAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
		if attr.Name.Local == "progId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIdAttr = &parsed
		}
		if attr.Name.Local == "shapeId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ShapeIdAttr = &parsed
		}
		if attr.Name.Local == "fieldCodes" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FieldCodesAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_ObjectLink: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_ObjectLink) Validate() error {
	return m.ValidateWithPath("CT_ObjectLink")
}
func (m *CT_ObjectLink) ValidateWithPath(path string) error {
	if m.UpdateModeAttr == ST_ObjectUpdateModeUnset {
		return fmt.Errorf("%s/UpdateModeAttr is a mandatory field", path)
	}
	if err := m.UpdateModeAttr.ValidateWithPath(path + "/UpdateModeAttr"); err != nil {
		return err
	}
	if m.LockedFieldAttr != nil {
		if err := m.LockedFieldAttr.ValidateWithPath(path + "/LockedFieldAttr"); err != nil {
			return err
		}
	}
	if err := m.DrawAspectAttr.ValidateWithPath(path + "/DrawAspectAttr"); err != nil {
		return err
	}
	return nil
}
