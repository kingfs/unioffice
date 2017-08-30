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

type CT_WebPublishItem struct {
	// Id
	IdAttr uint32
	// Destination Bookmark
	DivIdAttr string
	// Web Source Type
	SourceTypeAttr ST_WebSourceType
	// Source Id
	SourceRefAttr *string
	// Source Object Name
	SourceObjectAttr *string
	// Destination File Name
	DestinationFileAttr string
	// Title
	TitleAttr *string
	// Automatically Publish
	AutoRepublishAttr *bool
}

func NewCT_WebPublishItem() *CT_WebPublishItem {
	ret := &CT_WebPublishItem{}
	ret.SourceTypeAttr = ST_WebSourceType(1)
	return ret
}
func (m *CT_WebPublishItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "divId"},
		Value: fmt.Sprintf("%v", m.DivIdAttr)})
	attr, err := m.SourceTypeAttr.MarshalXMLAttr(xml.Name{Local: "sourceType"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.SourceRefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceRef"},
			Value: fmt.Sprintf("%v", *m.SourceRefAttr)})
	}
	if m.SourceObjectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceObject"},
			Value: fmt.Sprintf("%v", *m.SourceObjectAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "destinationFile"},
		Value: fmt.Sprintf("%v", m.DestinationFileAttr)})
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.AutoRepublishAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoRepublish"},
			Value: fmt.Sprintf("%v", *m.AutoRepublishAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WebPublishItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SourceTypeAttr = ST_WebSourceType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "divId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DivIdAttr = parsed
		}
		if attr.Name.Local == "sourceType" {
			m.SourceTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "sourceRef" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SourceRefAttr = &parsed
		}
		if attr.Name.Local == "sourceObject" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SourceObjectAttr = &parsed
		}
		if attr.Name.Local == "destinationFile" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DestinationFileAttr = parsed
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
		}
		if attr.Name.Local == "autoRepublish" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoRepublishAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WebPublishItem: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_WebPublishItem) Validate() error {
	return m.ValidateWithPath("CT_WebPublishItem")
}
func (m *CT_WebPublishItem) ValidateWithPath(path string) error {
	if m.SourceTypeAttr == ST_WebSourceTypeUnset {
		return fmt.Errorf("%s/SourceTypeAttr is a mandatory field", path)
	}
	if err := m.SourceTypeAttr.ValidateWithPath(path + "/SourceTypeAttr"); err != nil {
		return err
	}
	return nil
}
