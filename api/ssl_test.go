package api

import (
	"strconv"
	"testing"
)

// TestSSLGetListUrl tests `intnozero` struct tag
func TestSSLGetListUrl(t *testing.T) {
	opts := SSLGetListOptions{
		ListType:   "theListTyDomainsGetListOptionspe",
		SearchTerm: "theSearchTerm",
		SortBy:     "CREATEDATE",
		Page:       1, // int `namecheap:"intnozero"`
		PageSize:   0, // int `namecheap:"intnozero"`
	}

	u, err := testNewApiUrl("namecheap.ssl.getList", opts)

	if err != nil {
		t.Error(err)
	}

	q := u.Query()

	if q["Page"][0] != strconv.Itoa(opts.Page) {
		t.Fail()
	}

	// we should not have a 'PageSize' url value
	if _, ok := q["PageSize"]; ok {
		t.Fail()
	}
}
