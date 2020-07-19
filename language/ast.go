package language

type Document struct {
	Block Block
}

type Block struct {
	BlockType     string
	KeyValuePairs []KeyValuePair
	SubBlocks     []Block
}

type FunctionInvocation struct {
	FuncName  string
	Arguments []Argument
}

type Argument Value

type KeyValuePair struct {
	Key    Key
	Values []Value
}

type Key string

type Value struct {
	StrValue  string
	FuncValue []FunctionInvocation
}

type BlockEnum int

const (
	BlockHTTP BlockEnum = iota
	BlockPOST
	BlockDELETE
	BlockGET
	BlockPUT
	BlockPATCH
	BlockGraphQL
)

type SubBlockEnum int

const (
	SubBlockHeaders SubBlockEnum = iota
	SubBlockJson
	SubBlockData
)

type FunctionEnum int

const (
	FunctionReq FunctionEnum = iota
	FuctionResp
	FunctionJQ
	FunctionShell
)

// const commonHttp = []SubBlockEnum{SubBlockHeaders, SubBlockJson, SubBlockData}

// func SubBlockTypesForBlockType(bt BlockEnum) ([]SubBlockEnum, error) {

// 	switch bt {
// 	case BlockHTTP:
// 		commonHttp
// 	case BlockPOST:
// 		commonHttp
// 	case BlockDELETE:
// 		commonHttp
// 	case BlockGET:
// 		commonHttp
// 	case BlockPUT:
// 		commonHttp
// 	case BlockPATCH:
// 		commonHttp
// 	case BlockGraphQL:
// 		commonHttp
// 	default:
// 		return nil, errors.New("oh no")
// 	}

// 	return nil, nil
// }

