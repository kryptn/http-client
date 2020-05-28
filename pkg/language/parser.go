package language

func ParseLanguageFile(filename string) Block {

	pr, err := ParseFile(filename, Debug(false))
	if err != nil {
		panic(err)
	}

	return pr.(Block)
}