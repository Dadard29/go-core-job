package test

import (
	"github.com/Dadard29/go-core-job/connector"
	"testing"
)

func TestConnector(t *testing.T) {
	token := "protected"
	c := connector.NewCoreConnector("localhost", 8080, token)

	err := c.Job()
	if err != nil {
		t.Error(err)
	}
}
