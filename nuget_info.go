package golangnugetinfo

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const nugetAddress = "https://api.nuget.org/v3-flatcontainer/"

// Versions describes list of versions
type Versions struct {
	Versions []string `json:"versions"`
}

// GetNugetVersions retrieve list of versions of [nugetName] Nuget
func GetNugetVersions(nugetName string) ([]string, error) {
	response, e := http.Get(nugetAddress + nugetName + "/index.json")
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()

	var versions Versions
	decoder := json.NewDecoder(response.Body)

	decoder.Decode(&versions)

	return versions.Versions, nil
}

// Package describes xml package section
type Package struct {
	XMLName  xml.Name      `xml:"package"`
	Metadata NugetMetadata `xml:"metadata"`
}

// NugetMetadata describes Nuget metadata as Nuspec
type NugetMetadata struct {
	ID          string `xml:"id"`
	Version     string `xml:"version"`
	Authors     string `xml:"authors"`
	License     string `xml:"license"`
	LicenseURL  string `xml:"licenseUrl"`
	Icon        string `xml:"icon"`
	ProjectURL  string `xml:"projectUrl"`
	IconURL     string `xml:"iconUrl"`
	Description string `xml:"description"`
	Copyright   string `xml:"copyright"`
	Tags        string `xml:"tags"`
}

// GetNugetMetadata retrieve metadata of [nugetName] Nuget of [version]
func GetNugetMetadata(nugetName string, version string) (NugetMetadata, error) {
	var pckg Package
	url := nugetAddress + nugetName + "/" + version + "/" + nugetName + ".nuspec"
	response, e := http.Get(url)
	if e != nil {
		return pckg.Metadata, e
	}
	defer response.Body.Close()

	bodyBytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return pckg.Metadata, e
	}
	e = xml.Unmarshal(bodyBytes, &pckg)
	if e != nil {
		return pckg.Metadata, e
	}

	return pckg.Metadata, nil
}
