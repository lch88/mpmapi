package mpmapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// Manager is an interface to access MoPub Publisher Management API
type Manager interface {
	GetAllLineItems() ([]LineItem, error)
	GetAllAdUnits() ([]AdUnit, error)
}

// Client is a REST API Client to access MoPub Publisher Management API
// package users do not need direct access to this.
type Client struct {
	apiKey  string
	baseURL string
}

// NewClient returns a MoPub Publisher Management API Client with given API Key
func NewClient(apiKey string) Manager {
	return NewClientWithBaseURL(apiKey, "https://api.mopub.com/v2")
}

// NewClientWithBaseURL returns a MoPub Publisher Management API Client with given API Key and base url
func NewClientWithBaseURL(apiKey string, baseURL string) Manager {
	return Client{
		apiKey, baseURL,
	}
}

func (c Client) GetAllLineItems() ([]LineItem, error) {
	data, err := c.makeGetAllItemsRequest("line-items", "")
	if err != nil {
		return nil, err
	}

	lineItems := make([]LineItem, 0)
	for _, datum := range data {
		var lineItem LineItem
		if err := json.Unmarshal(datum, &lineItem); err != nil {
			return nil, err
		}
		lineItems = append(lineItems, lineItem)
	}

	return lineItems, nil
}

func (c Client) GetAllAdUnits() ([]AdUnit, error) {
	data, err := c.makeGetAllItemsRequest("adunits", "")
	if err != nil {
		return nil, err
	}

	adUnits := make([]AdUnit, 0)
	for _, datum := range data {
		var adUnit AdUnit
		if err := json.Unmarshal(datum, &adUnit); err != nil {
			return nil, err
		}
		adUnits = append(adUnits, adUnit)
	}

	return adUnits, nil
}

func (c Client) makeGetAllItemsRequest(resource string, filter string) ([]json.RawMessage, error) {

	result := make([]json.RawMessage, 0)
	resp, err := c.makeGetItemsRequest(resource, filter, 1)
	if err != nil {
		return nil, err
	}
	result = append(result, resp.Data...)

	g, _ := errgroup.WithContext(context.Background())
	lastPage := resp.Pagination.LastPage

	for i := 2; i <= lastPage; i++ {
		i := i
		time.Sleep(300 * time.Millisecond)
		g.Go(func() error {
			resp, err := c.makeGetItemsRequest(resource, filter, int64(i))

			if err != nil {
				return err
			}

			result = append(result, resp.Data...)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}

func (c Client) makeGetItemsRequest(resource string, filter string, page int64) (*Response, error) {
	body, err := c.makeRawAPIRequest(resource, filter, page)
	if err != nil {
		return nil, err
	}

	resp := Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c Client) makeRawAPIRequest(path string, filter string, page int64) ([]byte, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page":  strconv.FormatInt(page, 10),
			"limit": "500",
		}).
		SetHeader("x-api-key", c.apiKey).
		SetError(&ErrorResponse{}).
		Get(fmt.Sprintf("%s/%s?%s", c.baseURL, path, filter))

	if err != nil {
		// _ = ioutil.WriteFile(fmt.Sprintf("%d.json", page), resp.Body(), 0644)
		return nil, err
	}

	if err := resp.Error(); err != nil {
		errResp := err.(*ErrorResponse)
		errMsgs := make([]string, 0)
		for _, v := range errResp.Errors {
			errMsgs = append(errMsgs, fmt.Sprintf("%s, %s", v.Type, v.Message))
		}
		return nil, errors.Errorf("statusCode: %d, msgs: %s", errResp.StatusCode, strings.Join(errMsgs, ";"))
	}

	return resp.Body(), nil
}
