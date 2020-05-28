package language

func ParseLanguageFile(filename string) Block {

	pr, err := ParseFile(filename, Debug(true))
	if err != nil {
		panic(err)
	}

	return pr.(Block)
}