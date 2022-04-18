package discordwebhookfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func SendMessage(url string, message MessageFiles) error {

	client := &http.Client{}
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(message)

	if err != nil {
		return err
	}

	multipartbody := new(bytes.Buffer)
	w := multipart.NewWriter(multipartbody)

	if len(*message.Files) > 0 {

		for key, r := range *message.Files {
			r, err := os.Open(r)
			if err != nil {
				panic(err)
			}

			var fw io.Writer
			defer r.Close()
			if fw, err = w.CreateFormFile(fmt.Sprintf("file%v", key), r.Name()); err != nil {
				return err
			}
			if _, err = io.Copy(fw, r); err != nil {
				return err
			}

		}

		//#https:discord.com/developers/docs/reference#uploading-files

		_ = w.WriteField("payload_json", string(payload.String()))

		w.Close()

		req, err := http.NewRequest("POST", url, multipartbody)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		res, err := client.Do(req)
		if err != nil {
			return err
		}

		// Check the response
		if res.StatusCode != http.StatusOK {
			defer res.Body.Close()
			responseBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

			return fmt.Errorf(string(responseBody))

		}

	} else {

		resp, err := http.Post(url, "application/json", payload)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			defer resp.Body.Close()

			responseBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return fmt.Errorf(string(responseBody))
		}

	}

	return nil
}
