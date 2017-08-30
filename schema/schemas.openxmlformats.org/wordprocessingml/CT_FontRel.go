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

type CT_FontRel struct {
	FontKeyAttr   string
	SubsettedAttr sharedTypes.ST_OnOff
	IdAttr        string
}

func NewCT_FontRel() *CT_FontRel {
	ret := &CT_FontRel{}
	ret.FontKeyAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}
func (m *CT_FontRel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fontKey"},
		Value: fmt.Sprintf("%v", m.FontKeyAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:subsetted"},
		Value: fmt.Sprintf("%v", m.SubsettedAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FontRel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FontKeyAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "fontKey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FontKeyAttr = parsed
		}
		if attr.Name.Local == "subsetted" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.SubsettedAttr = parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FontRel: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_FontRel) Validate() error {
	return m.ValidateWithPath("CT_FontRel")
}
func (m *CT_FontRel) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.FontKeyAttr) {
		return fmt.Errorf(`%s/m.FontKeyAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.FontKeyAttr)
	}
	if err := m.SubsettedAttr.ValidateWithPath(path + "/SubsettedAttr"); err != nil {
		return err
	}
	return nil
}
