package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdDomainsDNS(t *testing.T) {
	cmd, err := newCmdDomainsDNS()
	if err != nil {
		t.Error(err)
	}

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	childCommands := cmd.Commands()

	if len(childCommands) != 7 {
		t.Error("Wrong number of child commands for `domains dns`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdDomainsDNSSetDefault(t *testing.T) {
	cmd, err := newCmdDomainsDNSSetDefault()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}
}

func TestNewCmdDomainsDNSSetCustom(t *testing.T) {
	cmd, err := newCmdDomainsDNSSetCustom()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	nameservers := cmd.Flag("nameservers")
	if !isFlagRequired(nameservers) {
		t.Error("Flag not marked as required `nameservers`")
	}
}

func TestNewCmdDomainsDNSGetList(t *testing.T) {
	cmd, err := newCmdDomainsDNSGetList()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}
}

func TestNewCmdDomainsDNSGetHosts(t *testing.T) {
	cmd, err := newCmdDomainsDNSGetHosts()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}
}

func TestNewCmdDomainsDNSGetEmailForwarding(t *testing.T) {
	cmd, err := newCmdDomainsDNSGetEmailForwarding()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}
}

func TestNewCmdDomainsDNSSetEmailForwarding(t *testing.T) {
	cmd, err := newCmdDomainsDNSSetEmailForwarding()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	mailbox := cmd.Flag("mailbox")
	if !isFlagRequired(mailbox) {
		t.Error("Flag not marked as required `mailbox`")
	}

	forwardTo := cmd.Flag("forwardto")
	if !isFlagRequired(forwardTo) {
		t.Error("Flag not marked as required `forwardto`")
	}
}

func TestNewCmdDomainsDNSSetHost(t *testing.T) {
	cmd, err := newCmdDomainsDNSSetHosts()
	if err != nil {
		t.Error(err)
	}

	domainname := cmd.Flag("domainname")
	if !isFlagRequired(domainname) {
		t.Error("Flag not marked as required `domainname`")
	}

	hostname := cmd.Flag("hostname")
	if !isFlagRequired(hostname) {
		t.Error("Flag not marked as required `hostname`")
	}

	recordtype := cmd.Flag("recordtype")
	if !isFlagRequired(recordtype) {
		t.Error("Flag not marked as required `recordtype`")
	}

	address := cmd.Flag("address")
	if !isFlagRequired(address) {
		t.Error("Flag not marked as required `address`")
	}

	ttl := cmd.Flag("ttl")
	if !isFlagRequired(ttl) {
		t.Error("Flag not marked as required `ttl`")
	}

}
