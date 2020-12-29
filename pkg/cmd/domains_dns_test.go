package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdDomainsDns(t *testing.T) {
	cmd := newCmdDomainsDns()

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	childCommands := cmd.Commands()

	if len(childCommands) != 1 {
		t.Error("Wrong number of child commands for `domains dns`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdDomainsDnsSetCustom(t *testing.T) {
	cmd := newCmdDomainsDnsSetCustom()

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	nameservers := cmd.Flag("nameservers")
	if !isFlagRequired(nameservers) {
		t.Error("Flag not marked as required `nameservers`")
	}
}
