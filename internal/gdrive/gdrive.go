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
package gdrive

import (
	"bytes"
	"context"
	"log"
	"net/http"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func Service(client *http.Client) (*drive.Service, error) {

	// Connect client to drive API
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	return srv, err
}

func CreateFolder(srv *drive.Service, path string) ([]string, error) {

	folder := &drive.File{
		Name:     path,
		MimeType: "application/vnd.google-apps.folder",
	}
	results, err := srv.Files.Create(folder).Do()
	if err != nil {
		return nil, err
	}
	folderId := []string{results.Id}
	return folderId, err

}

func CreateFile(srv *drive.Service, folderId []string, filename string, file *bytes.Reader) {

	dfile := &drive.File{
		Name:     filename,
		Parents:  folderId,
		MimeType: "audio/mp3",
	}

	_, err := srv.Files.Create(dfile).Media(file).Do()
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
}
