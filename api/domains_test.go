package api

import (
	"strconv"
	"testing"
)

func TestDomainsGetListUrl(t *testing.T) {
	opts := DomainsGetListOptions{
		ListType:   "theListType",
		SearchTerm: "theSearchTerm",
		SortBy:     "CREATEDATE",
		Page:       1,
		PageSize:   4,
	}

	u, err := testNewApiUrl("namecheap.users.getList", opts)

	if err != nil {
		t.Error(err)
	}

	q := u.Query()

	if q["ListType"][0] != opts.ListType {
		t.Fail()
	}

	if q["SearchTerm"][0] != opts.SearchTerm {
		t.Fail()
	}

	if q["SortBy"][0] != opts.SortBy {
		t.Fail()
	}

	if q["Page"][0] != strconv.Itoa(opts.Page) {
		t.Fail()
	}

	if q["PageSize"][0] != strconv.Itoa(opts.PageSize) {
		t.Fail()
	}
}

func TestDomainsGetContactsUrl(t *testing.T) {
	opts := DomainsGetContactsOptions{
		DomainName: "example.com",
	}

	u, err := testNewApiUrl("namecheap.users.getContacts", opts)

	if err != nil {
		t.Fail()
	}

	q := u.Query()

	if q["DomainName"][0] != opts.DomainName {
		t.Fail()
	}
}
