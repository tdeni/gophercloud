package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
	"github.com/gophercloud/gophercloud/v2/testhelper/fixture"
)

var (
	projectID = "{projectID}"
	resURL    = "/mgmt/" + "quotas/" + projectID
)

// getQuotasResp is a sample response to a Get call.
var getQuotasResp = `
{
    "quotas": [
        {
            "in_use": 5,
            "limit": 15,
            "reserved": 0,
            "resource": "instances"
        },
        {
            "in_use": 2,
            "limit": 50,
            "reserved": 0,
            "resource": "backups"
        },
        {
            "in_use": 1,
            "limit": 40,
            "reserved": 0,
            "resource": "volumes"
        }
    ]
}
`

// getQuotasResp is a sample response to a Update call.
var updateQuotaRequest = `
{
    "quotas": {
        "instances": 10,
        "backups": 30
    }
}
`

func HandleGet(t *testing.T) {
	fixture.SetupHandler(t, resURL, "GET", "", getQuotasResp, 200)
}

// HandleUpdateSuccessfully configures the test server to respond to an Update request.
func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(resURL,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, updateQuotaRequest)

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprint(w, updateQuotaRequest)
		})
}
