package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	//"strings"
	"reflect"
)

func main() {
	var i map[string]interface{}
	fpath := "/Users/zjgbz/go-workspace/src/github.com/tiangan10/TomlDoc/samples/default1.toml"
	bs, err := ioutil.ReadFile(fpath)
	fmt.Println(string(bs))
	if err != nil {
		panic(err)
	}
	
	if _, err = toml.Decode(string(bs), &i); err != nil {
		panic(err)
	}

	fmt.Printf("the decoded:\n%v", i)
	test(i["fruit"])
	
	
}
func test(t interface{}) {
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(t)
        fmt.Println("here\n", s.Index(1))
    }
}

// const annotationPrefix = "##user"
// const annotationType = "@type"
// const annotationKey = "@key"
// const annotationFormat = "@format"
// const annotationFormat = "@required"
// func generateOverrideInstructions (content string) {
// 	lines := strings.Split(content, "\n")
// 	for i, l := range lines {
// 		trimmed := strings.TrimSpace(l)
// 		if strings.HasPrefix(trimmed, annotatinoPrefix) {
// 			indEle, ele = findClosesElement(i+1, lines)
// 			indPrt, parent = findClosesElement(i-1, lines)
// 			getInstruction(parent, ele)
// 		} 
// 	}
// }
// func findClosestHeader(ind int, lines []string) {
// 	for i := ind; i >= 0; i-- {
// 		l := strings.TrimSpace(lines[i])
// 		if  l.HasPrefix('[[') {
// 			arrayInd = indArrayIndex(i, lines)
// 		}
// 	}
// } 


// type overrideInstructin struct {
// 	Rep string
// 	T string
// 	Dest string
// 	Format string
// 	Required bool
// 	ArrayInd 
// }