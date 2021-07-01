package api

import (
	"strconv"
	"testing"
)

func TestUsersAddressCreateUrl(t *testing.T) {
	opts := UsersAddressCreateOptions{
		AddressName:  "the_address",
		DefaultYN:    1,
		EmailAddress: "example@example.com",
		FirstName:    "the_first_name",
		LastName:     "the_last_name",
	}

	u, err := testNewApiUrl("namecheap.users.address.create", opts)
	if err != nil {
		t.FailNow()
	}

	q1 := u.Query()

	if q1["AddressName"][0] != opts.AddressName {
		t.FailNow()
	}

	if q1["DefaultYN"][0] != strconv.Itoa(opts.DefaultYN) {
		t.FailNow()
	}

	if q1["EmailAddress"][0] != opts.EmailAddress {
		t.FailNow()
	}

	if q1["FirstName"][0] != opts.FirstName {
		t.FailNow()
	}

	if q1["LastName"][0] != opts.LastName {
		t.FailNow()
	}
}

func TestUsersAddressDeleteUrl(t *testing.T) {
	opts := UsersAddressDeleteOptions{
		AddressId: 123456,
	}

	u, err := testNewApiUrl("namecheap.users.address.delete", opts)
	if err != nil {
		t.FailNow()
	}

	q1 := u.Query()

	if q1["AddressId"][0] != strconv.Itoa(opts.AddressId) {
		t.FailNow()
	}
}

func TestUsersAddressGetInfoUrl(t *testing.T) {
	opts := UsersAddressGetInfoOptions{
		AddressId: 123456,
	}

	u, err := testNewApiUrl("namecheap.users.address.getInfo", opts)
	if err != nil {
		t.FailNow()
	}

	q1 := u.Query()

	if q1["AddressId"][0] != strconv.Itoa(opts.AddressId) {
		t.FailNow()
	}
}

func TestUsersAddressSetDefaultUrl(t *testing.T) {
	opts := UsersAddressSetDefaultOptions{
		AddressId: 123456,
	}

	u, err := testNewApiUrl("namecheap.users.address.setDefault", opts)
	if err != nil {
		t.FailNow()
	}

	q1 := u.Query()

	if q1["AddressId"][0] != strconv.Itoa(opts.AddressId) {
		t.FailNow()
	}
}

func TestUsersAddressUpdateUrl(t *testing.T) {
	opts := UsersAddressUpdateOptions{
		AddressId:           123456,
		AddressName:         "the_address_name",
		EmailAddress:        "example@example.com",
		FirstName:           "the_first_name",
		LastName:            "the_last_name",
		Address1:            "the_address1",
		City:                "the_city",
		StateProvince:       "the_state_province",
		StateProvinceChoice: "the_state_province_choice",
		Zip:                 "the_zip",
		Country:             "the_country",
		Phone:               "the_phone",
	}

	u, err := testNewApiUrl("namecheap.users.address.update", opts)
	if err != nil {
		t.FailNow()
	}

	q1 := u.Query()

	if q1["AddressId"][0] != strconv.Itoa(opts.AddressId) {
		t.FailNow()
	}

	if q1["AddressName"][0] != opts.AddressName {
		t.FailNow()
	}

	if q1["EmailAddress"][0] != opts.EmailAddress {
		t.FailNow()
	}

	if q1["FirstName"][0] != opts.FirstName {
		t.FailNow()
	}

	if q1["LastName"][0] != opts.LastName {
		t.FailNow()
	}

	if q1["Address1"][0] != opts.Address1 {
		t.FailNow()
	}

	if q1["City"][0] != opts.City {
		t.FailNow()
	}

	if q1["StateProvince"][0] != opts.StateProvince {
		t.FailNow()
	}
}
