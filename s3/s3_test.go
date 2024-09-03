package s3_test

import (
	"os"
	"testing"

	"github.com/funmi4194/bifrost"
)

var (
	bridge bifrost.RainbowBridge
	err    error

	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_KEY")
	AWS_BUCKET_NAME       = os.Getenv("AWS_BUCKET_NAME")
)

func setup(t *testing.T) {

	bridge, err = bifrost.NewRainbowBridge(&bifrost.BridgeConfig{
		DefaultBucket:  AWS_BUCKET_NAME,
		DefaultTimeout: 10,
		Provider:       bifrost.SimpleStorageService,
		EnableDebug:    true,
		PublicRead:     true,
		AccessKey:      AWS_ACCESS_KEY_ID,
		SecretKey:      AWS_SECRET_ACCESS_KEY,
		Region:         "ap-northeast-1",
	})
	if err != nil {
		t.Error(err.(bifrost.Error).Code(), err)
		return
	}

	t.Logf("Connected to %s\n", bridge.Config().Provider)
}

func teardown() {
	bridge.Disconnect()
}

func TestS3(t *testing.T) {
	setup(t)
	defer teardown()

	t.Run("Tests UploadFile method", func(t *testing.T) {
		o, err := bridge.UploadFile(bifrost.File{
			Path:     "../shared/image/aand.png",
			Filename: "a_and_ampersand.png",
			Options: map[string]interface{}{
				bifrost.OptMetadata: map[string]string{
					"originalname": "aand.png",
				},
			},
		})
		if err != nil {
			t.Errorf("Failed to upload file: %v", err)
			return
		}
		t.Logf("Uploaded file: %s to %s\n", o.Name, o.Preview)
	})

	t.Run("Tests UploadMultiFile method", func(t *testing.T) {
		o, err := bridge.UploadMultiFile(bifrost.MultiFile{
			Files: []bifrost.File{
				{
					Path:     "../shared/image/aand.png",
					Filename: "a_and_ampersand.png",
					Options: map[string]interface{}{
						bifrost.OptMetadata: map[string]string{
							"originalname": "aand.png",
						},
						bifrost.OptACL: bifrost.ACLPublicRead,
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
			},

			// say 3 of 4 files need to share the same option, you can set globally for those 3 files and set the 4th file's option separately, bifrost won't override the option
			GlobalOptions: map[string]interface{}{
				bifrost.OptACL: bifrost.ACLPrivate,
			},
		})
		if err != nil {
			t.Errorf("Failed to upload file: %v", err)
			return
		}

		for _, file := range o {
			t.Logf("Uploaded file: %s to %s\n", file.Name, file.Preview)
		}
	})

}
