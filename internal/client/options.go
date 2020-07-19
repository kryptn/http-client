package client

type ClientOption func(*Client)

func WithDocument(documentName string) ClientOption {
	return func(c *Client) {
		c.documentName = documentName
	}
}

func SendRequest(b bool) ClientOption {
	return func(c *Client) {
		c.sendRequest = b
	}
}

func SetDebugDumpAst(b bool) ClientOption {
	return func(c *Client) {
		c.debugDumpAST = b
	}
}

func SetDebugDumpClient(b bool) ClientOption {
	return func(c *Client) {
		c.debugDumpClient = b
	}
}
