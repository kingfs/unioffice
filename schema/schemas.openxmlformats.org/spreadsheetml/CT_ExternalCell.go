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

	"baliance.com/gooxml"
)

type CT_ExternalCell struct {
	// Reference
	RAttr *string
	// Type
	TAttr ST_CellType
	// Value Metadata
	VmAttr *uint32
	// Value
	V *string
}

func NewCT_ExternalCell() *CT_ExternalCell {
	ret := &CT_ExternalCell{}
	return ret
}
func (m *CT_ExternalCell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr)})
	}
	if m.TAttr != ST_CellTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "vm"},
			Value: fmt.Sprintf("%v", *m.VmAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.V != nil {
		sev := xml.StartElement{Name: xml.Name{Local: "x:v"}}
		gooxml.AddPreserveSpaceAttr(&sev, *m.V)
		e.EncodeElement(m.V, sev)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ExternalCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = &parsed
		}
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "vm" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.VmAttr = &pt
		}
	}
lCT_ExternalCell:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "v":
				m.V = new(string)
				if err := d.DecodeElement(m.V, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalCell
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ExternalCell) Validate() error {
	return m.ValidateWithPath("CT_ExternalCell")
}
func (m *CT_ExternalCell) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	return nil
}
