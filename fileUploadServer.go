package main

import (
	"fmt"
	"net/http"

	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {
	store := filestore.FileStore{
		Path: "./uploads",
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:                "/files/",
		StoreComposer:           composer,
		NotifyCreatedUploads:    true,
		NotifyCompleteUploads:   true,
		NotifyTerminatedUploads: false,
		NotifyUploadProgress:    true,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			fmt.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()

	go func() {
		for {
			event := <-handler.CreatedUploads
			fmt.Printf("Upload %s created\n", event.Upload.MetaData["filename"])
		}
	}()

	go func() {
		for {
			event := <-handler.UploadProgress
			fi := event.Upload
			fmt.Printf("Uploaded %0.f%%\n", float64(fi.Offset)/float64(fi.Size)*100)
		}
	}()

	http.Handle("/files/", http.StripPrefix("/files/", handler))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}
}
