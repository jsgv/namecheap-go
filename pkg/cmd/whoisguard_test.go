package cmd

import "testing"

func TestNewCmdWhoisguardChangeEmailAddress(t *testing.T) {
	cmd, err := newCmdWhoisguardChangeEmailAddress()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("whoisguardid")) {
		t.Error("Flag not marked as required `whoisguardid`")
	}
}

func TestNewCmdWhoisguardEnable(t *testing.T) {
	cmd, err := newCmdWhoisguardEnable()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("whoisguardid")) {
		t.Error("Flag not marked as required `whoisguardid`")
	}

	if !isFlagRequired(cmd.Flag("forwardedtoemail")) {
		t.Error("Flag not marked as required `forwardedtoemail`")
	}
}

func TestNewCmdWhoisguardDisable(t *testing.T) {
	cmd, err := newCmdWhoisguardDisable()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("whoisguardid")) {
		t.Error("Flag not marked as required `whoisguardid`")
	}
}

func TestNewCmdWhoisguardRenew(t *testing.T) {
	cmd, err := newCmdWhoisguardRenew()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("whoisguardid")) {
		t.Error("Flag not marked as required `whoisguardid`")
	}

	if !isFlagRequired(cmd.Flag("years")) {
		t.Error("Flag not marked as required `years`")
	}
}
