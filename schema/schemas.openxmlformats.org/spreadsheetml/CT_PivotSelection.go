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

type CT_PivotSelection struct {
	// Pane
	PaneAttr ST_Pane
	// Show Header
	ShowHeaderAttr *bool
	// Label
	LabelAttr *bool
	// Data Selection
	DataAttr *bool
	// Extendable
	ExtendableAttr *bool
	// Selection Count
	CountAttr *uint32
	// Axis
	AxisAttr ST_Axis
	// Dimension
	DimensionAttr *uint32
	// Start
	StartAttr *uint32
	// Minimum
	MinAttr *uint32
	// Maximum
	MaxAttr *uint32
	// Active Row
	ActiveRowAttr *uint32
	// Active Column
	ActiveColAttr *uint32
	// Previous Row
	PreviousRowAttr *uint32
	// Previous Column Selection
	PreviousColAttr *uint32
	// Click Count
	ClickAttr *uint32
	IdAttr    *string
	// Pivot Area
	PivotArea *CT_PivotArea
}

func NewCT_PivotSelection() *CT_PivotSelection {
	ret := &CT_PivotSelection{}
	ret.PivotArea = NewCT_PivotArea()
	return ret
}
func (m *CT_PivotSelection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.PaneAttr != ST_PaneUnset {
		attr, err := m.PaneAttr.MarshalXMLAttr(xml.Name{Local: "pane"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowHeaderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showHeader"},
			Value: fmt.Sprintf("%v", *m.ShowHeaderAttr)})
	}
	if m.LabelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "label"},
			Value: fmt.Sprintf("%v", *m.LabelAttr)})
	}
	if m.DataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "data"},
			Value: fmt.Sprintf("%v", *m.DataAttr)})
	}
	if m.ExtendableAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "extendable"},
			Value: fmt.Sprintf("%v", *m.ExtendableAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.AxisAttr != ST_AxisUnset {
		attr, err := m.AxisAttr.MarshalXMLAttr(xml.Name{Local: "axis"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DimensionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dimension"},
			Value: fmt.Sprintf("%v", *m.DimensionAttr)})
	}
	if m.StartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "start"},
			Value: fmt.Sprintf("%v", *m.StartAttr)})
	}
	if m.MinAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "min"},
			Value: fmt.Sprintf("%v", *m.MinAttr)})
	}
	if m.MaxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "max"},
			Value: fmt.Sprintf("%v", *m.MaxAttr)})
	}
	if m.ActiveRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "activeRow"},
			Value: fmt.Sprintf("%v", *m.ActiveRowAttr)})
	}
	if m.ActiveColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "activeCol"},
			Value: fmt.Sprintf("%v", *m.ActiveColAttr)})
	}
	if m.PreviousRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "previousRow"},
			Value: fmt.Sprintf("%v", *m.PreviousRowAttr)})
	}
	if m.PreviousColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "previousCol"},
			Value: fmt.Sprintf("%v", *m.PreviousColAttr)})
	}
	if m.ClickAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "click"},
			Value: fmt.Sprintf("%v", *m.ClickAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sepivotArea := xml.StartElement{Name: xml.Name{Local: "x:pivotArea"}}
	e.EncodeElement(m.PivotArea, sepivotArea)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotSelection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotArea = NewCT_PivotArea()
	for _, attr := range start.Attr {
		if attr.Name.Local == "pane" {
			m.PaneAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "showHeader" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowHeaderAttr = &parsed
		}
		if attr.Name.Local == "label" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LabelAttr = &parsed
		}
		if attr.Name.Local == "data" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataAttr = &parsed
		}
		if attr.Name.Local == "extendable" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ExtendableAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
		if attr.Name.Local == "axis" {
			m.AxisAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "dimension" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DimensionAttr = &pt
		}
		if attr.Name.Local == "start" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.StartAttr = &pt
		}
		if attr.Name.Local == "min" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MinAttr = &pt
		}
		if attr.Name.Local == "max" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MaxAttr = &pt
		}
		if attr.Name.Local == "activeRow" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ActiveRowAttr = &pt
		}
		if attr.Name.Local == "activeCol" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ActiveColAttr = &pt
		}
		if attr.Name.Local == "previousRow" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PreviousRowAttr = &pt
		}
		if attr.Name.Local == "previousCol" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PreviousColAttr = &pt
		}
		if attr.Name.Local == "click" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ClickAttr = &pt
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
lCT_PivotSelection:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pivotArea":
				if err := d.DecodeElement(m.PivotArea, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotSelection
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotSelection) Validate() error {
	return m.ValidateWithPath("CT_PivotSelection")
}
func (m *CT_PivotSelection) ValidateWithPath(path string) error {
	if err := m.PaneAttr.ValidateWithPath(path + "/PaneAttr"); err != nil {
		return err
	}
	if err := m.AxisAttr.ValidateWithPath(path + "/AxisAttr"); err != nil {
		return err
	}
	if err := m.PivotArea.ValidateWithPath(path + "/PivotArea"); err != nil {
		return err
	}
	return nil
}
