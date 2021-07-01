package cmd

import "testing"

func TestNewCmdUsersAddressCreate(t *testing.T) {
	cmd, err := newCmdUsersAddressCreate()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("addressname")) {
		t.Error("Flag not marked as required `addressname`")
	}

	if !isFlagRequired(cmd.Flag("emailaddress")) {
		t.Error("Flag not marked as required `emailaddress`")
	}

	if !isFlagRequired(cmd.Flag("firstname")) {
		t.Error("Flag not marked as required `firstname`")
	}

	if !isFlagRequired(cmd.Flag("lastname")) {
		t.Error("Flag not marked as required `lastname`")
	}

	if !isFlagRequired(cmd.Flag("address1")) {
		t.Error("Flag not marked as required `address1`")
	}

	if !isFlagRequired(cmd.Flag("city")) {
		t.Error("Flag not marked as required `city`")
	}

	if !isFlagRequired(cmd.Flag("stateprovince")) {
		t.Error("Flag not marked as required `stateprovince`")
	}

	if !isFlagRequired(cmd.Flag("stateprovincechoice")) {
		t.Error("Flag not marked as required `stateprovincechoice`")
	}

	if !isFlagRequired(cmd.Flag("zip")) {
		t.Error("Flag not marked as required `zip`")
	}

	if !isFlagRequired(cmd.Flag("country")) {
		t.Error("Flag not marked as required `country`")
	}

	if !isFlagRequired(cmd.Flag("phone")) {
		t.Error("Flag not marked as required `phone`")
	}
}

func TestNewCmdUsersAddressDelete(t *testing.T) {
	cmd, err := newCmdUsersAddressDelete()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("addressid")) {
		t.Error("Flag not marked as required `addressid`")
	}
}

func TestNewCmdUsersAddressGetInfo(t *testing.T) {
	cmd, err := newCmdUsersAddressGetInfo()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("addressid")) {
		t.Error("Flag not marked as required `addressid`")
	}
}

func TestNewCmdUsersAddressSetDefault(t *testing.T) {
	cmd, err := newCmdUsersAddressSetDefault()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("addressid")) {
		t.Error("Flag not marked as required `addressid`")
	}
}

func TestNewCmdUsersAddressUpdate(t *testing.T) {
	cmd, err := newCmdUsersAddressUpdate()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("addressid")) {
		t.Error("Flag not marked as required `addressid`")
	}

	if !isFlagRequired(cmd.Flag("addressname")) {
		t.Error("Flag not marked as required `addressname`")
	}

	if !isFlagRequired(cmd.Flag("emailaddress")) {
		t.Error("Flag not marked as required `emailaddress`")
	}

	if !isFlagRequired(cmd.Flag("firstname")) {
		t.Error("Flag not marked as required `firstname`")
	}

	if !isFlagRequired(cmd.Flag("lastname")) {
		t.Error("Flag not marked as required `lastname`")
	}

	if !isFlagRequired(cmd.Flag("address1")) {
		t.Error("Flag not marked as required `address1`")
	}

	if !isFlagRequired(cmd.Flag("city")) {
		t.Error("Flag not marked as required `city`")
	}

	if !isFlagRequired(cmd.Flag("stateprovince")) {
		t.Error("Flag not marked as required `stateprovince`")
	}

	if !isFlagRequired(cmd.Flag("stateprovincechoice")) {
		t.Error("Flag not marked as required `stateprovincechoice`")
	}

	if !isFlagRequired(cmd.Flag("zip")) {
		t.Error("Flag not marked as required `zip`")
	}

	if !isFlagRequired(cmd.Flag("country")) {
		t.Error("Flag not marked as required `country`")
	}

	if !isFlagRequired(cmd.Flag("phone")) {
		t.Error("Flag not marked as required `phone`")
	}
}
