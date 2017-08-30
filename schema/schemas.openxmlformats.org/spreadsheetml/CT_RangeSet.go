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

type CT_RangeSet struct {
	// Field Item Index Page 1
	I1Attr *uint32
	// Field Item Index Page 2
	I2Attr *uint32
	// Field Item index Page 3
	I3Attr *uint32
	// Field Item Index Page 4
	I4Attr *uint32
	// Reference
	RefAttr *string
	// Named Range
	NameAttr *string
	// Sheet Name
	SheetAttr *string
	IdAttr    *string
}

func NewCT_RangeSet() *CT_RangeSet {
	ret := &CT_RangeSet{}
	return ret
}
func (m *CT_RangeSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.I1Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i1"},
			Value: fmt.Sprintf("%v", *m.I1Attr)})
	}
	if m.I2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i2"},
			Value: fmt.Sprintf("%v", *m.I2Attr)})
	}
	if m.I3Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i3"},
			Value: fmt.Sprintf("%v", *m.I3Attr)})
	}
	if m.I4Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i4"},
			Value: fmt.Sprintf("%v", *m.I4Attr)})
	}
	if m.RefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
			Value: fmt.Sprintf("%v", *m.RefAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.SheetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheet"},
			Value: fmt.Sprintf("%v", *m.SheetAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RangeSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "i1" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.I1Attr = &pt
		}
		if attr.Name.Local == "i2" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.I2Attr = &pt
		}
		if attr.Name.Local == "i3" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.I3Attr = &pt
		}
		if attr.Name.Local == "i4" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.I4Attr = &pt
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = &parsed
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "sheet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SheetAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RangeSet: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RangeSet) Validate() error {
	return m.ValidateWithPath("CT_RangeSet")
}
func (m *CT_RangeSet) ValidateWithPath(path string) error {
	return nil
}
