package handler

import (
	"context"
	"net/http"

	"github.com/kryptn/http-client/language"
)


func init() {
	registerBlockHandler("post", handleHttpPostBlock)
}

type HttpResult struct {
	resp *http.Response
	err error
}

func (hr *HttpResult) Value() *http.Response {
	return hr.resp
}

func (hr *HttpResult) HasValue() bool {
	return hr.resp != nil
}

func (hr *HttpResult) Error() error {
	return hr.err
}


func extractHeadersFromBlock(block language.Block) http.Header {
	headers := http.Header{}

	for _, sb := range block.SubBlocks {
		if sb.BlockType != "headers" {
			continue
		}

		for _, kv := range sb.KeyValuePairs {
			kv.
		}
	}
}

func httpRequest(method, url string, block language.Block) Result {


	headers := http.Header{}




	req := http.NewRequest(method, url)



}

func handleHttpPostBlock(language.Block) Result {

	

	return HttpResult{
		resp: resp,
		err: nil,
	}


	return HttpResult()
}