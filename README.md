# Namecheap Go

[![Go Reference](https://pkg.go.dev/badge/github.com/jsgv/namecheap-go.svg)](https://pkg.go.dev/github.com/jsgv/namecheap-go)

`namecheap-go` is a [Namecheap API](https://www.namecheap.com/support/api/methods/) CLI app. 
So far only methods that I personally use are implemented, but I plan to add the rest with time.


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

### Implemented methods

* Domains
    * [x] [domains.getList](https://www.namecheap.com/support/api/methods/domains/get-list/) 
    * [x] [domains.getContacts](https://www.namecheap.com/support/api/methods/domains/get-contacts/) 
    * [ ] [domains.create](https://www.namecheap.com/support/api/methods/domains/create/) 
    * [x] [domains.getTldList](https://www.namecheap.com/support/api/methods/domains/get-tld-list/) 
    * [ ] [domains.setContacts](https://www.namecheap.com/support/api/methods/domains/set-contacts/) 
    * [ ] [domains.check](https://www.namecheap.com/support/api/methods/domains/check/) 
    * [ ] [domains.reactivate](https://www.namecheap.com/support/api/methods/domains/reactivate/) 
    * [ ] [domains.renew](https://www.namecheap.com/support/api/methods/domains/renew/) 
    * [ ] [domains.getRegistrarLock](https://www.namecheap.com/support/api/methods/domains/get-registrar-lock/) 
    * [ ] [domains.setRegistrarLock](https://www.namecheap.com/support/api/methods/domains/set-registrar-lock/) 
    * [x] [domains.getInfo](https://www.namecheap.com/support/api/methods/domains/get-info/) 
* Domains DNS
    * [ ] [domains.dns.setDefault](https://www.namecheap.com/support/api/methods/domains-dns/set-default/) 
    * [x] [domains.dns.setCustom](https://www.namecheap.com/support/api/methods/domains-dns/set-custom/) 
    * [ ] [domains.dns.getList](https://www.namecheap.com/support/api/methods/domains-dns/get-list/) 
    * [ ] [domains.dns.getHosts](https://www.namecheap.com/support/api/methods/domains-dns/get-hosts/) 
    * [ ] [domains.dns.getEmailForwarding](https://www.namecheap.com/support/api/methods/domains-dns/get-email-forwarding/) 
    * [ ] [domains.dns.setEmailForwarding](https://www.namecheap.com/support/api/methods/domains-dns/set-email-forwarding/) 
    * [ ] [domains.dns.setHosts](https://www.namecheap.com/support/api/methods/domains-dns/set-hosts/) 
* Domains NS
    * [ ] [domains.ns.create](https://www.namecheap.com/support/api/methods/domains-ns/create/) 
    * [ ] [domains.ns.delete](https://www.namecheap.com/support/api/methods/domains-ns/delete/) 
    * [ ] [domains.ns.getInfo](https://www.namecheap.com/support/api/methods/domains-ns/getinfo/) 
    * [ ] [domains.ns.update](https://www.namecheap.com/support/api/methods/domains-ns/update/) 
* Domains Transfer
    * [ ] [domains.transfer.create](https://www.namecheap.com/support/api/methods/domains-transfer/create/) 
    * [ ] [domains.transfer.getStatus](https://www.namecheap.com/support/api/methods/domains-transfer/get-status/) 
    * [ ] [domains.transfer.updateStatus](https://www.namecheap.com/support/api/methods/domains-transfer/update-status/) 
    * [ ] [domains.transfer.getList](https://www.namecheap.com/support/api/methods/domains-transfer/get-list/) 
* SSL
    * [ ] [ssl.create](https://www.namecheap.com/support/api/methods/ssl/create/) 
    * [ ] [ssl.getList](https://www.namecheap.com/support/api/methods/ssl/get-list/) 
    * [ ] [ssl.parseCSR](https://www.namecheap.com/support/api/methods/ssl/parse-csr/) 
    * [ ] [ssl.getApproverEmailList](https://www.namecheap.com/support/api/methods/ssl/get-approver-email-list/) 
    * [ ] [ssl.activate](https://www.namecheap.com/support/api/methods/ssl/activate/) 
    * [ ] [ssl.resendApproverEmail](https://www.namecheap.com/support/api/methods/ssl/resend-approver-email/) 
    * [ ] [ssl.getInfo](https://www.namecheap.com/support/api/methods/ssl/get-info/) 
    * [ ] [ssl.renew](https://www.namecheap.com/support/api/methods/ssl/renew/) 
    * [ ] [ssl.reissue](https://www.namecheap.com/support/api/methods/ssl/reissue/) 
    * [ ] [ssl.resendfulfillmentemail](https://www.namecheap.com/support/api/methods/ssl/resend-fulfillment-email/) 
    * [ ] [ssl.purchasemoresans](https://www.namecheap.com/support/api/methods/ssl/purchasemoresans/) 
    * [ ] [ssl.revokecertificate](https://www.namecheap.com/support/api/methods/ssl/revokecertificate/) 
    * [ ] [ssl.editDCVMethod](https://www.namecheap.com/support/api/methods/ssl/editdcvmethod/) 
* Users
    * [ ] [users.getPricing](https://www.namecheap.com/support/api/methods/users/get-pricing/) 
    * [ ] [users.getBalances](https://www.namecheap.com/support/api/methods/users/get-balances/) 
    * [ ] [users.changePassword](https://www.namecheap.com/support/api/methods/users/change-password/) 
    * [ ] [users.update](https://www.namecheap.com/support/api/methods/users/update/) 
    * [ ] [users.createaddfundsrequest](https://www.namecheap.com/support/api/methods/users/create-add-funds-request/) 
    * [ ] [users.getAddFundsStatus](https://www.namecheap.com/support/api/methods/users/get-add-funds-status/) 
    * [ ] [users.create](https://www.namecheap.com/support/api/methods/users/create/) 
    * [ ] [users.login](https://www.namecheap.com/support/api/methods/users/login/) 
    * [ ] [users.resetPassword](https://www.namecheap.com/support/api/methods/users/reset-password/) 
* Users Address
    * [ ] [users.address.create](https://www.namecheap.com/support/api/methods/users-address/create/) 
    * [ ] [users.address.delete](https://www.namecheap.com/support/api/methods/users-address/delete/) 
    * [ ] [users.address.getInfo](https://www.namecheap.com/support/api/methods/users-address/get-info/) 
    * [ ] [users.address.getList](https://www.namecheap.com/support/api/methods/users-address/get-list/) 
    * [ ] [users.address.setDefault](https://www.namecheap.com/support/api/methods/users-address/set-default/) 
    * [ ] [users.address.update](https://www.namecheap.com/support/api/methods/users-address/update/) 
* Whoisguard
    * [ ] [whoisguard.changeemailaddress](https://www.namecheap.com/support/api/methods/whoisguard/change-email-address/) 
    * [ ] [whoisguard.enable](https://www.namecheap.com/support/api/methods/whoisguard/enable/) 
    * [ ] [whoisguard.disable](https://www.namecheap.com/support/api/methods/whoisguard/disable/) 
    * [ ] [whoisguard.getList](https://www.namecheap.com/support/api/methods/whoisguard/getlist/) 
    * [ ] [whoisguard.renew](https://www.namecheap.com/support/api/methods/whoisguard/renew/) 

