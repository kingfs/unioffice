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

type CT_TableStyle struct {
	// Table Style Name
	NameAttr string
	// Pivot Style
	PivotAttr *bool
	// Table
	TableAttr *bool
	// Table Style Count
	CountAttr *uint32
	// Table Style
	TableStyleElement []*CT_TableStyleElement
}

func NewCT_TableStyle() *CT_TableStyle {
	ret := &CT_TableStyle{}
	return ret
}
func (m *CT_TableStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.PivotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pivot"},
			Value: fmt.Sprintf("%v", *m.PivotAttr)})
	}
	if m.TableAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "table"},
			Value: fmt.Sprintf("%v", *m.TableAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.TableStyleElement != nil {
		setableStyleElement := xml.StartElement{Name: xml.Name{Local: "x:tableStyleElement"}}
		e.EncodeElement(m.TableStyleElement, setableStyleElement)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TableStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "pivot" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PivotAttr = &parsed
		}
		if attr.Name.Local == "table" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TableAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_TableStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tableStyleElement":
				tmp := NewCT_TableStyleElement()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TableStyleElement = append(m.TableStyleElement, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyle
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TableStyle) Validate() error {
	return m.ValidateWithPath("CT_TableStyle")
}
func (m *CT_TableStyle) ValidateWithPath(path string) error {
	for i, v := range m.TableStyleElement {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TableStyleElement[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
