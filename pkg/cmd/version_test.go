package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdVersion(t *testing.T) {
	cmd, err := newCmdVersion()
	if err != nil {
		t.Error(err)
	}

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}
