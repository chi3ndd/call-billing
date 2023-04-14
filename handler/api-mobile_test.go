package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestAPICallBilling(t *testing.T) {
	username := strconv.FormatInt(time.Now().UnixMilli(), 10)
	callDurations := [10]int64{}
	for idx, _ := range callDurations {
		callDurations[idx] = int64(rand.Intn(50000) + 1000)
	}
	client := resty.New()
	// Put
	for _, duration := range callDurations {
		res, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{"call_duration": duration}).
			Put(fmt.Sprintf("http://127.0.0.1:8910/mobile/%s/call", username))
		if err != nil {
			t.Fatal(err)
		}
		if !res.IsSuccess() {
			t.Fail()
		}
	}
	// Check billing
	res, err := client.R().Get(fmt.Sprintf("http://127.0.0.1:8910/mobile/%s/billing", username))
	if err != nil {
		t.Fatal(err)
	}
	type response struct {
		BlockCount int64 `json:"block_count"`
		CallCount  int64 `json:"call_count"`
	}
	var r response
	if err = json.Unmarshal(res.Body(), &r); err != nil {
		t.Fatal(err)
	}
	var expectBlockCount int64 = 0
	for _, duration := range callDurations {
		expectBlockCount += duration / 30000
		if duration%30000 > 0 {
			expectBlockCount += 1
		}
	}
	assert.Equal(t, expectBlockCount, r.BlockCount)
	assert.Equal(t, int64(len(callDurations)), r.CallCount)
}
