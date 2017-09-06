package data

import "encoding/xml"

// File is a data element from http://xmpp.org/extensions/xep-0096.html.
type File struct {
	XMLName xml.Name   `xml:"http://jabber.org/protocol/si/profile/file-transfer file"`
	Date    string     `xml:"date,attr,omitempty"`
	Hash    string     `xml:"hash,attr,omitempty"`
	Name    string     `xml:"name,attr,omitempty"`
	Size    int64      `xml:"size,attr,omitempty"`
	Desc    string     `xml:"desc,omitempty"`
	Range   *FileRange `xml:",omitempty"`
}

// FileRange is a data element from http://xmpp.org/extensions/xep-0096.html.
type FileRange struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/si/profile/file-transfer range"`
	Length  *int     `xml:"length,attr,omitempty"`
	Offset  *int     `xml:"offset,attr,omitempty"`
}

// IBBOpen is an element from http://xmpp.org/extensions/xep-0047.html.
type IBBOpen struct {
	XMLName   xml.Name `xml:"http://jabber.org/protocol/ibb open"`
	BlockSize int      `xml:"block-size,attr"`
	Sid       string   `xml:"sid,attr"`
	Stanza    string   `xml:"stanza,attr,omitempty"`
}

// IBBClose is an element from http://xmpp.org/extensions/xep-0047.html.
type IBBClose struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/ibb close"`
	Sid     string   `xml:"sid,attr"`
}

// IBBData is an element from http://xmpp.org/extensions/xep-0047.html.
type IBBData struct {
	XMLName  xml.Name `xml:"http://jabber.org/protocol/ibb data"`
	Sid      string   `xml:"sid,attr"`
	Sequence uint16   `xml:"seq,attr"`
	Base64   string   `xml:",chardata"`
}
