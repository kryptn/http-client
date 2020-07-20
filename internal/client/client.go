package client

import (
	"fmt"
	"os"

	"github.com/kryptn/http-client/internal/runner"

	"github.com/davecgh/go-spew/spew"
	"github.com/kryptn/http-client/language"
)

type Client struct {
	documentName     string
	documentFileName string

	sendRequest bool

	debugDumpAST    bool
	debugDumpClient bool
}

func NewClient(opts ...ClientOption) *Client {

	const (
		defaultDocumentName     = ""
		defaultDocumentFileName = ""
		defaultSendRequest      = false

		defaultDebugDumpAST    = false
		defaultDebugDumpClient = false
	)

	c := &Client{
		documentName:     defaultDocumentName,
		documentFileName: defaultDocumentFileName,
		sendRequest:      defaultSendRequest,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func identifyFile(name string) (string, error) {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		return name, nil
	}

	nameWithExt := fmt.Sprintf("%s.hc", name)
	_, err = os.Stat(nameWithExt)
	if !os.IsNotExist(err) {
		return nameWithExt, nil
	}

	return "", err
}

func (c *Client) Execute() (string, error) {

	if c.documentFileName == "" {
		name, err := identifyFile(c.documentName)
		if err != nil {
			return "", nil
		}
		c.documentFileName = name
	}

	result := language.ParseLanguageFile(c.documentFileName)

	if c.debugDumpAST {
		fmt.Println("dumping ast")
		spew.Dump(result)
	}

	if c.debugDumpClient {
		fmt.Println("dumping client")
		spew.Dump(c)
	}

	var response = ""

	if c.sendRequest {
		var err error
		response, err = runner.RunBlock(result)
		if err != nil {
			return response, err
		}

		fmt.Print("sent? ")
	}

	return response, nil
}
