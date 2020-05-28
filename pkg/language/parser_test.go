package language

import (
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"os"
	"testing"
)



func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSimpleParse(t *testing.T) {
	input := `blockType {
    key: value
	key2: more value
	block2 {
		haha: nested key values
		yes: yes haha
	}
}
`

	file, err := ioutil.TempFile("/tmp", "hc_parser_test")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(file.Name())

	_, err = file.Write([]byte(input))
	if err != nil  {
		log.Fatal(err)
	}

	expectedBlockType := "blockType"
	//expectedWords := []string{"firstWord", "secondWord", "thirdWord", "fourthWord", "fifthWorld"}

	// actually run test
	result := ParseLanguageFile(file.Name())

	if result.BlockType != expectedBlockType {
		t.Fail()
	}
	//
	//if !Equal(expectedWords, result.Words) {
	//	t.Fail()
	//}

	print("parsed result: \n\n")
	spew.Dump(result)


}