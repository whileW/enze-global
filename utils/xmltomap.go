package utils

import (
	"encoding/xml"
	"io"
)

type StringMap map[string]string

type xmlMapEntry struct {
	XMLName		xml.Name
	Value 		string `xml:",chardata"`
}

func (va StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(va) == 0 {
		return nil
	}
	errT := e.EncodeToken(start)
	if errT != nil {
		return errT
	}
	for k, v := range va {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}
	errT = e.EncodeToken(start.End())
	e.Flush()
	return errT
}

func (p *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*p = StringMap{}
	for {
		var e xmlMapEntry
		errT := d.Decode(&e)
		if errT == io.EOF {
			break
		} else if errT != nil {
			return errT
		}
		(*p)[e.XMLName.Local] = e.Value
	}
	return nil
}