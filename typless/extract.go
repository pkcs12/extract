package typless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Extract function performs /extract-data API call with given API Key and file
func Extract(apiKey, inFile, template string) ([]byte, error) {
	req, err := ExtractRequestWith(apiKey, inFile, template)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyFromEnvironment},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &ExtractResponse{}
	err = json.Unmarshal(buf, response)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// ExtractRequestWith generates multipart http.Request with given API Key and file
func ExtractRequestWith(apiKey, inFile, template string) (*http.Request, error) {
	file, err := os.Open(inFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer([]byte{})
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("document_type_name", template)
	_ = writer.WriteField("line_items", "true")

	part, err := writer.CreateFormFile("file", fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContent)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://developers.typless.com/api/document-types/extract-data/",
		body,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Token %v", apiKey))
	req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%v", writer.Boundary()))

	return req, nil
}
