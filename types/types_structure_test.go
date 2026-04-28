package types_test

import (
	"testing"

	"github.com/saintedlama/archscout"
	"github.com/saintedlama/archscout/internaltest"
	"github.com/saintedlama/archscout/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func findType(t *testing.T, ws *archscout.Workspace, name string) types.Item {
	t.Helper()
	for _, ti := range ws.Types.All() {
		if ti.Name == name {
			return ti
		}
	}
	t.Fatalf("type %q not found", name)
	return types.Item{}
}

func TestTypeFields_StructFieldsAndEmbeds(t *testing.T) {
	ws := internaltest.LoadFixtureWorkspace(t, "typestructfixture", archscout.WithTypeInfo())
	user := findType(t, ws, "User")

	require.Len(t, user.Fields, 5,
		"expected one embedded entry plus Name, Email, Age, Year")

	// Embedded field appears first in source order with Embedded == true.
	embed := user.Fields[0]
	assert.True(t, embed.Embedded, "first field should be the embedded inner.Base")
	assert.Empty(t, embed.Name)
	assert.Equal(t, "inner.Base", embed.TypeName, "syntactic source text preserved")
	assert.Equal(t,
		"example.com/typestructfixture/inner.Base",
		embed.TypeQName,
		"WithTypeInfo resolves cross-package embed qname",
	)

	// Tagged fields preserve their tag contents (without backticks).
	name := user.Fields[1]
	assert.Equal(t, "Name", name.Name)
	assert.Equal(t, "string", name.TypeName)
	assert.Equal(t, `json:"name"`, name.Tag)

	email := user.Fields[2]
	assert.Equal(t, "Email", email.Name)
	assert.Equal(t, `json:"email,omitempty"`, email.Tag)

	age := user.Fields[3]
	year := user.Fields[4]
	assert.Equal(t, "Age", age.Name)
	assert.Equal(t, "Year", year.Name)
	assert.Equal(t, "int", age.TypeName)
	assert.Equal(t, "int", year.TypeName)

	assert.Equal(t,
		[]string{"example.com/typestructfixture/inner.Base"},
		user.Embeds,
		"Embeds prefers qname over syntactic text when type info is present",
	)
}

func TestTypeMethods_InterfaceMethodsAndEmbeds(t *testing.T) {
	ws := internaltest.LoadFixtureWorkspace(t, "typestructfixture", archscout.WithTypeInfo())
	svc := findType(t, ws, "Service")

	// Methods lists only directly declared methods.
	require.Len(t, svc.Methods, 2)
	assert.Equal(t, "Run", svc.Methods[0].Name)
	assert.Equal(t, "Stop", svc.Methods[1].Name)

	// QName composition uses the package ID of the declaring file.
	assert.Equal(t,
		"example.com/typestructfixture.Service.Run",
		svc.Methods[0].QName,
	)
	assert.Equal(t,
		"example.com/typestructfixture.Service.Stop",
		svc.Methods[1].QName,
	)

	assert.Contains(t, svc.Embeds, "io.Reader")
	assert.Contains(t, svc.Embeds, "example.com/typestructfixture/inner.Pinger")
}

func TestTypeStructure_WithoutTypeInfoLeavesQNamesEmpty(t *testing.T) {
	ws := internaltest.LoadFixtureWorkspace(t, "typestructfixture")

	user := findType(t, ws, "User")
	require.Len(t, user.Fields, 5)

	for _, f := range user.Fields {
		assert.Empty(t, f.TypeQName,
			"TypeQName should be empty without WithTypeInfo: %s", f.Name)
	}

	// Embeds falls back to syntactic source text in the absence of type info.
	assert.Equal(t, []string{"inner.Base"}, user.Embeds)

	svc := findType(t, ws, "Service")
	for _, m := range svc.Methods {
		assert.Empty(t, m.QName,
			"interface method QName is composed from package ID even without type info")
	}
}
