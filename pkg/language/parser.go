package language

import "github.com/davecgh/go-spew/spew"

func dumps(origin string, values ...interface{}) {
	// helps debug
	for _, value := range values {
		print("dumping from ", origin, ":\n\n")
		spew.Dump(value)
		print("\n\n")
	}
}

func ParseLanguageFile(filename string) Block {

	pr, err := ParseFile(filename, Debug(false))
	if err != nil {
		panic(err)
	}

	return pr.(Block)
}
