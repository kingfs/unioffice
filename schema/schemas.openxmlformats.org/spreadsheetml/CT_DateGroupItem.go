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

type CT_DateGroupItem struct {
	// Year
	YearAttr uint16
	// Month
	MonthAttr *uint16
	// Day
	DayAttr *uint16
	// Hour
	HourAttr *uint16
	// Minute
	MinuteAttr *uint16
	// Second
	SecondAttr *uint16
	// Date Time Grouping
	DateTimeGroupingAttr ST_DateTimeGrouping
}

func NewCT_DateGroupItem() *CT_DateGroupItem {
	ret := &CT_DateGroupItem{}
	ret.DateTimeGroupingAttr = ST_DateTimeGrouping(1)
	return ret
}
func (m *CT_DateGroupItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "year"},
		Value: fmt.Sprintf("%v", m.YearAttr)})
	if m.MonthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "month"},
			Value: fmt.Sprintf("%v", *m.MonthAttr)})
	}
	if m.DayAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "day"},
			Value: fmt.Sprintf("%v", *m.DayAttr)})
	}
	if m.HourAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hour"},
			Value: fmt.Sprintf("%v", *m.HourAttr)})
	}
	if m.MinuteAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minute"},
			Value: fmt.Sprintf("%v", *m.MinuteAttr)})
	}
	if m.SecondAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "second"},
			Value: fmt.Sprintf("%v", *m.SecondAttr)})
	}
	attr, err := m.DateTimeGroupingAttr.MarshalXMLAttr(xml.Name{Local: "dateTimeGrouping"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DateGroupItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.DateTimeGroupingAttr = ST_DateTimeGrouping(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "year" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			m.YearAttr = uint16(parsed)
		}
		if attr.Name.Local == "month" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.MonthAttr = &pt
		}
		if attr.Name.Local == "day" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.DayAttr = &pt
		}
		if attr.Name.Local == "hour" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.HourAttr = &pt
		}
		if attr.Name.Local == "minute" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.MinuteAttr = &pt
		}
		if attr.Name.Local == "second" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.SecondAttr = &pt
		}
		if attr.Name.Local == "dateTimeGrouping" {
			m.DateTimeGroupingAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DateGroupItem: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_DateGroupItem) Validate() error {
	return m.ValidateWithPath("CT_DateGroupItem")
}
func (m *CT_DateGroupItem) ValidateWithPath(path string) error {
	if m.DateTimeGroupingAttr == ST_DateTimeGroupingUnset {
		return fmt.Errorf("%s/DateTimeGroupingAttr is a mandatory field", path)
	}
	if err := m.DateTimeGroupingAttr.ValidateWithPath(path + "/DateTimeGroupingAttr"); err != nil {
		return err
	}
	return nil
}
