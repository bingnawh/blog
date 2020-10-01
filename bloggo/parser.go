package main

import (
	"io/ioutil"
    "strings"
    "bufio"
)


const numMetadataLines = 4

func split(input []byte) (map[string]string, []byte){
	frontMatter := make(map[string]string)


    array := strings.Split(string(input), "---")
    r1 := strings.NewReader(array[0])
    r2 := strings.NewReader(array[1])

    scanner := bufio.NewScanner(r1)
    for scanner.Scan() {
        metadataLine := strings.Split(scanner.Text(), ":")
        frontMatter[metadataLine[0]] = metadataLine[1]
   }

    content, _ := ioutil.ReadAll(r2)
	return frontMatter, content
}

func parse(fileName string) (map[string]string, []byte) {
	b, _ := ioutil.ReadFile(fileName)

    frontMatter, content:= split(b)
	return frontMatter, content

}
