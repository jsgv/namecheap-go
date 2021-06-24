package api

import "encoding/xml"

// UsersGetPricingResponse represents the response
// for the `namecheap.users.getPricing` method.
type UsersGetPricingResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName              xml.Name `xml:"CommandResponse" json:"-"`
		Type                 string   `xml:"Type,attr"`
		UserGetPricingResult struct {
			XMLName     xml.Name `xml:"" json:"-"`
			IsSuccess   bool     `xml:"IsSuccess,attr"`
			ProductType struct {
				XMLName         xml.Name `xml:"" json:"-"`
				Name            string   `xml:"Name,attr"`
				ProductCategory []struct {
					XMLName xml.Name `xml:"" json:"-"`
					Name    string   `xml:"Name,attr"`
					Product struct {
						XMLName xml.Name `xml:"" json:"-"`
						Name    string   `xml:"Name,attr"`
						Price   []struct {
							XMLName                   xml.Name `xml:"" json:"-"`
							Duration                  int      `xml:"Duration,attr"`
							DurationType              string   `xml:"DurationType,attr"`
							Price                     string   `xml:"Price,attr"`
							PricingType               string   `xml:"PricingType,attr"`
							AdditionalCost            string   `xml:"AdditionalCost,attr"`
							RegularPrice              string   `xml:"RegularPrice,attr"`
							RegularPriceType          string   `xml:"RegularPriceType,attr"`
							RegularAdditionalCost     string   `xml:"RegularAdditionalCost,attr"`
							RegularAdditionalCostType string   `xml:"RegularAdditionalCostType,attr"`
							YourPrice                 string   `xml:"YourPrice,attr"`
							YourPriceType             string   `xml:"YourPriceType,attr"`
							YourAdditonalCost         string   `xml:"YourAdditonalCost,attr"`
							YourAdditonalCostType     string   `xml:"YourAdditonalCostType,attr"`
							PromotionPrice            string   `xml:"PromotionPrice,attr"`
							Currency                  string   `xml:"Currency,attr"`
						}
					}
				}
			}
		}
	}
}

// UsersGetPricingOptions represents the options
// for the `namecheap.users.getPricing` method.
type UsersGetPricingOptions struct {
	ProductType     string
	ProductCategory string
	PromotionCode   string
	ActionName      string
	ProductName     string
}

// UsersGetPricing returns pricing information for a requested product type.
func (c *Client) UsersGetPricing(opts UsersGetPricingOptions) (*UsersGetPricingResponse, error) {
	r := &UsersGetPricingResponse{}
	err := c.do("namecheap.users.getPricing", opts, r)
	return r, err
}

// UsersGetBalancesResponse represents the response
// for the `namecheap.users.getBalances` method.
type UsersGetBalancesResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName               xml.Name `xml:"CommandResponse" json:"-"`
		Type                  string   `xml:"Type,attr"`
		UserGetBalancesResult struct {
			XMLName                   xml.Name `xml:"" json:"-"`
			Currency                  string   `xml:"Currency,attr"`
			AvailableBalance          string   `xml:"AvailableBalance,attr"`
			AccountBalance            string   `xml:"AccountBalance,attr"`
			EarnedAmount              string   `xml:"EarnedAmount,attr"`
			WithdrawableAmount        string   `xml:"WithdrawableAmount,attr"`
			FundsRequiredForAutoRenew string   `xml:"FundsRequiredForAutoRenew,attr"`
		}
	}
}

// UsersGetBalances gets information about fund in the user's account.
func (c *Client) UsersGetBalances() (*UsersGetBalancesResponse, error) {
	r := &UsersGetBalancesResponse{}
	err := c.do("namecheap.users.getBalances", nil, r)
	return r, err
}

// UsersChangePasswordResponse represents the response
// for the `namecheap.users.changePassword` method.
type UsersChangePasswordResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                  xml.Name `xml:"CommandResponse" json:"-"`
		Type                     string   `xml:"Type,attr"`
		UserChangePasswordResult struct {
			XMLName xml.Name `xml:"" json:"-"`
			Success bool     `xml:"Success,attr" json:",omitempty"`
			UserId  int      `xml:"UserId,attr" json:",omitempty"`
		}
	}
}

// UsersChangePasswordOptions represents the options
// for the `namecheap.users.changePassword` method.
type UsersChangePasswordOptions struct {
	OldPassword string
	NewPassword string
}

// UsersChangePassword
func (c *Client) UsersChangePassword(opts UsersChangePasswordOptions) (*UsersChangePasswordResponse, error) {
	r := &UsersChangePasswordResponse{}
	err := c.do("namecheap.users.changePassword", opts, r)
	return r, err
}

// UsersUpdateResponse represents the response
// for the `namecheap.users.update` method.
type UsersUpdateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName          xml.Name `xml:"CommandResponse" json:"-"`
		Type             string   `xml:"Type,attr"`
		UserUpdateResult struct {
			XMLName xml.Name `xml:"" json:"-"`
			Success bool     `xml:"Success,attr"`
			UserId  int      `xml:"UserId,attr"`
		}
	}
}

// UsersUpdateOptions represents the options
// for the `namecheap.users.update` method.
type UsersUpdateOptions struct {
	FirstName     string
	LastName      string
	JobTitle      string
	Organization  string
	Address1      string
	Address2      string
	City          string
	StateProvince string
	Zip           string
	Country       string
	EmailAddress  string
	Phone         string
	PhoneExt      string
	Fax           string
}

// UsersUpdate updates user account information for the particular user.
func (c *Client) UsersUpdate(opts UsersUpdateOptions) (*UsersUpdateResponse, error) {
	r := &UsersUpdateResponse{}
	err := c.do("namecheap.users.update", opts, r)
	return r, err
}

