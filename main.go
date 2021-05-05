package main

import (
	"encoding/json"
	"flag"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"time"
	"vaccineNotifier/src"

	"github.com/sirupsen/logrus"
)

const DateFormat string = "02-01-2006"
const toPhone string = ""
const fromPhone string = ""
const twilioID string = ""
const twilioSecret string = ""
const cowinAPi string = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/calendarByDistrict"
const maxRetry int = 5

func main() {
	startDate := flag.String("s", time.Now().Format(DateFormat), "start date to run the script from")
	interval := flag.Int("i", 30, "number of days to look for")
	districtId := flag.Int("d", 294, "the id of the district to search for")

	flag.Parse()

	logger := logrus.Logger{}
	logger.Infof("looking from: %v for interval: %v", startDate, interval)

	var headers = map[string]string{
		"authority":       "cdn-api.co-vin.in",
		"credentials":     "include",
		"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0",
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "en-US,en;q=0.5",
		"referer":         "https://selfregistration.cowin.gov.in/",
	}

	sDate, err := time.Parse(DateFormat, *startDate)
	if err != nil {
		logrus.Errorf("unable to parse time")
		return
	}
	retries := 0
	for i := 0; i < *interval; {
		availableCenters := make([]src.Center, 0)
		centresWith18plus := make([]src.Center, 0)
		date := sDate.Add(time.Duration(i) * 24 * time.Hour)
		resp, err := src.DoSecureGet(cowinAPi, "", map[string]string{"date": date.Format(DateFormat),
			"district_id": strconv.Itoa(*districtId)}, headers)
		if err != nil {
			logrus.Errorf("Unable to fetch data for date: %v. Error: %v", err, date)
			if retries > maxRetry {
				i += 1
				retries = 0
				logrus.Errorf("Unable to fetch data for date: %v even after retries. Error: %v", err, date)
				continue
			}
			retries += 1
			time.Sleep(5 * time.Second)
			continue
		}
		centers := src.CenterList{}
		err = json.Unmarshal(resp, &centers)
		if err != nil {
			logger.Errorf("unable to fetch data. Error: %v", err)
			continue
		}
		for _, c := range centers.Centers {
			for _, s := range c.Sessions {
				if s.MinAgeLimit < 45 {
					centresWith18plus = append(centresWith18plus, c)
					if s.AvailableCapacity > 0 {
						availableCenters = append(availableCenters, c)
					}
				}
			}
		}
		logrus.Infof("Date %v: %v / %v relevant centers found", date, len(centresWith18plus), len(centers.Centers))
		if len(availableCenters) > 0 {
			spew.Dump(availableCenters)
		}
		time.Sleep(5 * time.Second)
		i += 1
		retries = 0
	}

}
