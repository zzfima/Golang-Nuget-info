package golangnugetinfo

import (
	"encoding/json"
	"net/http"
)

// Versions describes list of versions
type Versions struct {
	Versions []string `json:"versions"`
}

// GetNugetVersions retrieve list of versions of [nugetName] Nuget
func GetNugetVersions(nugetName string) ([]string, error) {
	response, e := http.Get("https://api.nuget.org/v3-flatcontainer/" + nugetName + "/index.json")
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()

	var versions Versions
	decoder := json.NewDecoder(response.Body)

	decoder.Decode(&versions)

	return versions.Versions, nil
}
