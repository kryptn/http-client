package language

type Document struct {
	Block Block
}

type Block struct {
	BlockType string
	KeyValuePairs []KeyValuePair
	SubBlocks []Block
}


type FunctionInvocation struct {
	FuncName string
	Arguments []Argument
}

type Argument Value

type KeyValuePair struct {
	Key Key
	Values []Value
}

type Key string

type Value struct {
	StrValue string
	FuncValue []FunctionInvocation
}
