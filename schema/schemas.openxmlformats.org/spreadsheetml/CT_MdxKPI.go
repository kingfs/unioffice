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
	"strconv"
)

type CT_MdxKPI struct {
	// Member Unique Name Index
	NAttr uint32
	// KPI Index
	NpAttr uint32
	// KPI Property
	PAttr ST_MdxKPIProperty
}

func NewCT_MdxKPI() *CT_MdxKPI {
	ret := &CT_MdxKPI{}
	ret.PAttr = ST_MdxKPIProperty(1)
	return ret
}
func (m *CT_MdxKPI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
		Value: fmt.Sprintf("%v", m.NAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "np"},
		Value: fmt.Sprintf("%v", m.NpAttr)})
	attr, err := m.PAttr.MarshalXMLAttr(xml.Name{Local: "p"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MdxKPI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PAttr = ST_MdxKPIProperty(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NAttr = uint32(parsed)
		}
		if attr.Name.Local == "np" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NpAttr = uint32(parsed)
		}
		if attr.Name.Local == "p" {
			m.PAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MdxKPI: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_MdxKPI) Validate() error {
	return m.ValidateWithPath("CT_MdxKPI")
}
func (m *CT_MdxKPI) ValidateWithPath(path string) error {
	if m.PAttr == ST_MdxKPIPropertyUnset {
		return fmt.Errorf("%s/PAttr is a mandatory field", path)
	}
	if err := m.PAttr.ValidateWithPath(path + "/PAttr"); err != nil {
		return err
	}
	return nil
}
