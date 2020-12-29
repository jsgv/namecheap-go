package cmd

import (
	"bytes"
	"testing"
)

func TestNewCmdVersion(t *testing.T) {
	cmd := newCmdVersion()

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}
