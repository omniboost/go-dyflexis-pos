package dyflexis_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDyflexisPOSPost(t *testing.T) {
	b := []byte(`
{
  "BusinessDate": "2022-06-07T00:00:00.000000Z",
  "Reports": {
    "Turnover": {
      "turnover": 121.0,
      "tax": 21.0
    },
    "Hourly": [
      {
        "time": "2022-06-07T00:00:00.000000Z",
        "turnover": 121.0,
        "tax": 21.0
      },
      {
        "time": "2022-06-07T01:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T02:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T03:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T04:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T05:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T06:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T07:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T08:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T09:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T10:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T11:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T12:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T13:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T14:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T15:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T16:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T17:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T18:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T19:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T20:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T21:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T22:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      },
      {
        "time": "2022-06-07T23:00:00.000000Z",
        "turnover": 0.0,
        "tax": 0.0
      }
    ],
    "Departments": [
    ]
  }
}`)
	req := client.NewDyflexisPOSPostRequest()
	json.Unmarshal(b, req.RequestBody())
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
