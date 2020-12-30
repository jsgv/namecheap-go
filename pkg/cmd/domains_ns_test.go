package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdDomainsNS(t *testing.T) {
	cmd, err := newCmdDomainsNS()
	if err != nil {
		t.Error(err)
	}

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	childCommands := cmd.Commands()

	if len(childCommands) != 4 {
		t.Error("Wrong number of child commands for `domains ns`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdDomainsNSCreate(t *testing.T) {
	cmd, err := newCmdDomainsNSCreate()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	nameserver := cmd.Flag("nameserver")
	if !isFlagRequired(nameserver) {
		t.Error("Flag not marked as required `nameserver`")
	}

	ip := cmd.Flag("ip")
	if !isFlagRequired(ip) {
		t.Error("Flag not marked as required `ip`")
	}
}
func TestNewCmdDomainsNSDelete(t *testing.T) {
	cmd, err := newCmdDomainsNSDelete()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	nameserver := cmd.Flag("nameserver")
	if !isFlagRequired(nameserver) {
		t.Error("Flag not marked as required `nameserver`")
	}
}

func TestNewCmdDomainsNSGetInfo(t *testing.T) {
	cmd, err := newCmdDomainsNSGetInfo()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	nameserver := cmd.Flag("nameserver")
	if !isFlagRequired(nameserver) {
		t.Error("Flag not marked as required `nameserver`")
	}

}

func TestNewCmdDomainsNSUpdate(t *testing.T) {
	cmd, err := newCmdDomainsNSUpdate()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	nameserver := cmd.Flag("nameserver")
	if !isFlagRequired(nameserver) {
		t.Error("Flag not marked as required `nameserver`")
	}

	oldip := cmd.Flag("oldip")
	if !isFlagRequired(oldip) {
		t.Error("Flag not marked as required `oldip`")
	}

	ip := cmd.Flag("ip")
	if !isFlagRequired(ip) {
		t.Error("Flag not marked as required `ip`")
	}
}
