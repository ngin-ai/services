package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type MyData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(ctx context.Context) ([]MyData, error) {
	// Load the Google Sheets API credentials from the credentials file
	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		return nil, err
	}

	// Create a new Sheets service using the credentials
	srv, err := sheets.NewService(ctx, sheets.WithCredentials(creds))
	if err != nil {
		return nil, err
	}

	// Specify the spreadsheet ID and range to read from
	spreadsheetID := "<your spreadsheet ID here>"
	readRange := "Sheet1!A2:B"

	// Read the data from the spreadsheet
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, err
	}

	// Convert the response to a slice of MyData structs
	var data []MyData
	for _, row := range resp.Values {
		if len(row) == 2 {
			item := MyData{
				Name:  row[0].(string),
				Email: row[1].(string),
			}
			data = append(data, item)
		}
	}

	// Return the data as JSON
	return data, nil
}

func main() {
	lambda.Start(handler)
}
