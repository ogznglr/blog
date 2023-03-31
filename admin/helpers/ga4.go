package helpers

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

func Ga4Request(pagepath string) (string, error) {

	propertyid := "359945658"
	property := fmt.Sprintf("properties/%s", propertyid)

	// Read the credentials file
	creds, err := ioutil.ReadFile("./quickstart.json")
	if err != nil {
		fmt.Printf("Unable to read credentials file: %v", err)
		return "", err
	}

	// Parse the credentials file
	config, err := google.JWTConfigFromJSON(creds, analyticsdata.AnalyticsReadonlyScope)

	if err != nil {
		fmt.Printf("Unable to parse credentials file: %v", err)
		return "", err
	}

	// Create the HTTP client with the credentials
	client := config.Client(context.Background())

	// Create an Analytics Data API service object
	service, err := analyticsdata.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return "", err
	}

	// Make a request to the API
	req := &analyticsdata.RunReportRequest{
		Property: property,
		DateRanges: []*analyticsdata.DateRange{
			{
				StartDate: "30daysAgo",
				EndDate:   "today",
			},
		},
		Dimensions: []*analyticsdata.Dimension{
			{
				Name: "pagePath",
			},
		},

		Metrics: []*analyticsdata.Metric{
			{
				Name: "activeUsers",
			},
		},
		DimensionFilter: &analyticsdata.FilterExpression{
			Filter: &analyticsdata.Filter{
				FieldName: "pagePath",
				StringFilter: &analyticsdata.StringFilter{
					CaseSensitive: true,
					Value:         pagepath,
				},
			},
		},
	}

	resp, err := service.Properties.RunReport(req.Property, req).Do()
	if err != nil || resp.RowCount == 0 {
		return "", err
	}

	count := resp.Rows[0].MetricValues[0].Value
	return count, nil

}
