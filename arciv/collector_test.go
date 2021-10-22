package arciv

import (
	"log"
	"strings"
	"testing"
)

// Testdata
const pdfUrl = "http://arxiv.org/pdf/astro-ph/0608371v1"

var resultExample = SearchResult{
	ItemsPerPage: "2",
	Entries: []Entry{
		{
			ID: "http://arxiv.org/abs/cond-mat/0102536v1",
			Title: "Impact of Electron-Electron Cusp on Configuration " +
				"Interaction Energies",
			Links: []Link{
				{
					Title: "pdf",
					Type:  "application/pdf",
					Rel:   "related",
					Href:  "http://arxiv.org/pdf/cond-mat/0102536v1",
				},
			},
		},
		{
			ID: "http://arxiv.org/abs/1501.04914v1",
			Title: "Hamiltonian of a many-electron system with " +
				"single-electron and electron-pair states in a two-dimensional " +
				"periodic potential",
			Links: []Link{
				{
					Title: "pdf",
					Type:  "application/pdf",
					Rel:   "related",
					Href:  "http://arxiv.org/pdf/1501.04914v1",
				},
			},
		},
	},
}

// Test downloading a file from url
// and that it fills the bytesbuffer with pdf content
func TestDownloadFileFromUrl(t *testing.T) {

	buf, err := DownloadFileFromUrl(pdfUrl)
	if err != nil || buf == nil {
		log.Printf(
			"failed to download from url: %v\n"+
				"Error: %v", pdfUrl, err)
		t.FailNow()
	}

	if buf.Len() == 0 {
		log.Printf("no empty buffer for url: %v", pdfUrl)
		t.FailNow()
	}

	content := buf.String()
	if content == "" || content[0:4] != "%PDF" {
		log.Printf("failed to decode for url: %v", pdfUrl)
		t.FailNow()
	}
}

// Test getting url for pdf from an entry
func TestEntryPdfUrl(t *testing.T) {
	url, err := resultExample.Entries[0].PdfUrl()

	if err != nil {
		log.Printf("error in collecting pdfurl: %v", err)
		t.FailNow()
	}

	if url == "" || !strings.Contains(url, "http") {
		log.Printf("Pdfurl is not a url: %v", url)
		t.FailNow()
	}
}

// Download a pdf for an entry
func TestEntryDownloadPdf(t *testing.T) {
	filename, buf, err := resultExample.Entries[0].DownloadPdf()

	if err != nil {
		log.Printf("error in collecting pdf: %v", err)
		t.FailNow()
	}

	if filename == "" || !strings.Contains(filename, ".pdf") {
		log.Printf("failure to create filename for pdf: %v", filename)
		t.FailNow()
	}

	content := buf.String()
	if content == "" || content[0:4] != "%PDF" {
		log.Print("failed to decode pdf")
		t.FailNow()
	}
}

func TestSearchResultDownloadPdfs(t *testing.T) {
	filenames, bufs, err := resultExample.DownloadPdfs()

	if err != nil {
		log.Printf("Error in collecting pdfs: %v", err)
		t.FailNow()
	}

	for i := 0; i < len(filenames); i++ {

		filename, content := filenames[i], bufs[i].String()
		if filename == "" || !strings.Contains(filename, ".pdf") {
			log.Printf("failure to create filename for pdf: %v", filename)
			t.FailNow()
		}

		if content == "" || content[0:4] != "%PDF" {
			log.Print("failed to decode pdf")
			t.FailNow()
		}
	}

}
