package tenderduty

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/blockpane/tenderduty/v2/pkg/namada"
	"github.com/near/borsh-go"
)

func TestXxx(t *testing.T) {
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(`http://222.106.187.14:55102/abci_query?path="/vp/pos/pos_params"`)
	if err != nil {
		t.Logf("err: %s", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Logf("err: %s", err)
		return
	}

	type ABCIResponse struct {
		Result struct {
			Response struct {
				Code  int    `json:"code"`
				Value string `json:"value"`
			} `json:"response"`
		} `json:"result"`
	}

	var result ABCIResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		t.Logf("err: %s", err)
		return
	}
	decodedBz, _ := base64.StdEncoding.DecodeString(result.Result.Response.Value)
	t.Logf("%s", result.Result.Response.Value)
	t.Logf("%X", decodedBz)

	posParams := namada.PosParams{}
	err = borsh.Deserialize(&posParams, decodedBz)
	if err != nil {
		t.Logf("err: %s", err)
		return
	}

	t.Logf("%+v", posParams)

	// val := &staking.QueryValidatorResponse{}
	// err = val.Unmarshal(decodedBz)
	// if err != nil {
	// 	t.Logf("err: %s", err)
	// 	return
	// }

	// t.Logf("%v", val.GetValidator())

}
