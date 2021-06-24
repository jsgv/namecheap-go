package cmd

import "testing"

func TestNewCmdUsersCreate(t *testing.T) {
	cmd, err := newCmdUsersCreate()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("newusername")) {
		t.Error("Flag not marked as required `newusername`")
	}

	if !isFlagRequired(cmd.Flag("newuserpassword")) {
		t.Error("Flag not marked as required `newuserpassword`")
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

func TestNewCmdUsersLogin(t *testing.T) {
	cmd, err := newCmdUsersLogin()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("password")) {
		t.Error("Flag not marked as required `password`")
	}
}

func TestNewCmdUsersResetPassword(t *testing.T) {
	cmd, err := newCmdUsersResetPassword()
	if err != nil {
		t.Error(err)
	}

	if !isFlagRequired(cmd.Flag("findby")) {
		t.Error("Flag not marked as required `findby`")
	}

	if !isFlagRequired(cmd.Flag("findbyvalue")) {
		t.Error("Flag not marked as required `findbyvalue`")
	}
}
