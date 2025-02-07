package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/supabase-community/gotrue-go/types"
)

const recoverPath = "/recover"

// POST /recover
//
// Password recovery. Will deliver a password recovery mail to the user based
// on email address.
//
// By default recovery links can only be sent once every 60 seconds.
func (c *Client) Recover(req types.RecoverRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	apiURL := recoverPath
	if req.RedirectURL != nil {
		apiURL += "?redirect_to=" + *req.RedirectURL
	}

	r, err := c.newRequest(apiURL, http.MethodPost, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fullBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	return nil
}
