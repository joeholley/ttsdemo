/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package gsheets

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
)

func Service(client *http.Client) (*sheets.Service, error) {
	// Connect client to sheets API
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	//srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv, err
}

// gets data from specified cells
func RetrieveCells(srv *sheets.Service, spreadsheetId string, readRange string) (results []string, err error) {

	// Get data
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err == nil {
		if len(resp.Values) == 0 {
			fmt.Println("No data found.")
		} else {
			fmt.Printf("Retreived cells %s from %s\n", readRange, spreadsheetId)
			for _, row := range resp.Values {
				// first retrieved column is index 0
				results = append(results, row[0].(string))
			}
		}
	}
	return results, err
}

// puts data into cells (to track checksums, for example)
func WriteCells(srv *sheets.Service, spreadsheetId string, valueRange string, checksums []string) error {
	ctx := context.Background()

	// sheets.ValueRange requires an array of interface{}, so convert checksums to that format.
	csums := make([]interface{}, len(checksums))
	for i, v := range checksums {
		csums[i] = v
	}

	values := &sheets.ValueRange{
		MajorDimension: "COLUMNS",
		Values: [][]interface{}{
			csums,
		},
	}

	// write data
	resp, err := srv.Spreadsheets.Values.Update(spreadsheetId, valueRange, values).ValueInputOption("RAW").Context(ctx).Do()
	if err == nil {
		if resp.UpdatedCells == 0 {
			fmt.Println("No data found.")
		} else {
			fmt.Printf("Updated %v cells %s from %s\n", resp.UpdatedCells, valueRange, spreadsheetId)
		}
	}
	return err
}
