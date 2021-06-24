package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdSSL(t *testing.T) {
	cmd, err := newCmdSSL()
	if err != nil {
		t.Error(err)
	}

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	if len(cmd.Commands()) != 13 {
		t.Error("Wrong number of child commands for `ssl`")
	}

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}

func TestNewCmdSSLCreate(t *testing.T) {
	cmd, err := newCmdSSLCreate()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("years")) {
		t.Error("Flag not marked as required `years`")
	}

	if !isFlagRequired(cmd.Flag("type")) {
		t.Error("Flag not marked as required `type`")
	}
}

func TestNewCmdSSLParseCSR(t *testing.T) {
	cmd, err := newCmdSSLParseCSR()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("csr")) {
		t.Error("Flag not marked as required `csr`")
	}
}

func TestNewCmdSSLGetApproverEmailList(t *testing.T) {
	cmd, err := newCmdSSLGetApproverEmailList()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("domainname")) {
		t.Error("Flag not marked as required `domainname`")
	}
	if !isFlagRequired(cmd.Flag("certificatetype")) {
		t.Error("Flag not marked as required `certificatetype`")
	}
}
