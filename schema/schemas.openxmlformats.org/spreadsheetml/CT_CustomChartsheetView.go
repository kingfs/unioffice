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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_CustomChartsheetView struct {
	// GUID
	GuidAttr string
	// Print Scale
	ScaleAttr *uint32
	// Visible State
	StateAttr ST_SheetState
	// Zoom To Fit
	ZoomToFitAttr *bool
	PageMargins   *CT_PageMargins
	// Chart Sheet Page Setup
	PageSetup    *CT_CsPageSetup
	HeaderFooter *CT_HeaderFooter
}

func NewCT_CustomChartsheetView() *CT_CustomChartsheetView {
	ret := &CT_CustomChartsheetView{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}
func (m *CT_CustomChartsheetView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	if m.ScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scale"},
			Value: fmt.Sprintf("%v", *m.ScaleAttr)})
	}
	if m.StateAttr != ST_SheetStateUnset {
		attr, err := m.StateAttr.MarshalXMLAttr(xml.Name{Local: "state"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ZoomToFitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomToFit"},
			Value: fmt.Sprintf("%v", *m.ZoomToFitAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "x:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "x:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "x:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CustomChartsheetView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
		}
		if attr.Name.Local == "scale" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ScaleAttr = &pt
		}
		if attr.Name.Local == "state" {
			m.StateAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "zoomToFit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ZoomToFitAttr = &parsed
		}
	}
lCT_CustomChartsheetView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pageMargins":
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case "pageSetup":
				m.PageSetup = NewCT_CsPageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case "headerFooter":
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomChartsheetView
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CustomChartsheetView) Validate() error {
	return m.ValidateWithPath("CT_CustomChartsheetView")
}
func (m *CT_CustomChartsheetView) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if err := m.StateAttr.ValidateWithPath(path + "/StateAttr"); err != nil {
		return err
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	return nil
}
