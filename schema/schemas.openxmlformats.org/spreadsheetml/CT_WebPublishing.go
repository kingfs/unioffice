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

type CT_WebPublishing struct {
	// Use CSS
	CssAttr *bool
	// Thicket
	ThicketAttr *bool
	// Enable Long File Names
	LongFileNamesAttr *bool
	// VML in Browsers
	VmlAttr *bool
	// Allow PNG
	AllowPngAttr *bool
	// Target Screen Size
	TargetScreenSizeAttr ST_TargetScreenSize
	// DPI
	DpiAttr *uint32
	// Code Page
	CodePageAttr *uint32
	// Character Set
	CharacterSetAttr *string
}

func NewCT_WebPublishing() *CT_WebPublishing {
	ret := &CT_WebPublishing{}
	return ret
}
func (m *CT_WebPublishing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CssAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "css"},
			Value: fmt.Sprintf("%v", *m.CssAttr)})
	}
	if m.ThicketAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thicket"},
			Value: fmt.Sprintf("%v", *m.ThicketAttr)})
	}
	if m.LongFileNamesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "longFileNames"},
			Value: fmt.Sprintf("%v", *m.LongFileNamesAttr)})
	}
	if m.VmlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "vml"},
			Value: fmt.Sprintf("%v", *m.VmlAttr)})
	}
	if m.AllowPngAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowPng"},
			Value: fmt.Sprintf("%v", *m.AllowPngAttr)})
	}
	if m.TargetScreenSizeAttr != ST_TargetScreenSizeUnset {
		attr, err := m.TargetScreenSizeAttr.MarshalXMLAttr(xml.Name{Local: "targetScreenSize"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DpiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dpi"},
			Value: fmt.Sprintf("%v", *m.DpiAttr)})
	}
	if m.CodePageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codePage"},
			Value: fmt.Sprintf("%v", *m.CodePageAttr)})
	}
	if m.CharacterSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "characterSet"},
			Value: fmt.Sprintf("%v", *m.CharacterSetAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_WebPublishing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "css" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CssAttr = &parsed
		}
		if attr.Name.Local == "thicket" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ThicketAttr = &parsed
		}
		if attr.Name.Local == "longFileNames" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LongFileNamesAttr = &parsed
		}
		if attr.Name.Local == "vml" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VmlAttr = &parsed
		}
		if attr.Name.Local == "allowPng" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowPngAttr = &parsed
		}
		if attr.Name.Local == "targetScreenSize" {
			m.TargetScreenSizeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "dpi" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DpiAttr = &pt
		}
		if attr.Name.Local == "codePage" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CodePageAttr = &pt
		}
		if attr.Name.Local == "characterSet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CharacterSetAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WebPublishing: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_WebPublishing) Validate() error {
	return m.ValidateWithPath("CT_WebPublishing")
}
func (m *CT_WebPublishing) ValidateWithPath(path string) error {
	if err := m.TargetScreenSizeAttr.ValidateWithPath(path + "/TargetScreenSizeAttr"); err != nil {
		return err
	}
	return nil
}
