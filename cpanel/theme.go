package cpanel

import "github.com/arzahs/cpanelgo"

type GetThemeAPIResponse struct {
	cpanelgo.BaseUAPIResponse
	Theme string `json:"data"`
}

func (c CpanelApi) GetTheme() (GetThemeAPIResponse, error) {
	var out GetThemeAPIResponse
	err := c.Gateway.UAPI("Themes", "get_theme_base", cpanelgo.Args{}, &out)
	if err == nil {
		err = out.Error()
	}
	return out, err
}
