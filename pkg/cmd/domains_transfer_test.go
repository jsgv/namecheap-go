package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdDomainsTransfer(t *testing.T) {
	cmd, err := newCmdDomainsTransfer()
	if err != nil {
		t.Error(err)
	}

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	if len(cmd.Commands()) != 4 {
		t.Error("Wrong number of child commands for `domains transfer`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdDomainsTransferCreate(t *testing.T) {
	cmd, err := newCmdDomainsTransferCreate()

	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("domainname")) {
		t.Error("Flag not marked as required `domainname`")
	}

	if !isFlagRequired(cmd.Flag("years")) {
		t.Error("Flag not marked as required `years`")
	}
}

func TestNewCmdDomainsTransferGetStatus(t *testing.T) {
	cmd, err := newCmdDomainsTransferGetStatus()

	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("transferid")) {
		t.Error("Flag not marked as required `transferid`")
	}
}

func TestNewCmdDomainsTransferUpdateStatus(t *testing.T) {
	cmd, err := newCmdDomainsTransferUpdateStatus()

	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("transferid")) {
		t.Error("Flag not marked as required `transferid`")
	}
}

func TestNewCmdDomainsTransferGetList(t *testing.T) {
	_, err := newCmdDomainsTransferGetList()
	if err != nil {
		t.Error(err)
	}
}
