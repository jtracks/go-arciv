package arciv

import "encoding/xml"

type QueryResult struct {
	XMLName      xml.Name `xml:"feed"`
	Title        string   `xml:"title"`
	Link         Link     `xml:"link"`
	ID           string   `xml:"id"`
	Updated      string   `xml:"updated"`
	TotalResults string   `xml:"totalResults"`
	StartIndex   string   `xml:"startIndex"`
	ItemsPerPage string   `xml:"itemsPerPage"`
	Entries      []Entry  `xml:"entry"`
}

type Entry struct {
	XMLName         xml.Name `xml:"entry"`
	ID              string   `xml:"id"`
	Updated         string   `xml:"updated"`
	Published       string   `xml:"published"`
	Title           string   `xml:"title"`
	Summary         string   `xml:"summary"`
	Authors         []Author `xml:"author"`
	DOI             string   `xml:"doi"`
	Links           []Link   `xml:"link"`
	Comment         string   `xml:"comment"`
	JournalRef      string   `xml:"journal_ref"`
	PrimaryCategory Category `xml:"primary_category"`
	Category        Category `xml:"category"`
}

type Category struct {
	Name string `xml:"term,attr"`
}

type Link struct {
	XMLName xml.Name `xml:"link"`
	Title   string   `xml:"title,attr"`
	Type    string   `xml:"type,attr"`
	Rel     string   `xml:"rel,attr"`
	Href    string   `xml:"href,attr"`
}

type Author struct {
	XMLName      xml.Name      `xml:"author"`
	Name         string        `xml:"name"`
	Affiliations []Affiliation `xml:"affiliation"`
}

type Affiliation struct {
	Name string `xml:",chardata"`
}
