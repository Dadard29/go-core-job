package connector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (c CoreConnector) Job(weekday int) error {
	var resp *response
	var err error

	todayWeekday := time.Now().Weekday()
	if int(todayWeekday) != weekday {
		fmt.Println("wrong weekday")
		return nil
	}

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
