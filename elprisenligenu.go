package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const ELPRISENLIGENU_URL = "https://elprisenligenu.dk/api/v1/"
const ELPRISENLIEGNU_PRICE_CLASS = "DK2"

type ElprisenLigeNuPricePoint struct {
	DKK       float32   `json:"DKK_per_kWh"`
	EUR       float32   `json:"EUR_per_kWh"`
	EXR       float32   `json:"EXR"`
	TimeStart time.Time `json:"time_start"`
}

type HTTPGetter interface {
	Get(url string) (*http.Response, error)
}

type ElPrisenLigeNuAPI struct {
	Getter HTTPGetter
}

func NewElPrisenLigeNuAPI(getter HTTPGetter) *ElPrisenLigeNuAPI {
	if getter == nil {
		getter = http.DefaultClient
	}
	return &ElPrisenLigeNuAPI{
		Getter: getter,
	}
}

func (api *ElPrisenLigeNuAPI) GetPrices(day time.Time) ([]ElprisenLigeNuPricePoint, error) {
	url := ELPRISENLIGENU_URL + "prices/" + day.Format("2006/01-02") + "_" + ELPRISENLIEGNU_PRICE_CLASS + ".json"
	resp, err := api.Getter.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get price data (path=%s)", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("price data request failed with unexpected status code (code=%d)", resp.StatusCode)
	}

	var pricePoints []ElprisenLigeNuPricePoint
	if err := json.NewDecoder(resp.Body).Decode(&pricePoints); err != nil {
		return nil, errors.Wrap(err, "failed to decode price data")
	}

	return pricePoints, nil
}
