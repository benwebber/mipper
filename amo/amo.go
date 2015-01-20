// Package amo implements a client for the addons.mozilla.org (AMO) API.
package amo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	APIURL     = "https://services.addons.mozilla.org/api" // AMO API URL
	APIVersion = "1.5"                                     // AMO API version
)

// Addon represents a Mozilla addon.
type Addon struct {
	XMLName  xml.Name `xml:"addon"   json:"-"`
	ID       uint     `xml:"id,attr" json:"id"`
	GUID     string   `xml:"guid"    json:"guid"`
	Name     string   `xml:"name"    json:"name"`
	Version  string   `xml:"version" json:"version"`
	URL      string   `xml:"install" json:"url"`
	Homepage string   `xml:"homepage" json:"homepage"`
	Summary  string   `xml:"summary" json:"summary"`
}

// searchResults represents a list of addon search results.
type searchResults struct {
	XMLName xml.Name `xml:"searchresults"`
	Addons  []Addon  `xml:"addon"`
}

// AMOClient represents an AMO API client.
type AMOClient struct {
	URL     string
	Version string
}

func (a *Addon) Fetch() ([]byte, error) {
	resp, err := http.Get(a.URL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

// NewAMOClient constructs a new AMOClient using the constant URL and versions
// defined in this package.
func NewAMOClient() *AMOClient {
	return &AMOClient{
		URL:     APIURL,
		Version: APIVersion,
	}
}

// get retrieves a resource from the AMO API and returns a byteslice containing
// the data.
//
// If there is an error requesting or reading the data, get returns nil and the
// error.
func (c *AMOClient) get(endpoint string) ([]byte, error) {
	// Construct endpoint URL.
	url := fmt.Sprintf("%v%v", c.URL, endpoint)
	// Request data.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// Read response data.
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return content, err
}

// Addon requests an addon description from the AMO API, parses the data, and
// returns a pointer to the Addon object.
func (c *AMOClient) Addon(id uint) (*Addon, error) {
	endpoint := fmt.Sprintf("/addon/%v", id)
	data, err := c.get(endpoint)
	if err != nil {
		return &Addon{}, err
	}
	var addon Addon
	err = xml.Unmarshal(data, &addon)
	if err != nil {
		return &Addon{}, err
	}
	return &addon, nil
}

// Search searches the AMO API for addons matching a particular string, and
// returns a slice of Addons.
func (c *AMOClient) Search(s string) ([]Addon, error) {
	endpoint := fmt.Sprintf("/search/%v", url.QueryEscape(s))
	data, err := c.get(endpoint)
	if err != nil {
		return nil, err
	}
	var results searchResults
	err = xml.Unmarshal(data, &results)
	if err != nil {
		return []Addon{}, err
	}
	return results.Addons, nil
}
