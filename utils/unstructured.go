// Obtain chunks from Unstructured and convert to Golang structures
package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"

	"github.com/carlmjohnson/requests"
)

// Structure for JSON demarshalling (string to data structure)
type responseFormat struct {
	Type string
	ElementID string
	Text string
	Metadata map[string]interface{}
}

func Partition(filename string) []responseFormat {
	jsonString := getResponse(filename)
	jsonResps := demarshalResponse(jsonString)

	return jsonResps
}

func getResponse(filename string) string {

	// Creating the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Need to upload file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create a form field
	err = writer.WriteField("strategy", "hi_res")
	if err != nil{
		panic(err)
	}

	part, err := writer.CreateFormFile("files", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	var s string
	var j responseFormat

	err = requests.
		URL("https://api.unstructured.io/general/v0/general").
		Accept("application/json").
		ContentType(writer.FormDataContentType()).
		Header("unstructured-api-key", "3TzVzYZI6Dh7r5FTlOyK3aB5xzagmZ").
		BodyReader(&requestBody).
		ToJSON(&j).
		ToString(&s).
		Fetch(context.Background())

	if err != nil{
		panic(err)
	}

	return s
	
}

func demarshalResponse(s string) []responseFormat {
	count := 1
	var json_resps []responseFormat
	var json_string string
	json_string += "{"

	// Uses stack approach to determine end of string
	var letter string
	var response responseFormat
	i := 2
	for ;i < len(s); {
		for ;i < len(s) && count != 0; {
			letter = string(s[i])
			switch letter {
				case "{":
					count += 1
				case "}":
					count -= 1
			}
			json_string += letter
			i++
		}

		// Demarshalling JSON string
		err := json.Unmarshal([]byte(json_string), &response)
		if err != nil {
			panic(err)
		}
		json_resps = append(json_resps, response)

		// Resetting for next JSON string to parse
		json_string = "{"
		count = 1
		i += 2 // skipping comma and next starting bracket
	}

	return json_resps
}