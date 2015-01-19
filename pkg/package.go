package pkg

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/satori/go.uuid"

	"github.com/benwebber/mipper/amo"
)

var metadataTemplate = `<RDF xmlns="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
     xmlns:em="http://www.mozilla.org/2004/em-rdf#">

  <Description about="urn:mozilla:install-manifest">
    <em:id>{{% .GUID %}}</em:id>
    <em:type>32</em:type>
    <em:targetApplication>
      <Description>
        <!-- Firefox -->
        <em:id>{ec8030f7-c20a-464f-9b0e-13a3a9e97384}</em:id>
      </Description>
    </em:targetApplication>
  </Description>
</RDF>`

type Package struct {
	Addons []amo.Addon `json:"addons"`
	GUID   string
}

func (p *Package) Contains(addon amo.Addon) bool {
	for _, a := range p.Addons {
		if a.GUID == addon.GUID {
			return true
		}
	}
	return false
}

func (p *Package) Add(addon amo.Addon) {
	if !p.Contains(addon) {
		p.Addons = append(p.Addons, addon)
	}
}

func (p *Package) Remove(addon amo.Addon) {
	if p.Contains(addon) {
		i := p.IndexOf(addon)
		p.Addons = append(p.Addons[:i], p.Addons[i+1:]...)
	}
}

func (p *Package) IndexOf(addon amo.Addon) int {
	for i, a := range p.Addons {
		if a.GUID == addon.GUID {
			return i
		}
	}
	return -1
}

func (p *Package) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Package) Marshal() ([]byte, error) {
	data, err := json.MarshalIndent(&p, "", "  ")
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

func (p *Package) ReadFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = p.Unmarshal(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *Package) WriteFile(filename string) error {
	data, err := p.Marshal()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func NewFromFile(filename string) (*Package, error) {
	var p Package
	err := p.ReadFile(filename)
	if err != nil {
		return &Package{}, nil
	}
	return &p, nil
}

func (p *Package) FetchAll() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir := filepath.Join(pwd, "mipper_cache")
	err = os.Mkdir(cacheDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatal(err)
		}
	}
	var filename string
	for _, addon := range p.Addons {
		filename = filepath.Join(cacheDir, fmt.Sprintf("%v.xpi", addon.GUID))
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		resp, err := http.Get(addon.URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		_, err = io.Copy(f, resp.Body)
	}
}

func (p *Package) Build(filename string) {
	// Downlod all addons to cache.
	p.FetchAll()

	archive, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	zw := zip.NewWriter(archive)
	defer zw.Close()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir := filepath.Join(pwd, "mipper_cache")
	var cachedAddon string
	for _, addon := range p.Addons {
		zf, err := zw.Create(addon.GUID + ".xpi")
		if err != nil {
			log.Fatal(err)
		}
		cachedAddon = filepath.Join(cacheDir, fmt.Sprintf("%v.xpi", addon.GUID))
		f, err := ioutil.ReadFile(cachedAddon)
		_, err = zf.Write(f)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Assign the package a random ID.
	p.GUID = uuid.NewV4().String()

	// Populate install.rdf metadata.
	var metadata bytes.Buffer
	t := template.New("install.rdf")
	// This template requires literal curly braces ({}). Override the default
	// delimiters.
	t.Delims("{%", "%}")
	t.Parse(metadataTemplate)
	err = t.Execute(&metadata, p)
	if err != nil {
		log.Fatal(err)
	}

	// Write the metadata to the XPI.
	f, err := zw.Create("install.rdf")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(metadata.String()))
	if err != nil {
		log.Fatal(err)
	}
}
