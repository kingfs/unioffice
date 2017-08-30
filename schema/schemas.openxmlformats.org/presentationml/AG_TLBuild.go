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

type AG_TLBuild struct {
	SpidAttr     *uint32
	GrpIdAttr    *uint32
	UiExpandAttr *bool
}

func NewAG_TLBuild() *AG_TLBuild {
	ret := &AG_TLBuild{}
	return ret
}
func (m *AG_TLBuild) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.GrpIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grpId"},
			Value: fmt.Sprintf("%v", *m.GrpIdAttr)})
	}
	if m.UiExpandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiExpand"},
			Value: fmt.Sprintf("%v", *m.UiExpandAttr)})
	}
	return nil
}
func (m *AG_TLBuild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpidAttr = &pt
		}
		if attr.Name.Local == "grpId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.GrpIdAttr = &pt
		}
		if attr.Name.Local == "uiExpand" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiExpandAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_TLBuild: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *AG_TLBuild) Validate() error {
	return m.ValidateWithPath("AG_TLBuild")
}
func (m *AG_TLBuild) ValidateWithPath(path string) error {
	return nil
}
