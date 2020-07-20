package runner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/kryptn/http-client/language"
)

type BlockHandler func([]language.KeyValuePair, []language.Block) (string, error)

type FunctionHandler func(language.FunctionInvocation) (string, error)

func handleEnv(fn language.FunctionInvocation) (string, error) {
	var defaultValue = ""
	if len(fn.Arguments) == 2 {
		defaultValue = coalesceArguments(fn.Arguments).StrValue
	}
	value, ok := os.LookupEnv(fn.Arguments[0].StrValue)
	if !ok {
		return defaultValue, nil
	}
	return value, nil
}

func handleFunction(fns []language.FunctionInvocation) (string, error) {
	var result = ""
	var err error
	for _, fn := range fns {
		handler := getHandlerForFunction(fn)
		result, err = handler(fn)

	}
	if err != nil {
		return "", err
	}
	return result, nil

}

func coalesceValues(values []language.Value) language.Value {
	var result = ""
	for _, value := range values {
		if value.FuncValue != nil {
			fnResult, err := handleFunction(value.FuncValue)
			if err != nil {
				return language.Value{StrValue: ""}
			}
			result = result + fnResult
		} else {
			result = result + value.StrValue
		}
	}
	return language.Value{StrValue: result}

}

func coalesceArguments(values []language.Argument) language.Value {
	var result = ""
	for _, value := range values {
		if value.FuncValue != nil {
			fnResult, err := handleFunction(value.FuncValue)
			if err != nil {
				return language.Value{StrValue: ""}
			}
			result = result + fnResult
		} else {
			result = result + value.StrValue
		}
	}
	return language.Value{StrValue: result}
}

func handleHttpPostBlock(kv []language.KeyValuePair, subblocks []language.Block) (string, error) {

	var kvs = map[string]string{}
	for _, kv := range kv {
		kvs[string(kv.Key)] = kv.Values[0].StrValue
	}

	url, ok := kvs["url"]
	if !ok {
		return "", errors.New("no url")
	}

	payload, err := handleJsonBlock(subblocks[0])
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	spew.Dump(resp)

	fmt.Println(url)
	fmt.Println(resp.Status)

	

	var content []byte
	_, err = resp.Body.Read(content)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func handleJsonBlock(block language.Block) ([]byte, error) {

	data := map[string]string{}

	for _, kvs := range block.KeyValuePairs {
		data[string(kvs.Key)] = coalesceValues(kvs.Values).StrValue
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return payload, nil

}

func getHandlerForFunction(fn language.FunctionInvocation) FunctionHandler {
	switch fn.FuncName {
	case "env":
		return handleEnv
	default:
		return handleEnv
	}
}

func getHandlerForBlockType(block language.Block) BlockHandler {
	switch block.BlockType {
	case "post":
		return handleHttpPostBlock
	default:
		return handleHttpPostBlock
	}
}

func RunBlock(block language.Block) (string, error) {

	handler := getHandlerForBlockType(block)

	result, err := handler(block.KeyValuePairs, block.SubBlocks)
	if err != nil {
		return "", err
	}

	return result, nil

}
