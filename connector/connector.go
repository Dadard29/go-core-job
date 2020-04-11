package connector

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (c CoreConnector) Job() error {
	var resp *response
	var err error
	//if resp, err = c.checkUsersInactivity(); err != nil {
	//	return err
	//} else if !resp.status {
	//	return errors.New(resp.message)
	//}

	if resp, err = c.resetRequestCount(); err != nil {
		return err
	} else if !resp.Status {
		return errors.New(resp.Message)
	}

	return nil
}

func (c CoreConnector) unmarshal(resp *http.Response) (*response, error) {
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respObj response
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

func (c CoreConnector) checkUsersInactivity() (*response, error) {
	r, err := http.NewRequest(c.checkUserInactivityRoute.method,
		c.baseUrl + c.checkUserInactivityRoute.endpoint, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", c.protectedToken)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(resp)
}

func (c CoreConnector) resetRequestCount() (*response, error) {
	r, err := http.NewRequest(c.resetRequestCountRoute.method,
		c.baseUrl + c.resetRequestCountRoute.endpoint, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", c.protectedToken)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(resp)
}
