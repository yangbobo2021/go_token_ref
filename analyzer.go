package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

// 判断是否包含go文件
func hasGoFile(path string) bool {
	list, err := os.ReadDir(path)
	if err != nil {
		print("Error: Read dir err", err)
		return false
	}

	for _, d := range list {
		if strings.HasSuffix(d.Name(), ".go") {
			return true
		}
	}

	return false
}

// find all sub dirs
func findGoDirs(path string) []string {
	path_list := []string{}
	path_list = append(path_list, path)

	fmt.Println("dir:", path)
	list, err := os.ReadDir(path)
	if err != nil {
		print("Error: Read dir err", err)
		return path_list
	}

	for _, d := range list {
		if d.IsDir() && d.Name()[0:1] != "." && hasGoFile(path+"/"+d.Name()) {
			path_list = append(path_list, findGoDirs(path+"/"+d.Name())...)
		}
	}

	return path_list
}

func main() {
	// 文件夹转化为文件夹列表
	path_list := findGoDirs(os.Args[1])

	// 解析go项目
	cfg := &packages.Config{Mode: packages.LoadTypes | packages.LoadSyntax}
	pkgs, err := packages.Load(cfg, path_list...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		//os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		//os.Exit(1)
	}

	// 对包进行遍历
	for _, pkg := range pkgs {
		fmt.Println(pkg.ID, pkg.GoFiles)
		fmt.Println(pkg.Types)
		// 符号定义，暂时不考虑
		// for id, obj := range pkg.TypesInfo.Defs {
		// 	fmt.Printf("%s: %q defines %v\n",
		// 		pkg.Fset.Position(id.Pos()), id.Name, obj)
		// }

		// 符号引用处理
		for id, obj := range pkg.TypesInfo.Uses {
			fmt.Println("----<<begin>>----")
			fmt.Printf("%s: %q uses %v\n",
				pkg.Fset.Position(id.Pos()), id.Name, obj)
			fmt.Println("----<<begin2>>----")
			if obj.Pkg() != nil && obj.Pkg().Path() != "xxxx" {
				fmt.Println(obj.Pkg().Path())
			}
			fmt.Println("----<<end>>----")

			// TODO
			// 变量的类型如果定义于第三方库，对于变量的使用，也是一种第三方库相关知识
		}
	}

	os.Exit(0)
}
