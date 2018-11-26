package whm

import "github.com/arzahs/cpanelgo"

type ListAccountsApiResponse struct {
	BaseWhmApiResponse
	Data struct {
		Accounts []struct {
			User  string `json:"user"`
			Theme string `json:"theme"`
		} `json:"acct"`
	} `json:"data"`
}

func (a WhmApi) ListAccounts() (ListAccountsApiResponse, error) {
	var out ListAccountsApiResponse

	err := a.WHMAPI1("listaccts", cpanelgo.Args{}, &out)
	if err == nil {
		err = out.Error()
	}

	return out, err
}

type AccountSummaryApiResponse struct {
	BaseWhmApiResponse
	Data struct {
		Account []struct {
			Email     string `json:"email"`
			Suspended int    `json:"suspended"`
		} `json:"acct"`
	} `json:"data"`
}

func (r AccountSummaryApiResponse) HasEmail() bool {
	e := r.Email()
	return e != "" && e != "*unknown*"
}

func (r AccountSummaryApiResponse) Email() string {
	for _, v := range r.Data.Account {
		if v.Email != "" {
			return v.Email
		}
	}
	return ""
}

func (r AccountSummaryApiResponse) Suspended() bool {
	for _, v := range r.Data.Account {
		if v.Suspended != 0 {
			return true
		}
	}
	return false
}

func (a WhmApi) AccountSummary(username string) (AccountSummaryApiResponse, error) {
	var out AccountSummaryApiResponse

	err := a.WHMAPI1("accountsummary", cpanelgo.Args{
		"user": username,
	}, &out)
	if err == nil {
		err = out.Error()
	}

	return out, err
}
