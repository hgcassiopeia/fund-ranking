/*
Copyright © 2022 THITIKAN PHAGAMAS <thitikan.phagamas@gmail.com>

*/
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallApiData_Success(t *testing.T) {
	baseAPI := "https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y"
	responseData := callApiData(baseAPI)
	assert.True(t, len(responseData) > 0, "Call API successfully.")
}

func TestCallApiData_Fail(t *testing.T) {
	baseAPI := "https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/"
	responseData := callApiData(baseAPI)
	assert.True(t, len(responseData) == 0, "Wrong URI cannot call API.")
}

func TestGetFundByRange(t *testing.T) {
	timeRange := "1M"
	isSuccess := getFundByRange(timeRange)
	assert.True(t, isSuccess, "Can get fund by range.")
}
