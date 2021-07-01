package api

import (
	"strconv"
	"testing"
)

func TestWhoisguardChangeEmailAddressUrl(t *testing.T) {
	opts := WhoisguardChangeEmailAddressOptions{
		WhoisguardID: 123,
	}

	u, err := testNewApiUrl("namecheap.whoisguard.changeemailaddress", opts)

	if err != nil {
		t.Fail()
	}

	q1 := u.Query()

	if q1["WhoisguardID"][0] != strconv.Itoa(opts.WhoisguardID) {
		t.Fail()
	}
}

func TestWhoisguardEnableUrl(t *testing.T) {
	opts := WhoisguardEnableOptions{
		WhoisguardID:     123,
		ForwardedToEmail: "example@example.com",
	}

	u, err := testNewApiUrl("namecheap.whoisguard.enable", opts)

	if err != nil {
		t.Fail()
	}

	q1 := u.Query()

	if q1["WhoisguardID"][0] != strconv.Itoa(opts.WhoisguardID) {
		t.Fail()
	}

	if q1["ForwardedToEmail"][0] != opts.ForwardedToEmail {
		t.Fail()
	}
}

func TestWhoisguardDisableUrl(t *testing.T) {
	opts := WhoisguardDisableOptions{
		WhoisguardID: 123,
	}

	u, err := testNewApiUrl("namecheap.whoisguard.disable", opts)

	if err != nil {
		t.Fail()
	}

	q1 := u.Query()

	if q1["WhoisguardID"][0] != strconv.Itoa(opts.WhoisguardID) {
		t.Fail()
	}
}

func TestWhoisguardGetListUrl(t *testing.T) {
	opts := WhoisguardGetListOptions{
		ListType: "ALL",
		Page:     1,
		PageSize: 20,
	}

	u, err := testNewApiUrl("namecheap.whoisguard.getList", opts)

	if err != nil {
		t.Fail()
	}

	q1 := u.Query()

	if q1["ListType"][0] != opts.ListType {
		t.Fail()
	}

	if q1["Page"][0] != strconv.Itoa(opts.Page) {
		t.Fail()
	}

	if q1["PageSize"][0] != strconv.Itoa(opts.PageSize) {
		t.Fail()
	}
}

func TestWhoisguardRenewUrl(t *testing.T) {
	opts := WhoisguardRenewOptions{
		WhoisguardID:  "1234",
		Years:         1,
		PromotionCode: 1234,
	}

	u, err := testNewApiUrl("namecheap.whoisguard.renew", opts)

	if err != nil {
		t.Fail()
	}

	q1 := u.Query()

	if q1["WhoisguardID"][0] != opts.WhoisguardID {
		t.Fail()
	}

	if q1["Years"][0] != strconv.Itoa(opts.Years) {
		t.Fail()
	}

	if q1["PromotionCode"][0] != strconv.Itoa(opts.PromotionCode) {
		t.Fail()
	}
}
