package jpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	appKey       string
	masterSecret string

	client *http.Client
}

var (
	_c *Client
)

func Init(key, secret string) {
	_c = &Client{
		appKey:       key,
		masterSecret: secret,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func Push(data *PushBody) (*PushResp, error) {
	bs, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", JPushPushUrl, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(_c.appKey, _c.masterSecret)

	resp, err := _c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: code=%d, error=%s", resp.StatusCode, string(respBs))
	}

	r := PushResp{}
	if err := json.Unmarshal(respBs, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func NewBasicPushBody(platform string, rid string, prod bool, msg string) *PushBody {
	return &PushBody{
		Platform: []string{platform},
		Audience: &Audience{
			RegistrationId: []string{rid},
		},
		Notification: &Notification{
			AIOpportunity: false,
			IOS: IOSBase{
				Alert: msg,
			},
		},
		InappMessage: &InappMessage{
			InappMessage: true,
		},
		Options: &Options{
			ApnsProduction: prod,
		},
	}
}
