package utils

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	tilFiles         map[string][]string
	cnt           = 0
	header        = "### TIL"
	lineSeperator = "\n\n"
	line          = "---"
	introduction  string
)
func init() {
	tilFiles = make(map[string][]string)
}
func SetIntroduction(intro string) {
	introduction = intro
}

func SetHeader(customHeader string) {
	header += customHeader
}

func CreateDocs(tils map[string][]string) {
	file,err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	writeHeader(file, header)
	writeIntro(file, introduction)
	writeContent(file)

}

func writeHeader(file *os.File, header string) {
	_, err := file.Write([]byte(header + lineSeperator + line + lineSeperator))
	if err != nil {
		log.Fatal(err)
	}
}

func writeIntro(file *os.File, intro string) {
	tilCntStr := fmt.Sprintf("*There are %d TIL files!*",cnt)
	_, err := file.Write([]byte(intro + lineSeperator + tilCntStr + lineSeperator))
	if err != nil {
		log.Fatal(err)
	}
}

func writeContent(file *os.File) {
	keys := []string{}
	for key := range tilFiles {
		keys = append(keys,key)
	}

	sort.Strings(keys)

	for _,topic := range keys {
		_, err := file.WriteString(generateTILTopic(topic))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func generateTILTopic(topic string) string{
	topicHeader := fmt.Sprintf("### %s\n\n",topic)
	topicContent := ""
	for _,content := range tilFiles[topic] {
		topicContent += fmt.Sprintf("- [%s](%s)\n",convertToTitle(content),content)
	}
	return topicHeader + topicContent
}

func convertToTitle(path string) string {
	filename := strings.Split(strings.Split(path,"/")[1],".")[0]
	return strings.Replace(filename,"_", " ", -1)
}
func SetTils(tils map[string][]string) {
	tilFiles = tils
	for _,file := range tils {
		cnt += len(file)
	}
}
