package lib

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func Build(serviceSchemaPath string, outputPath string) error {
	rawSchema, err := ioutil.ReadFile(serviceSchemaPath)
	if err != nil {
		return err
	}

	service := Service{}
	err = yaml.Unmarshal(rawSchema, &service)

	if err != nil {
		return fmt.Errorf("can't parse schema: %v/n", err)
	}

	typesFileText, err := buildTypesFile(&service)
	if err != nil {
		return err
	}

	handlerInterfaceFileText, err := buildHandlerInterfaceFile(&service)
	if err != nil {
		return err
	}

	executorFileText, err := buildExecutorFile(&service)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(outputPath, "types.go"), []byte(typesFileText), 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(outputPath, "handler_interface.go"), []byte(handlerInterfaceFileText), 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(outputPath, "executor.go"), []byte(executorFileText), 0777)
	if err != nil {
		return err
	}

	fmt.Println("Success!")
	return nil
}
