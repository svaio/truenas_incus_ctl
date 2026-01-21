package cmd

import (
	"testing"
	"truenas/truenas_incus_ctl/core"

	"github.com/spf13/cobra"
)

// Helper to get the create command for a category
func getIscsiCrudCreateCmd(category string) *cobra.Command {
	for _, cmd := range iscsiCmd.Commands() {
		if cmd.Use == category {
			for _, subCmd := range cmd.Commands() {
				if subCmd.Use == "create" {
					return subCmd
				}
			}
		}
	}
	return nil
}

// Helper to get the update command for a category
func getIscsiCrudUpdateCmd(category string) *cobra.Command {
	for _, cmd := range iscsiCmd.Commands() {
		if cmd.Use == category {
			for _, subCmd := range cmd.Commands() {
				if subCmd.Use == "update" {
					return subCmd
				}
			}
		}
	}
	return nil
}

// Wrapper to adapt iscsiCrudUpdateCreate to the DoTest signature
func wrapIscsiInitiatorCreate(cmd *cobra.Command, api core.Session, args []string) error {
	return iscsiCrudUpdateCreate(cmd, "initiator", api)
}

func TestIscsiInitiatorCreateWithCommas(t *testing.T) {
	FailIf(t, DoTest(
		t,
		getIscsiCrudCreateCmd("initiator"),
		wrapIscsiInitiatorCreate,
		map[string]interface{}{
			"initiators": "iqn.1993-08.org.debian:01:abc,iqn.1993-08.org.debian:01:def",
			"comment":    "test group",
		},
		[]string{},
		[]string{
			`[{"comment":"test group","initiators":["iqn.1993-08.org.debian:01:abc","iqn.1993-08.org.debian:01:def"]}]`,
		},
		[]string{
			`{"jsonrpc":"2.0","result":{"id":1},"id":1}`,
		},
		"",
	))
}

func TestIscsiInitiatorCreateWithJsonArray(t *testing.T) {
	FailIf(t, DoTest(
		t,
		getIscsiCrudCreateCmd("initiator"),
		wrapIscsiInitiatorCreate,
		map[string]interface{}{
			"initiators": `["iqn1","iqn2","iqn3"]`,
			"comment":    "json test",
		},
		[]string{},
		[]string{
			`[{"comment":"json test","initiators":["iqn1","iqn2","iqn3"]}]`,
		},
		[]string{
			`{"jsonrpc":"2.0","result":{"id":2},"id":1}`,
		},
		"",
	))
}

func TestIscsiInitiatorUpdateById(t *testing.T) {
	FailIf(t, DoTest(
		t,
		getIscsiCrudUpdateCmd("initiator"),
		wrapIscsiInitiatorCreate,
		map[string]interface{}{
			"id":         "1",
			"initiators": "newiqn1,newiqn2",
		},
		[]string{},
		[]string{
			`[1,{"initiators":["newiqn1","newiqn2"]}]`,
		},
		[]string{
			`{"jsonrpc":"2.0","result":{"id":1},"id":1}`,
		},
		"",
	))
}

func TestIscsiInitiatorCreateEmpty(t *testing.T) {
	FailIf(t, DoTest(
		t,
		getIscsiCrudCreateCmd("initiator"),
		wrapIscsiInitiatorCreate,
		map[string]interface{}{
			"comment": "empty group",
		},
		[]string{},
		[]string{
			`[{"comment":"empty group"}]`,
		},
		[]string{
			`{"jsonrpc":"2.0","result":{"id":3},"id":1}`,
		},
		"",
	))
}
