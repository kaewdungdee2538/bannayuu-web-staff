package googlestorage

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func cloudStorage() {
	// [START cloud_storage_golang]
	config := &firebase.Config{
		StorageBucket: "gs://bannayuu-admin.appspot.com",
	}
	opt := option.WithCredentialsFile("googlestorage/bannayuu-admin-firebase-adminsdk-zhr8k-0f2216c5e5.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.
	// [END cloud_storage_golang]

	log.Printf("Created bucket handle: %v\n", bucket)
}


func CloudStorageCustomBucket(app *firebase.App) {
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// [START cloud_storage_custom_bucket_golang]
	bucket, err := client.Bucket("my-custom-bucket")
	// [END cloud_storage_custom_bucket_golang]
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Created bucket handle: %v\n", bucket)
}