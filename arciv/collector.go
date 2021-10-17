package arciv

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Download all pfd for the entries in QueryResult
//	return error if any fetch has failed
func (qr QueryResult) DownloadPdfs() ([]string, []*bytes.Buffer, error) {
	// TODO, could be made concurrent?

	var numUrls int = len(qr.Entries)
	var resultFilename []string = make([]string, numUrls)
	var resultContent []*bytes.Buffer = make([]*bytes.Buffer, numUrls)

	// No entries?
	if numUrls == 0 {
		return resultFilename, resultContent, errors.New(
			"no entries in QueryResult")
	}

	// Collect all pdfs
	for _, entry := range qr.Entries {
		filename, content, err := entry.DownloadPdf()

		if err != nil {
			log.Printf("Cannot get pdf for entry with ID: %v", entry.ID)
			filename = ""
			content = bytes.NewBuffer(make([]byte, 0))
		}

		resultFilename = append(resultFilename, filename)
		resultContent = append(resultContent, content)

	}
	return resultFilename, resultContent, nil
}

// Download pdf related to entry and create an appropriate filename
// return filename, content, error
func (e Entry) DownloadPdf() (string, *bytes.Buffer, error) {
	url, err := e.PdfUrl()

	if err != nil {
		return "", nil, err
	}

	buf, err := DownloadFileFromUrl(url)

	if err != nil {
		return "", nil, err
	}

	title := strings.TrimSpace(e.Title)
	title = strings.ToLower(title)
	title = strings.Replace(title, " ", "_", -1)

	return title + ".pdf", buf, nil
}

// Get url of pdf in entry
// return url, error
func (e Entry) PdfUrl() (string, error) {

	// pdf entry is usually the last
	for i := len(e.Links) - 1; i >= 0; i-- {
		link := e.Links[i]

		if link.Title == "pdf" || link.Title == "application/pdf" {
			return link.Href, nil
		}
	}
	return "", fmt.Errorf("no pdf url for entry: %v", e.Title)
}

// Collect a file from url return content of file in url
// Requires active internet connection
func DownloadFileFromUrl(url string) (*bytes.Buffer, error) {
	var err error

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"error for request for pdf at %v\n"+
				"Reason: %v", url, resp.Body)
	}

	if resp.ContentLength == 0 {
		msg := fmt.Sprintf("no content for pdf with url: %v", url)
		log.Printf("%v", msg)
		return nil, fmt.Errorf("%v", msg)
	}

	// Read all into the buffer and return pointer to the buffer
	buf := new(bytes.Buffer)
	bufferWriter := bufio.NewWriter(buf)
	_, err = io.Copy(bufferWriter, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
