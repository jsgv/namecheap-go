package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdDomains(t *testing.T) {
	cmd := NewCmdDomains()

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	childCommands := cmd.Commands()

	if len(childCommands) != 5 {
		t.Error("Wrong number of child commands for `domains`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdDomainsGetList(t *testing.T) {
	cmd := NewCmdDomainsGetList()

	flags := cmd.Flags()

	listtypeValue, err := flags.GetString("listtype")
	if err != nil {
		t.Error(err)
	}
	if listtypeValue != "ALL" {
		t.Error("Wrong default value set for `listtype`")
	}

	pageValue, err := flags.GetString("page")
	if err != nil {
		t.Error(err)
	}
	if pageValue != "1" {
		t.Error("Wrong default value set for `page`")
	}

	pagesizeValue, err := flags.GetString("pagesize")
	if err != nil {
		t.Error(err)
	}
	if pagesizeValue != "20" {
		t.Error("Wrong default value set for `pagesize`")
	}

	searchtermFlag := flags.Lookup("searchterm")
	if searchtermFlag == nil {
		t.Error("Missing flag `searchterm`")
	}

	sortbyFlag := flags.Lookup("sortby")
	if sortbyFlag == nil {
		t.Error("Missing flag `sortby`")
	}
}
