package golangnugetinfo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	v, e := GetNugetVersions("Castle.Core")
	require.Nil(t, e)
	require.Equal(t, "1.1.0", v[0])
}

func TestLicense(t *testing.T) {
	v, e := GetNugetMetadata("Castle.Core", "5.1.0")
	require.Nil(t, e)
	require.Equal(t, "Apache-2.0", v.License)
}

func TestMetadata(t *testing.T) {
	m, e := GetNugetMetadata("Castle.Core", "5.1.0")
	require.Nil(t, e)
	require.Equal(t, "Castle.Core", m.ID)
	require.Equal(t, "5.1.0", m.Version)
	require.Equal(t, "Castle Project Contributors", m.Authors)
	require.Equal(t, "Apache-2.0", m.License)
	require.Equal(t, "https://licenses.nuget.org/Apache-2.0", m.LicenseURL)
	require.Equal(t, "castle-logo.png", m.Icon)
	require.Equal(t, "http://www.castleproject.org/", m.ProjectURL)
	require.Equal(t, "http://www.castleproject.org/img/castle-logo.png", m.IconURL)
	require.Equal(t, "Castle Core, including DynamicProxy, Logging Abstractions and DictionaryAdapter", m.Description)
	require.Equal(t, "Copyright (c) 2004-2022 Castle Project - http://www.castleproject.org/", m.Copyright)
	require.Equal(t, "castle dynamicproxy dynamic proxy dynamicproxy2 dictionaryadapter emailsender", m.Tags)
}
