package api

import (
	"testing"
)

// TestDomainsTransferCreateUrl tests the `yesno` struct tag.
func TestDomainsTransferCreateUrl(t *testing.T) {
	opts := DomainsTransferCreateOptions{
		DomainName:        "example.com",
		AddFreeWhoisguard: true,  // bool `namecheap:"yesno"`
		WGenable:          false, // bool `namecheap:"yesno"`
	}

	u, err := testNewApiUrl("namecheap.users.getContacts", opts)

	if err != nil {
		t.Fail()
	}

	q := u.Query()

	if q["DomainName"][0] != opts.DomainName {
		t.Fail()
	}

	if q["AddFreeWhoisguard"][0] != "yes" {
		t.Fail()
	}

	if q["WGenable"][0] != "no" {
		t.Fail()
	}
}
