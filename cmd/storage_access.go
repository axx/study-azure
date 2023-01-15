package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	fmt.Println("Azure Blob storage access study")

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	blobServiceUrl := "https://ststudy.blob.core.windows.net"
	client, err := azblob.NewClient(blobServiceUrl, credential, nil)
	handleError(err)

	suffix := randomString()

	containerName := fmt.Sprintf("container-%s", suffix)
	fmt.Println("Creating a container", containerName)
	_, err = client.CreateContainer(context.TODO(), containerName, nil)
	handleError(err)

	blobName := fmt.Sprintf("myblob-%s.txt", suffix)
	file, err := os.Create(blobName)
	handleError(err)
	defer os.Remove(blobName)
	defer file.Close()

	_, err = fmt.Fprintln(file, "this goes to blob", blobName, "in container", containerName)
	handleError(err)

	fmt.Println("Uploading file", blobName, "to container", containerName)
	_, err = client.UploadFile(context.TODO(), containerName, blobName, file, nil)
	handleError(err)
}
