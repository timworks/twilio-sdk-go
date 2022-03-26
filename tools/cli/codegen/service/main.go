package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	apiclient "github.com/timworks/twilio-sdk-go-tools/cli/codegen/service/api_client"
	apioperation "github.com/timworks/twilio-sdk-go-tools/cli/codegen/service/api_operation"
	client "github.com/timworks/twilio-sdk-go-tools/cli/codegen/service/client"

	"github.com/iancoleman/strcase"
	"github.com/santhosh-tekuri/jsonschema/v2"
)

func main() {
	var definitionPath string
	var outputPath string

	flag.StringVar(&definitionPath, "definition", "",
		"The `path` to client definition",
	)

	flag.StringVar(&outputPath, "target", "",
		"The target `path` to generate the api into",
	)

	flag.Parse()

	_, filename, _, _ := runtime.Caller(0)
	schema, _ := jsonschema.Compile(fmt.Sprintf("%s/schema.json", path.Dir(filename)))
	fixture, _ := ioutil.ReadFile(fmt.Sprintf("%s/api.json", definitionPath))

	var api map[string]interface{}
	json.Unmarshal(fixture, &api)

	if err := schema.ValidateInterface(api); err != nil {
		fmt.Println(err)
		return
	}

	if err := translateAndGenerateClient(outputPath, api); err != nil {
		fmt.Println(err)
		return
	}

	if err := generateApiClients(api["subClients"].([]interface{}), api["structures"].(map[string]interface{}), outputPath); err != nil {
		fmt.Println(err)
		return
	}
}

func generateApiClients(apiClients []interface{}, structures map[string]interface{}, path string) error {
	for _, apiClient := range apiClients {
		apiClientMap := apiClient.(map[string]interface{})

		filePath := fmt.Sprintf("%s/%s", path, apiClientMap["packageName"].(string))

		if apiClientMap["subClients"] != nil {
			if err := generateApiClients(apiClientMap["subClients"].([]interface{}), structures, filePath); err != nil {
				return err
			}
		}

		if err := translateAndGenerateApiClient(filePath, apiClientMap); err != nil {
			return err
		}

		for _, operation := range apiClientMap["operations"].([]interface{}) {
			operationMap := operation.(map[string]interface{})
			operationMap["packageName"] = apiClientMap["packageName"]
			operationMap["config"] = apiClientMap["config"]
			operationMap["structures"] = structures

			bytes, err := json.Marshal(operation)
			if err != nil {
				return err
			}
			operationResp, err := apioperation.Translate(bytes)
			if err != nil {
				return err
			}

			contents, err := apioperation.Generate(operationResp, false)
			if err != nil {
				return err
			}
			if err := createAndWriteFile(filePath, fmt.Sprintf("api_op_%s.go", strcase.ToSnake(operationMap["name"].(string))), string(*contents)); err != nil {
				return err
			}
		}
	}
	return nil
}

func translateAndGenerateApiClient(path string, content map[string]interface{}) error {
	bytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	translationResp, err := apiclient.Translate(bytes)
	if err != nil {
		return err
	}
	contents, err := apiclient.Generate(translationResp, false)
	if err != nil {
		return err
	}
	if err := createAndWriteFile(path, "api_op_client.go", string(*contents)); err != nil {
		return err
	}
	return nil
}

func translateAndGenerateClient(path string, content map[string]interface{}) error {
	bytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	translationResp, err := client.Translate(bytes)
	if err != nil {
		return err
	}
	contents, err := client.Generate(translationResp, false)
	if err != nil {
		return err
	}
	if err := createAndWriteFile(path, "client.go", string(*contents)); err != nil {
		return err
	}
	return nil
}

func createAndWriteFile(path string, fileName string, content string) error {
	os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	defer file.Close()

	if err != nil {
		return fmt.Errorf("Unable to create file on disk. %s", err)
	}

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("Unable to write to file. %s", err)
	}
	return nil
}
