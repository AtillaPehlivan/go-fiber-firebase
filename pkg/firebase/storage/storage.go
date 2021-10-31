package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
	"mocerize-api/pkg/firebase"
)

var (
	defaultBucket *storage.BucketHandle
)

type FirebaseStorage struct {
}

func Setup(ctx context.Context) error {
	client, err := firebase.GetFirebaseApp().Storage(ctx)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket("mockerize.appspot.com")
	if err != nil {
		return err
	}

	defaultBucket = bucket

	return nil
}

func DefaultBucket() *storage.BucketHandle {
	return defaultBucket
}

func (fs *FirebaseStorage) UploadSingle() {

	// Open local file.
	//f, err := os.Open("notes.txt")
	//if err != nil {
	//	log.Printf("os.Open: %v\n", err)
	//}
	//defer f.Close()
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancel()
	//
	//wc := defaultBucket.Object("atilla.txt").NewWriter(ctx)
	//
	//if _, err = io.Copy(wc, f); err != nil {
	//	log.Printf("io.Copy: %v\n\n", err)
	//}
	//
	//
	//if err := wc.Close(); err != nil {
	//	log.Printf("Writer.Close: %v\n", err)
	//}
	//log.Printf("Blob %v uploaded.\n\n", "atilla")

	//https://firebasestorage.googleapis.com/v0/b/mockerize.appspot.com/o/atilla.txt?alt=media

	attr, _ := defaultBucket.Object("atilla.txt").Attrs(context.Background())
	log.Println(attr.MediaLink)
	for key, val := range attr.Metadata {
		log.Println(key, val)

	}

}
