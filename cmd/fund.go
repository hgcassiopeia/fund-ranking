/*
Copyright © 2022 THITIKAN PHAGAMAS <thitikan.phagamas@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// fundCmd represents the fund command
var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Fund is a command for displays a list of funds.",
	Long: `
Fund is a command for displays a list of funds over a specified time period.
Use flag 'time' for setup range such as 1d, 1w, 1m and 1y (1y is default)
For Example : 

codetest fund --time=1d`,
	Run: func(cmd *cobra.Command, args []string) {
		time, _ := cmd.Flags().GetString("time")

		correctTime := contains([]string{"1D", "1W", "1M", "1Y"}, strings.ToUpper(time))
		if correctTime {
			time := strings.ToUpper(time)
			getFundByRange(time)
		} else {
			log.Printf("Your input: %v is incorrect type please input time contain 1D, 1W, 1M, 1Y", time)
		}
	},
}

func init() {
	rootCmd.AddCommand(fundCmd)
	fundCmd.PersistentFlags().String("time", "1Y", "A flag for get list of funds by range in 1D, 1W, 1M, 1Y for example \n codetest fund --time=1d \t 1d is range of day \n codetest fund --time=1w \t 1w is range of week \n codetest fund --time=1m \t 1m is range of month \n codetest fund --time=1y \t 1y is range of year")
}

func getFundByRange(timeRange string) bool {

	isSuccess := false
	url := "https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-"
	url += timeRange + ".json"

	responseData := callApiData(url)
	if len(responseData) > 0 {
		fundRaw := Fund{}

		if err := json.Unmarshal(responseData, &fundRaw); err != nil {
			log.Printf("Could not unmarshal responseData - %v", err)
		}

		fund := []Fund{}
		if err := json.Unmarshal(fundRaw.Data, &fund); err != nil {
			log.Printf("Could not unmarshal fundRaw data - %v", err)
		}

		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "ชื่อกองทุน"},
				{Align: simpletable.AlignCenter, Text: "อันดับของกองทุน"},
				{Align: simpletable.AlignCenter, Text: "เวลาที่ข้อมูลถูกอัพเดต"},
				{Align: simpletable.AlignCenter, Text: "ผลตอบแทน"},
				{Align: simpletable.AlignCenter, Text: "ราคา"},
			},
		}

		for _, row := range fund {
			t, _ := time.Parse("2006-01-02T00:00:00.000Z", row.NavDate)
			r := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: row.ThailandFundCode},
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.NavReturn)},
				{Align: simpletable.AlignCenter, Text: t.Format("02 Jan 2006")},
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.AvgReturn)},
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%.2f", row.Nav)},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleDefault)
		fmt.Println(table.String())
		isSuccess = true
	}

	return isSuccess
}

func callApiData(baseAPI string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not request fund data - %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Codetes CLI (github.com/hgcassiopeia/codetest)")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read response body - %v", err)
	}

	return responseData
}

type Fund struct {
	Status           bool            `json:"status"`
	Error            string          `json:"error"`
	Data             json.RawMessage `json:"data"`
	MStarID          string          `json:"mstar_id"`
	ThailandFundCode string          `json:"thailand_fund_code"`
	NavReturn        float64         `json:"nav_return"`
	Nav              float64         `json:"nav"`
	NavDate          string          `json:"nav_date"`
	AvgReturn        float64         `json:"avg_return"`
}

// Helper function
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
