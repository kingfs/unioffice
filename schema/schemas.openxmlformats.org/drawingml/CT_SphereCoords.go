// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_SphereCoords struct {
	LatAttr int32
	LonAttr int32
	RevAttr int32
}

func NewCT_SphereCoords() *CT_SphereCoords {
	ret := &CT_SphereCoords{}
	ret.LatAttr = 0
	ret.LonAttr = 0
	ret.RevAttr = 0
	return ret
}
func (m *CT_SphereCoords) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lat"},
		Value: fmt.Sprintf("%v", m.LatAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lon"},
		Value: fmt.Sprintf("%v", m.LonAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rev"},
		Value: fmt.Sprintf("%v", m.RevAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SphereCoords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.LatAttr = 0
	m.LonAttr = 0
	m.RevAttr = 0
	for _, attr := range start.Attr {
		if attr.Name.Local == "lat" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.LatAttr = int32(parsed)
		}
		if attr.Name.Local == "lon" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.LonAttr = int32(parsed)
		}
		if attr.Name.Local == "rev" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.RevAttr = int32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SphereCoords: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_SphereCoords) Validate() error {
	return m.ValidateWithPath("CT_SphereCoords")
}
func (m *CT_SphereCoords) ValidateWithPath(path string) error {
	if m.LatAttr < 0 {
		return fmt.Errorf("%s/m.LatAttr must be >= 0 (have %v)", path, m.LatAttr)
	}
	if m.LatAttr >= 21600000 {
		return fmt.Errorf("%s/m.LatAttr must be < 21600000 (have %v)", path, m.LatAttr)
	}
	if m.LonAttr < 0 {
		return fmt.Errorf("%s/m.LonAttr must be >= 0 (have %v)", path, m.LonAttr)
	}
	if m.LonAttr >= 21600000 {
		return fmt.Errorf("%s/m.LonAttr must be < 21600000 (have %v)", path, m.LonAttr)
	}
	if m.RevAttr < 0 {
		return fmt.Errorf("%s/m.RevAttr must be >= 0 (have %v)", path, m.RevAttr)
	}
	if m.RevAttr >= 21600000 {
		return fmt.Errorf("%s/m.RevAttr must be < 21600000 (have %v)", path, m.RevAttr)
	}
	return nil
}
