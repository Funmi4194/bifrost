package request

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"github.com/funmi4194/bifrost/shared/config"
	"github.com/funmi4194/bifrost/shared/types"
)

func (c *Client) PostForm(url string, params types.Param) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, pf := range params.Files {
		if pf.Path != "" {
			// open file
			file, err := os.Open(pf.Path)
			if err != nil {
				return nil, err
			}
			pf.Handle = file
			// close file
			defer file.Close()
		}

		part, err := writer.CreateFormFile(pf.Key, pf.Name)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, pf.Handle)
		if err != nil {
			return nil, err
		}
	}

	for _, pd := range params.Data {
		_ = writer.WriteField(pd.Key, pd.Value)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// copy request
	req := c.Request.Clone(c.Request.Context())
	req.Method = config.MethodPost
	req.URL, _ = c.Request.URL.Parse(url)
	req.Header.Add(config.ReqContentType, writer.FormDataContentType())
	req.Body = io.NopCloser(body)

	// make request
	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response
	return io.ReadAll(resp.Body)
}
