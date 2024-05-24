package novofon

// CatchInfoBalance https://novofon.com/instructions/api/#api_info_balance_nvf
type CatchInfoBalance struct {
	Status   string  `json:"status"`
	Balance  float32 `json:"balance"`
	Currency string  `json:"currency"`
	Message  string  `json:"message"`
}

// CatchInfoPrice https://novofon.com/instructions/api/#api_info_price_nvf
type CatchInfoPrice struct {
	Status string `json:"status"`
	Info   struct {
		Prefix      string  `json:"prefix"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		Currency    string  `json:"currency"`
	}
	Message string `json:"message"`
}

// CatchInfoTimezone https://novofon.com/instructions/api/#api_info_timezone
type CatchInfoTimezone struct {
	Status   string `json:"status"`
	Unixtime int    `json:"unixtime"`
	Datetime string `json:"datetime"`
	Timezone string `json:"timezone"`
	Message  string `json:"message"`
}

// CatchTariff https://novofon.com/instructions/api/#api_tariff_nvf
type CatchTariff struct {
	Status string `json:"status"`
	Info   struct {
		TariffId              int     `json:"tariff_id"`
		TariffName            string  `json:"tariff_name"`
		IsActive              bool    `json:"is_active"`
		Cost                  float32 `json:"cost"`
		Currency              string  `json:"currency"`
		UsedSeconds           int     `json:"used_seconds"`
		UsedSecondsMobile     int     `json:"used_seconds_mobile"`
		UsedSecondsFix        int     `json:"used_seconds_fix"`
		TariffIdForNextPeriod int     `json:"tariff_id_for_next_period"`
		TariffForNextPeriod   int     `json:"tariff_for_next_period"`
	}
	Message string `json:"message"`
}

// CatchRequestCallback https://novofon.com/instructions/api/#api_callback_nvf
type CatchRequestCallback struct {
	Status  string `json:"status"`
	From    int    `json:"from"`
	To      int    `json:"to"`
	Time    int    `json:"time"`
	Message string `json:"message"`
}

// CatchRequestChecknumber https://novofon.com/instructions/api/#api_request_checknumber_nvf
type CatchRequestChecknumber struct {
	Status  string `json:"status"`
	From    int    `json:"from"`
	To      int    `json:"to"`
	Lang    string `json:"lang"`
	Time    int    `json:"time"`
	Message string `json:"message"`
}

// CatchInfoNumber_lookup https://novofon.com/instructions/api/#api_info_number_lookup
//type CatchInfoNumber struct{}

// CatchInfoListsCurrencies https://novofon.com/instructions/api/#api_info_lists_currencies_nvf
type CatchInfoListsCurrencies struct {
	Status     string   `json:"status"`
	Currencies []string `json:"currencies"`
	Message    string   `json:"message"`
}

// CatchInfoListsLanguages https://novofon.com/instructions/api/#api_info_lists_languages_nvf
type CatchInfoListsLanguages struct {
	Status    string   `json:"status"`
	Languages []string `json:"languages"`
	Message   string   `json:"message"`
}

// CatchInfoListsTariffs https://novofon.com/instructions/api/#api_info_lists_tariffs_nvf
//type CatchInfoListsTariffs struct{}

// CatchSip https://novofon.com/instructions/api/#api_sip_method
type CatchSip struct {
	Status string `json:"status"`
	Sips   []struct {
		Id          string `json:"id"`
		DisplayName string `json:"display_name"`
		Lines       int    `json:"lines"`
	} `json:"sips"`
	Left    int    `json:"left"`
	Message string `json:"message"`
}

// CatchSipSipStatus https://novofon.com/instructions/api/#api_sip_status
type CatchSipSipStatus struct {
	Status   string `json:"status"`
	Sip      string `json:"sip"`
	IsOnline bool   `json:"is_online"`
	Message  string `json:"message"`
}

// CatchSipCallerid  https://novofon.com/instructions/api/#api_sip_callerid_nvf
type CatchSipCallerid struct {
	Status      string `json:"status"`
	Sip         string `json:"sip"`
	NewCallerId bool   `json:"new_caller_id"`
	Message     string `json:"message"`
}

// CatchSipRedirection https://novofon.com/instructions/api/#api_sip_redirection_get_nvf
type CatchSipRedirection struct {
	Status string `json:"status"`
	Info   []struct {
		SipId            string `json:"sip_id"`
		Status           string `json:"status"`
		Condition        string `json:"condition"`
		Destination      string `json:"destination"`
		DestinationValue string `json:"destination_value"`
	} `json:"info"`
	Message string `json:"message"`
}

// CatchSipRedirectionPUT https://novofon.com/instructions/api/#api_sip_redirection_put
type CatchSipRedirectionPUT struct {
	Status        string `json:"status"`
	Sip           string `json:"sip"`
	CurrentStatus string `json:"current_status"`
	Destination   string `json:"destination"`
	Message       string `json:"message"`
}

// CatchSipCreate https://novofon.com/instructions/api/#api_sip_post_create
type CatchSipCreate struct {
	Status  string `json:"status"`
	Sip     string `json:"sip"`
	Message string `json:"message"`
}

// CatchStatistics https://novofon.com/instructions/api/#api_statistic_nvf
type CatchStatistics struct {
	Status string `json:"status"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Stats  []struct {
		Id          string  `json:"id"`
		Sip         string  `json:"sip"`
		Callstart   string  `json:"callstart"`
		From        int     `json:"from"`
		To          int     `json:"to"`
		Description string  `json:"description"`
		Disposition string  `json:"disposition"`
		Billseconds int     `json:"billseconds"`
		Cost        float32 `json:"cost"`
		Billcost    float32 `json:"billcost"`
		Currency    string  `json:"currency"`
		Seconds     int     `json:"seconds"`
	}
	Message string `json:"message"`
}
