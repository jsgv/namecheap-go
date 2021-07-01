# Namecheap Go

[![Go Reference](https://pkg.go.dev/badge/github.com/jsgv/namecheap-go.svg)](https://pkg.go.dev/github.com/jsgv/namecheap-go)
![Go](https://github.com/jsgv/namecheap-go/workflows/Go/badge.svg?branch=master)
[![Build Status](https://travis-ci.com/jsgv/namecheap-go.svg?branch=master)](https://travis-ci.com/jsgv/namecheap-go)

`namecheap-go` is a [Namecheap API](https://www.namecheap.com/support/api/methods/) CLI app.

## Installation
```
go get github.com/jsgv/namecheap-go/cmd/namecheap
```

## Usage

* Call a command

`namecheap domains getlist`

* Print help message for command with description and available flags

`namecheap domains getlist -h`

## Auth
Auth configuration can be set with flags or a file.
You will need to get the auth values from [Namecheap](https://www.namecheap.com/support/api/intro/).

### Flags
| Flag | Description |
| ---- | ----------- |
| --apiKey | Password required used to access the API |
| --apiUser | Username required to access the API |
| --clientIP | An IP address of the server from which our system receives API calls (only IPv4 can be used) |

```
namecheap domains getinfo --domainname example.com \
      --apiKey XXXXXXXXXXXXXXX --apiUser XXXXXXXX --clientIP 127.0.0.1
```

### File
Configuration file path is `~/.config/namecheap/config.yml` with the same names as the flags.
With this option you don't have to set the configuration params with every call.

```yml
apiKey: XXXXXXXXXXXXXXX
apiUser: XXXXXXXX
clientIP: 127.0.0.1
```

`namecheap domains getinfo --domainname example.com`

## Note
I have not tested all the live responses returned from the API. There might be some wrong ones.
I have found incorrect details in the Namecheap API documentation.
Feel free to open an issue letting me know and I will fix it.

## Implemented methods

* Domains
    * [x] [domains.getList](https://www.namecheap.com/support/api/methods/domains/get-list/)
    * [x] [domains.getContacts](https://www.namecheap.com/support/api/methods/domains/get-contacts/)
    * [x] [domains.create](https://www.namecheap.com/support/api/methods/domains/create/)
    * [x] [domains.getTldList](https://www.namecheap.com/support/api/methods/domains/get-tld-list/)
    * [x] [domains.setContacts](https://www.namecheap.com/support/api/methods/domains/set-contacts/)
    * [x] [domains.check](https://www.namecheap.com/support/api/methods/domains/check/)
    * [x] [domains.reactivate](https://www.namecheap.com/support/api/methods/domains/reactivate/)
    * [x] [domains.renew](https://www.namecheap.com/support/api/methods/domains/renew/)
    * [x] [domains.getRegistrarLock](https://www.namecheap.com/support/api/methods/domains/get-registrar-lock/)
    * [x] [domains.setRegistrarLock](https://www.namecheap.com/support/api/methods/domains/set-registrar-lock/)
    * [x] [domains.getInfo](https://www.namecheap.com/support/api/methods/domains/get-info/)
* Domains DNS
    * [x] [domains.dns.setDefault](https://www.namecheap.com/support/api/methods/domains-dns/set-default/)
    * [x] [domains.dns.setCustom](https://www.namecheap.com/support/api/methods/domains-dns/set-custom/)
    * [x] [domains.dns.getList](https://www.namecheap.com/support/api/methods/domains-dns/get-list/)
    * [x] [domains.dns.getHosts](https://www.namecheap.com/support/api/methods/domains-dns/get-hosts/)
    * [x] [domains.dns.getEmailForwarding](https://www.namecheap.com/support/api/methods/domains-dns/get-email-forwarding/)
    * [x] [domains.dns.setEmailForwarding](https://www.namecheap.com/support/api/methods/domains-dns/set-email-forwarding/)
    * [x] [domains.dns.setHosts](https://www.namecheap.com/support/api/methods/domains-dns/set-hosts/)
* Domains NS
    * [x] [domains.ns.create](https://www.namecheap.com/support/api/methods/domains-ns/create/)
    * [x] [domains.ns.delete](https://www.namecheap.com/support/api/methods/domains-ns/delete/)
    * [x] [domains.ns.getInfo](https://www.namecheap.com/support/api/methods/domains-ns/getinfo/)
    * [x] [domains.ns.update](https://www.namecheap.com/support/api/methods/domains-ns/update/)
* Domains Transfer
    * [x] [domains.transfer.create](https://www.namecheap.com/support/api/methods/domains-transfer/create/)
    * [x] [domains.transfer.getStatus](https://www.namecheap.com/support/api/methods/domains-transfer/get-status/)
    * [x] [domains.transfer.updateStatus](https://www.namecheap.com/support/api/methods/domains-transfer/update-status/)
    * [x] [domains.transfer.getList](https://www.namecheap.com/support/api/methods/domains-transfer/get-list/)
* SSL
    * [x] [ssl.create](https://www.namecheap.com/support/api/methods/ssl/create/)
    * [x] [ssl.getList](https://www.namecheap.com/support/api/methods/ssl/get-list/)
    * [x] [ssl.parseCSR](https://www.namecheap.com/support/api/methods/ssl/parse-csr/)
    * [x] [ssl.getApproverEmailList](https://www.namecheap.com/support/api/methods/ssl/get-approver-email-list/)
    * [x] [ssl.activate](https://www.namecheap.com/support/api/methods/ssl/activate/)
    * [x] [ssl.resendApproverEmail](https://www.namecheap.com/support/api/methods/ssl/resend-approver-email/)
    * [x] [ssl.getInfo](https://www.namecheap.com/support/api/methods/ssl/get-info/)
    * [x] [ssl.renew](https://www.namecheap.com/support/api/methods/ssl/renew/)
    * [x] [ssl.reissue](https://www.namecheap.com/support/api/methods/ssl/reissue/)
    * [x] [ssl.resendfulfillmentemail](https://www.namecheap.com/support/api/methods/ssl/resend-fulfillment-email/)
    * [x] [ssl.purchasemoresans](https://www.namecheap.com/support/api/methods/ssl/purchasemoresans/)
    * [x] [ssl.revokecertificate](https://www.namecheap.com/support/api/methods/ssl/revokecertificate/)
    * [x] [ssl.editDCVMethod](https://www.namecheap.com/support/api/methods/ssl/editdcvmethod/)
* Users
    * [x] [users.getPricing](https://www.namecheap.com/support/api/methods/users/get-pricing/)
    * [x] [users.getBalances](https://www.namecheap.com/support/api/methods/users/get-balances/)
    * [x] [users.changePassword](https://www.namecheap.com/support/api/methods/users/change-password/)
    * [x] [users.update](https://www.namecheap.com/support/api/methods/users/update/)
    * [x] [users.createaddfundsrequest](https://www.namecheap.com/support/api/methods/users/create-add-funds-request/)
    * [x] [users.getAddFundsStatus](https://www.namecheap.com/support/api/methods/users/get-add-funds-status/)
    * [x] [users.create](https://www.namecheap.com/support/api/methods/users/create/)
    * [x] [users.login](https://www.namecheap.com/support/api/methods/users/login/)
    * [x] [users.resetPassword](https://www.namecheap.com/support/api/methods/users/reset-password/)
* Users Address
    * [x] [users.address.create](https://www.namecheap.com/support/api/methods/users-address/create/)
    * [x] [users.address.delete](https://www.namecheap.com/support/api/methods/users-address/delete/)
    * [x] [users.address.getInfo](https://www.namecheap.com/support/api/methods/users-address/get-info/)
    * [x] [users.address.getList](https://www.namecheap.com/support/api/methods/users-address/get-list/)
    * [x] [users.address.setDefault](https://www.namecheap.com/support/api/methods/users-address/set-default/)
    * [x] [users.address.update](https://www.namecheap.com/support/api/methods/users-address/update/)
* Whoisguard
    * [x] [whoisguard.changeemailaddress](https://www.namecheap.com/support/api/methods/whoisguard/change-email-address/)
    * [x] [whoisguard.enable](https://www.namecheap.com/support/api/methods/whoisguard/enable/)
    * [x] [whoisguard.disable](https://www.namecheap.com/support/api/methods/whoisguard/disable/)
    * [x] [whoisguard.getList](https://www.namecheap.com/support/api/methods/whoisguard/getlist/)
    * [x] [whoisguard.renew](https://www.namecheap.com/support/api/methods/whoisguard/renew/)

