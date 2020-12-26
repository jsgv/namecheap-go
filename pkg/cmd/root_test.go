package cmd

import (
	"bytes"
	"testing"
)

func TestRootCmd(t *testing.T) {
	cmd := newRootCmd()

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}
