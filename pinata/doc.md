# How to use Bifrost with Pinata Cloud

Welcome to the Bifrost documentation for Pinata Cloud! In this guide, we'll show you how to use Bifrost to upload files to Pinata Cloud.

## Overview

Pinata Cloud is a powerful decentralized cloud storage platform that enables users to securely store and manage digital assets. It allows you to store and access your data on IPFS. Bifrost provides a simple and intuitive way to upload files to Pinata Cloud without having to write complex code.

## Prerequisites

Before you can start using Bifrost to upload files to Pinata Cloud, you'll need to make sure you have the following:

- A Pinata Cloud account with API access
- A Pinata Cloud JWT (we hope to add support for secret/key soon)
- Bifrost installed on your local machine

## Mount a Bifrost bridge to Pinata

1. Install Bifrost using: `go get github.com/bifrost-cloud/bifrost`
2. Initialize a new Bifrost client and mount a Pinata Cloud bridge using the following code:

```go
package main

import (
	"fmt"
	"os"
	"github.com/funmi4194/bifrost"
)
// mount a bridge to Pinata
pinataBridge, _ := bifrost.NewRainbowBridge(&bifrost.BridgeConfig{
	Provider:    bifrost.PinataCloud,
	PinataJWT:   os.Getenv("PINATA_JWT"),
	EnableDebug: true,
	PublicRead:  true,
})
defer pinataBridge.Disconnect()
fmt.Printf("Connected to %s\n", pinataBridge.Config().Provider)
```

And that's it! You have now mounted a Bifrost bridge to your Pinata account and can start uploading files via this bridge.

## Shipping a file to Pinata Cloud via the rainbow bridge

Uploading a file to Pinata Cloud with Bifrost is just as easy. Here's how:

```go
// Upload a file
uploadedFile, err := bridge.UploadFile(bifrost.File{
	Path:     "../shared/image/aand.png",
	Filename: "pinata_aand.png",
	Options: map[string]interface{}{
		bifrost.OptPinata: map[string]interface{}{
			"cidVersion": 1,
		},
		bifrost.OptMetadata: map[string]string{
			"originalname": "aand.png",
		},
	},
})
if err != nil {
	fmt.Println(err)
	return
}
fmt.Printf("Uploaded file: %s to %s\n", uploadedFile.Name, uploadedFile.Preview)
```

As you can see, uploading a file to Pinata Cloud using Bifrost is as simple as calling the UploadFile method on the Bifrost client with the path to the file on your local machine and the name to give the file on Pinata Cloud, and any other metatadata via the `Options` field

## Uploading Multiple Files to Pinata Cloud with Bifrost

Bifrost also provides a simple way to upload multiple files to Pinata Cloud. Here's an example code snippet:

```go
// Upload multiple files

f, _ := os.Open("../shared/image/hair.jpg")

uploadedFiles, err := bridge.UploadMultiFile(bifrost.MultiFile{
	Files: []bifrost.File{
		{
			Path:     "../shared/image/aand.png",
			Filename: "a_and_ampersand.png",
			Options: map[string]interface{}{
				bifrost.OptMetadata: map[string]string{
					"originalname": "aand.png",
				},
			},
		},
		{
			Path:     "../shared/image/bifrost.webp",
			Filename: "bifrost_bridge.webp",
			Options: map[string]interface{}{
				bifrost.OptMetadata: map[string]string{
					"originalname": "bifrost.jpg",
					"universe":     "Marvel",
				},
			},
		},
        {
            Path:     "",
            Handle:   f,
            Filename: "sammy.jpg",
            Options: map[string]interface{}{
                bifrost.OptMetadata: map[string]string{
                    "originalname": "hair.jpg",
                    "specie":       "Human",
                },
                bifrost.OptACL: bifrost.ACLPublicRead,
            },
        },
	},
	GlobalOptions: map[string]interface{}{
		bifrost.OptACL: bifrost.ACLPrivate,
	},
})
if err != nil {
	fmt.Println(err)
	return
}

for _, file := range uploadedFiles {
	fmt.Printf("Uploaded file: %s to %s\n", file.Name, file.Preview)
}
```

## Additional Resources

- [Pinata Cloud Documentation](https://pinata.cloud/documentation)
- [What is decentralized storage](https://www.enterprisenetworkingplanet.com/data-center/decentralized-data-storage)
- [What is IPFS](https://docs.ipfs.tech/concepts/what-is-ipfs)

We hope this guide has been helpful in using Bifrost with Pinata Cloud. If you have any questions or feedback, please don't hesitate to open an [issue](https://github.com/funmi4194/bifrost/issues)!
