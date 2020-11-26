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
	for _, str := range gendoc() {
		fmt.Println(str)
	}
	fmt.Println("General markdown finished")
}

//gendoc 生成文档的函数，期间会调用各个内容的生成函数
func gendoc() []string {
	info := genInfo()
	servers := genServers()
	paths := genPaths()
	result := append(info, servers...)
	result = append(result, paths...)
	return result
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

//genServers 生成Servers的信息
func genServers() []string {
	//Servers信息
	var result []string
	result = make([]string, 0)
	//录入
	if openapi.Servers != nil {
		servers := openapi.Servers
		if len(servers) != 0 {
			result = append(result, "**服务端列表：**")
			for _, server := range servers {
				if len(server.URL) != 0 {
					if len(server.Description) != 0 {
						result = append(result, "- "+server.URL)
						result = append(result, "    "+server.Description)
					} else {
						result = append(result, "- "+server.URL)
					}
				}
			}
		}
	}
	//servers 中的variable未考虑，不输出
	return result
}

//genPaths 生成api文档中最关键的内容
func genPaths() []string {
	//paths信息
	var result []string
	result = make([]string, 0)
	//录入
	paths := openapi.Paths
	if len(paths) != 0 {
		result = append(result, "## 接口信息")
		for api, pathitem := range paths {
			result = append(result, "### "+api)
			if len(pathitem.Summary) != 0 {
				result = append(result, "**"+pathitem.Summary+"**")
			}
			if len(pathitem.Description) != 0 {
				result = append(result, "> "+pathitem.Description)
			}
			if pathitem.Servers != nil {
				result = append(result, "**服务端列表：**")
				for _, server := range pathitem.Servers {
					if len(server.URL) != 0 {
						result = append(result, "- "+server.URL)

					}
				}
			}
			//以上是基本信息，以下是不同请求方法信息
			if pathitem.Get != nil {
				result = append(result, "#### GET")
				if len(pathitem.Get.Summary) != 0 {
					result = append(result, "**"+pathitem.Get.Summary+"**")
				}
				if len(pathitem.Get.Description) != 0 {
					result = append(result, "> "+pathitem.Get.Description)
				}
				fmt.Println(pathitem.Get)
				fmt.Println("ll")
			}
			if pathitem.Post != nil {
				result = append(result, "#### POST")
				if len(pathitem.Post.Summary) != 0 {
					result = append(result, "**"+pathitem.Post.Summary+"**")
				}
				if len(pathitem.Post.Description) != 0 {
					result = append(result, "> "+pathitem.Post.Description)
				}
			}
			if pathitem.Put != nil {
				result = append(result, "#### PUT")
				if len(pathitem.Put.Summary) != 0 {
					result = append(result, "**"+pathitem.Put.Summary+"**")
				}
				if len(pathitem.Put.Description) != 0 {
					result = append(result, "> "+pathitem.Put.Description)
				}
			}
			if pathitem.Delete != nil {
				result = append(result, "#### DELETE")
				if len(pathitem.Delete.Summary) != 0 {
					result = append(result, "**"+pathitem.Delete.Summary+"**")
				}
				if len(pathitem.Delete.Description) != 0 {
					result = append(result, "> "+pathitem.Delete.Description)
				}
			}
		}

	}

	//	fmt.Println(paths)
	return result
}
