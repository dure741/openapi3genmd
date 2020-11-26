package main

import (
	"flag"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

var (
	inputFilePath  string
	outputFilePath string
	follow         bool
	openapi        *openapi3.Swagger
	openapiErr     error
)

func init() {
	flag.StringVar(&inputFilePath, "input", "openapi.yaml", "OpenAPI specification path")
	flag.StringVar(&outputFilePath, "output", "openapi.md", "Output markdown path")
	flag.BoolVar(&follow, "follow", false, "If the file already exists, write it consecutively at the end, use \"true\" or \"false\"")
	flag.Parse()
}

func main() {
	openapi, openapiErr = openapi3.NewSwaggerLoader().LoadSwaggerFromFile(inputFilePath)
	if openapiErr != nil {
		fmt.Println(openapiErr)
		return
	}
	fmt.Println("Openapi3gendoc start...")
	gendoc()
	fmt.Println("General markdown finished")
}

//gendoc 生成文档的函数，期间会调用各个内容的生成函数
func gendoc() {
	genInfo()
}

//genInfo 在markdown中生成info信息
func genInfo() []string {
	//Info信息
	var result []string
	result = make([]string, 0)
	//录入
	info := openapi.Info
	if len(info.Title) != 0 {
		result = append(result, "# "+info.Title)
	}
	if len(info.Description) != 0 {
		result = append(result, info.Description)
	}
	if len(info.TermsOfService) != 0 {
		result = append(result, "["+info.TermsOfService+"}("+info.TermsOfService+")")
	}
	if info.Contact != nil {
		contact := info.Contact
		if len(contact.Name) != 0 {
			if len(contact.URL) != 0 {
				result = append(result, "[**"+contact.Name+"**]("+contact.URL+")")
			} else {
				result = append(result, "**"+contact.Name+"**")
			}
		}
		if len(contact.Email) != 0 {
			result = append(result, "["+contact.Email+"]("+contact.Email+")")
		}
	}
	if info.License != nil {
		licence := info.License
		if len(licence.URL) != 0 {
			result = append(result, "["+licence.Name+"]("+licence.URL+")")
		} else {
			result = append(result, licence.Name)
		}
	}
	if len(info.Version) != 0 {
		result = append(result, info.Version)
	}
	return result
}

// func content(content interface{}) {
// 	if content!=nil

// }
