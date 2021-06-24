package api

import (
	"strconv"
	"testing"
)

func TestUsersCreateUrl(t *testing.T) {
	opts := UsersCreateOptions{
		NewUserName:                 "thenewusername",
		NewUserPassword:             "thenewuserpassword",
		EmailAddress:                "theemailaddress",
		IgnoreDuplicateEmailAddress: true,
		AcceptTerms:                 1,
		AcceptNews:                  0,
	}

	u, err := testNewApiUrl("namecheap.users.create", opts)

	if err != nil {
		t.Fail()
	}

	q1 := u.Query()

	if q1["NewUserName"][0] != opts.NewUserName {
		t.Fail()
	}

	if q1["NewUserPassword"][0] != opts.NewUserPassword {
		t.Fail()
	}

	if q1["EmailAddress"][0] != opts.EmailAddress {
		t.Fail()
	}

	if q1["IgnoreDuplicateEmailAddress"][0] != "yes" {
		t.Fail()
	}

	if q1["AcceptTerms"][0] != strconv.Itoa(opts.AcceptTerms) {
		t.Fail()
	}

	if q1["AcceptNews"][0] != strconv.Itoa(opts.AcceptNews) {
		t.Fail()
	}

	/////////////////////////

	opts2 := UsersCreateOptions{
		IgnoreDuplicateEmailAddress: false,
	}

	url2, err := testNewApiUrl("namecheap.users.create", opts2)

	if err != nil {
		t.Fail()
	}

	q2 := url2.Query()

	if q2["IgnoreDuplicateEmailAddress"][0] != "no" {
		t.Fail()
	}
}

func TestUsersLoginUrl(t *testing.T) {
	opts := UsersLoginOptions{
		Password: "the_password",
	}

	u, err := testNewApiUrl("namecheap.users.login", opts)

	if err != nil {
		t.Fail()
	}

	q := u.Query()

	if q["Password"][0] != opts.Password {
		t.Fail()
	}
}

func TestUsersResetPasswordUrl(t *testing.T) {
	opts := UsersResetPasswordOptions{
		FindBy:        "the_findby",
		FindByValue:   "the_findbyvalue",
		EmailFromName: "the_emailfromname",
		EmailFrom:     "the_emailfrom",
		URLPattern:    "the_urlpattern",
	}

	u, err := testNewApiUrl("namecheap.users.resetPassword", opts)

	if err != nil {
		t.Fail()
	}

	q := u.Query()

	if q["FindBy"][0] != opts.FindBy {
		t.Fail()
	}

	if q["FindByValue"][0] != opts.FindByValue {
		t.Fail()
	}

	if q["EmailFromName"][0] != opts.EmailFromName {
		t.Fail()
	}

	if q["EmailFrom"][0] != opts.EmailFrom {
		t.Fail()
	}

	if q["URLPattern"][0] != opts.URLPattern {
		t.Fail()
	}
}