// UsersCreateAddFundsRequestResponse represents the response
// for the `namecheap.users.createaddfundsrequest` method.
type UsersCreateAddFundsRequestResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                     xml.Name `xml:"CommandResponse" json:"-"`
		Type                        string   `xml:"Type,attr" json:",omitempty"`
		Createaddfundsrequestresult struct {
			XMLName     xml.Name `xml:"" json:"-"`
			TokenID     string   `xml:"TokenID,attr" json:",omitempty"`
			ReturnURL   string   `xml:"ReturnURL,attr" json:",omitempty"`
			RedirectURL string   `xml:"RedirectURL,attr" json:",omitempty"`
		}
	}
}

// UsersCreateAddFundsRequestOptions represents the options
// for the `namecheap.users.createaddfundsrequest` method.
type UsersCreateAddFundsRequestOptions struct {
	Username    string
	PaymentType string
	Amount      string
	ReturnUrl   string
}

// UsersCreateAddFundsRequest
func (c *Client) UsersCreateAddFundsRequest(opts UsersCreateAddFundsRequestOptions) (*UsersCreateAddFundsRequestResponse, error) {
	r := &UsersCreateAddFundsRequestResponse{}
	err := c.do("namecheap.users.createaddfundsrequest", opts, r)
	return r, err
}

// UsersGetAddFundsStatusResponse represents the response
// for the `namecheap.users.getaddfundsstatus` method.
type UsersGetAddFundsStatusResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		GetAddFundsStatusResult struct {
			XMLName       xml.Name `xml:"" json:"-"`
			TransactionId int      `xml:"TransactionId,attr" json:",omitempty"`
			Amount        int      `xml:"Amount,attr" json:",omitempty"`
			Status        string   `xml:"Status,attr" json:",omitempty"`
		}
	}
}

// UsersGetAddFundsStatusOptions represents the options
// for the `namecheap.users.getaddfundsstatus` method.
type UsersGetAddFundsStatusOptions struct {
	TokenId string
}

// UsersGetAddFundsStatus gets the status of add funds request.
func (c *Client) UsersGetAddFundsStatus(opts UsersGetAddFundsStatusOptions) (*UsersGetAddFundsStatusResponse, error) {
	r := &UsersGetAddFundsStatusResponse{}
	err := c.do("namecheap.users.getaddfundsstatus", opts, r)
	return r, err
}

// UsersCreateResponse represents the response
// for the `namecheap.users.create` method.
type UsersCreateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName          xml.Name `xml:"CommandResponse" json:"-"`
		Type             string   `xml:"Type,attr"`
		UserCreateResult struct {
			XMLName xml.Name `xml:"UserCreateResult" json:"-"`
			Success bool     `xml:"Success,attr"`
			UserId  int      `xml:"UserId,attr"`
		}
	}
}

// UsersCreateOptions represents the options
// for the `namecheap.users.create` method.
type UsersCreateOptions struct {
	NewUserName                 string
	NewUserPassword             string
	EmailAddress                string
	IgnoreDuplicateEmailAddress bool `namecheap:"yesno"`
	FirstName                   string
	LastName                    string
	AcceptTerms                 int
	AcceptNews                  int
	JobTitle                    string
	Organization                string
	Address1                    string
	Address2                    string
	City                        string
	StateProvince               string
	Zip                         string
	Country                     string
	Phone                       string
	PhoneExt                    string
	Fax                         string
}

// UsersCreate creates a new account at NameCheap under this ApiUser.
func (c *Client) UsersCreate(opts UsersCreateOptions) (*UsersCreateResponse, error) {
	r := &UsersCreateResponse{}
	err := c.do("namecheap.users.create", opts, r)
	return r, err
}

// UsersLoginResponse represents the response
// for the `namecheap.users.login` method.
type UsersLoginResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName         xml.Name `xml:"CommandResponse" json:"-"`
		Type            string   `xml:"Type,attr"`
		UserLoginResult struct {
			XMLName      xml.Name `xml:"UserLoginResult" json:"-"`
			LoginSuccess bool     `xml:"LoginSuccess,attr"`
			UserName     int      `xml:"UserName,attr"`
		}
	}
}

// UsersLoginOptions represents the options
// for the `namecheap.users.login` method.
type UsersLoginOptions struct {
	Password string
}

// UsersLogin validates the username and password of user accounts you have created
// using the API command namecheap.users.create. You cannot use this command to
// validate user accounts created directly at namecheap.com.
func (c *Client) UsersLogin(opts UsersLoginOptions) (*UsersLoginResponse, error) {
	r := &UsersLoginResponse{}
	err := c.do("namecheap.users.login", opts, r)
	return r, err
}

// UsersResetPasswordResponse represents the response
// for the `namecheap.users.resetpassword` method.
type UsersResetPasswordResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		UserResetPasswordResult struct {
			XMLName xml.Name `xml:"UserResetPasswordResult" json:"-"`
			Success bool     `xml:"Success,attr"`
		}
	}
}

// UsersResetPasswordOptions represents the options
// for the `namecheap.users.resetpassword` method.
type UsersResetPasswordOptions struct {
	FindBy        string
	FindByValue   string
	EmailFromName string
	EmailFrom     string
	URLPattern    string
}

// UsersResetPassword
func (c *Client) UsersResetPassword(opts UsersResetPasswordOptions) (*UsersResetPasswordResponse, error) {
	r := &UsersResetPasswordResponse{}
	err := c.do("namecheap.users.resetpassword", opts, r)
	return r, err
}
