package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

func isFlagRequired(f *flag.Flag) bool {
	if f.Annotations == nil {
		return false
	}

	for k, v := range f.Annotations {
		if k == cobra.BashCompOneRequiredFlag && v[0] == "true" {
			return true
		}
	}

	return false
}

func TestNewCmdDomains(t *testing.T) {
	cmd := newCmdDomains()

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	childCommands := cmd.Commands()

	if len(childCommands) != 12 {
		t.Error("Wrong number of child commands for `domains`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdDomainsGetList(t *testing.T) {
	cmd := newCmdDomainsGetList()

	listtype := cmd.Flag("listtype")
	if listtype == nil {
		t.Error("Missing flag `listtype`")
	}
	if listtype.Value.String() != "ALL" {
		t.Error("Wrong default value set for `listtype`")
	}

	page := cmd.Flag("page")
	if page == nil {
		t.Error("Missing flag `page`")
	}
	if page.Value.String() != "1" {
		t.Error("Wrong default value set for `page`")
	}

	pagesize := cmd.Flag("pagesize")
	if pagesize == nil {
		t.Error("Missing flag `pagesize`")
	}
	if pagesize.Value.String() != "20" {
		t.Error("Wrong default value set for `pagesize`")
	}

	searchterm := cmd.Flag("searchterm")
	if searchterm == nil {
		t.Error("Missing flag `searchterm`")
	}

	sortby := cmd.Flag("sortby")
	if sortby == nil {
		t.Error("Missing flag `sortby`")
	}
}

func TestNewCmdDomainsGetContacts(t *testing.T) {
	cmd := newCmdDomainsGetContacts()

	domainname := cmd.Flag("domainname")
	if domainname == nil {
		t.Error("Missing flag `domainname`")
	}
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}
}

func TestNewCmdDomainsCreate(t *testing.T) {
	newCmdDomainsCreate()
}

func TestNewCmdDomainsGetTldList(t *testing.T) {
	cmd := newCmdDomainsGetTldList()

	if cmd == nil {
		t.Error("Missing command `domains gettldlist`")
	}
}

func TestNewCmdDomainsCheck(t *testing.T) {
	cmd := newCmdDomainsCheck()

	domainlist := cmd.Flag("domainlist")
	if !isFlagRequired(domainlist) {
		t.Error("Flag not markes as required `domainlist`")
	}
}

func TestNewCmdDomainsReactivate(t *testing.T) {
	newCmdDomainsReactivate()
}

func TestNewCmdDomainsRenew(t *testing.T) {
	newCmdDomainsRenew()
}

func TestNewCmdDomainsGetRegistrarLock(t *testing.T) {
	cmd := newCmdDomainsGetRegistrarLock()

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not markes as required `domainname`")
	}
}

func TestNewCmdDomainsSetRegistrarLock(t *testing.T) {
	newCmdDomainsSetRegistrarLock()
}

func TestNewCmdDomainsGetInfo(t *testing.T) {
	cmd := newCmdDomainsGetInfo()

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	hostname := cmd.Flag("hostname")
	if hostname == nil {
		t.Error("Missing flag `hostname`")
	}
}
