package mpmapi

import (
	"encoding/json"
	"time"
)

type Pagination struct {
	Count       int64  `json:"count"`
	CurrentPage int64  `json:"currentPage"`
	PerPage     int64  `json:"perPage"`
	PrevPage    *int64 `json:"prevPage"`
	NextPage    *int64 `json:"nextPage"`
	LastPage    int    `json:"lastPage"`
}

type LineItem struct {
	AdUnitKeys           []string `json:"adUnitKeys"`
	Advertiser           string   `json:"advertiser"`
	AllocationPercentage float64  `json:"allocationPercentage"`
	AutoCpm              float64  `json:"autoCpm"`
	Bid                  float64  `json:"bid"`
	Budget               float64  `json:"budget"`
	BudgetStrategy       string   `json:"budgetStrategy"`
	BudgetType           string   `json:"budgetType"`
	DayParts             []struct {
		Days      []string `json:"days"`
		StartTime int      `json:"startTime"`
		EndTime   int      `json:"endTime"`
	} `json:"dayParts"`
	DayPartTargeting  string    `json:"dayPartTargeting"`
	DeviceTargeting   bool      `json:"deviceTargeting"`
	DisallowAutoCpm   bool      `json:"disallowAutoCpm"`
	MaxAndroidVersion string    `json:"maxAndroidVersion"`
	MinAndroidVersion string    `json:"minAndroidVersion"`
	MaxIosVersion     string    `json:"maxIosVersion"`
	MinIosVersion     string    `json:"minIosVersion"`
	TargetAndroid     bool      `json:"targetAndroid"`
	TargetIos         bool      `json:"targetIos"`
	TargetIphone      bool      `json:"targetIphone"`
	TargetIpad        bool      `json:"targetIpad"`
	TargetIpod        bool      `json:"targetIpod"`
	End               time.Time `json:"end"`
	FrequencyCaps     []struct {
		Cap         int    `json:"cap"`
		Duration    string `json:"duration"`
		NumDuration int    `json:"numDuration"`
	} `json:"frequencyCaps"`
	FrequencyCapsEnabled         bool      `json:"frequencyCapsEnabled"`
	IncludeConnectivityTargeting string    `json:"includeConnectivityTargeting"`
	TargetedCarriers             []string  `json:"targetedCarriers"`
	IncludeGeoTargeting          string    `json:"includeGeoTargeting"`
	Key                          string    `json:"key"`
	Keywords                     []string  `json:"keywords"`
	Name                         string    `json:"name"`
	NetworkType                  string    `json:"networkType"`
	OrderKey                     string    `json:"orderKey"`
	OrderName                    string    `json:"orderName"`
	Priority                     int64     `json:"priority"`
	RefreshInterval              int64     `json:"refreshInterval"`
	Start                        time.Time `json:"start"`
	Status                       string    `json:"status"`
	TargetedCountries            []string  `json:"targetedCountries"`
	TargetedRegions              []string  `json:"targetedRegions"`
	TargetedCities               []string  `json:"targetedCities"`
	TargetedZipCodes             []string  `json:"targetedZipCodes"`
	Type                         string    `json:"type"`
	UserAppsTargeting            string    `json:"userAppsTargeting"`
	UserAppsTargetingList        []string  `json:"userAppsTargetingList"`
	EnableOverrides              bool      `json:"enableOverrides"`
	OverrideFields               struct {
		NetworkAppId     string `json:"network_app_id"`
		NetworkAdUnitId  string `json:"network_adunit_id"`
		NetworkAccountId string `json:"network_account_id"`
	} `json:"overrideFields"`
}

type Response struct {
	Data       []json.RawMessage `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

type LineItemsResponse struct {
	Data       []LineItem `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type AdUnit struct {
	AppKey              string `json:"appKey"`
	AppName             string `json:"appName"`
	AppType             string `json:"appType"`
	Active              bool   `json:"active"`
	DailyImpressionCap  int    `json:"dailyImpressionCap"`
	HourlyImpressionCap int    `json:"hourlyImpressionCap"`
	Format              string `json:"format"`
	Key                 string `json:"key"`
	Name                string `json:"name"`

	// For Banner
	RefreshInterval int `json:"refreshInterval"`

	// For Native
	NativePlacement       string            `json:"native_placement"`
	NativePositioningData NativePositioning `json:"native_positioning_data"`

	// For Rewarded
	RewardCallbackUrl string `json:"rewardCallbackUrl"`
	Rewards           []struct {
		Amount       int    `json:"amount"`
		CurrencyName string `json:"currencyName"`
	} `json:"rewards"`
}

type NativePositioning struct {
	Fixed []struct {
		Position int `json:"position"`
	} `json:"fixed"`
	Repeating struct {
		Interval int `json:"interval"`
	} `json:"repeating"`
}

type AdUnitsResponse struct {
	Data       []AdUnit   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ErrorResponse struct {
	StatusCode int `json:"statusCode"`
	Errors     []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"errors"`
}
