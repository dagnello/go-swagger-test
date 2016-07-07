package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

/*GeneratorType Generator concrete type

swagger:discriminator GeneratorType type
*/
type GeneratorType interface {
	runtime.Validatable

	/* type
	 */
	Type() string
	SetType(string)
}

// UnmarshalGeneratorTypeSlice unmarshals polymorphic slices of GeneratorType
func UnmarshalGeneratorTypeSlice(reader io.Reader, consumer runtime.Consumer) ([]GeneratorType, error) {
	var elements [][]byte
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GeneratorType
	for _, element := range elements {
		obj, err := unmarshalGeneratorType(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGeneratorType unmarshals polymorphic GeneratorType
func UnmarshalGeneratorType(reader io.Reader, consumer runtime.Consumer) (GeneratorType, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGeneratorType(data, consumer)
}

func unmarshalGeneratorType(data []byte, consumer runtime.Consumer) (GeneratorType, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "Type1":
		var result Type1
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "Type2":
		var result Type2
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}
