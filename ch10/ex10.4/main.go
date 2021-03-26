package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var pkgF = flag.String("p", "go/build", "package path")

/* go list -json
   {
   	"Dir": "C:\\Program Files\\Go\\src\\go\\build",
   	"ImportPath": "go/build",
   	"Name": "build",
   	"Deps": [
   		"go/ast",
   		"go/doc",
   		"go/parser",
   		"go/scanner",
   		"go/token",
   		"io",
   		"io/fs",
   		"io/ioutil",
   	],
   }
*/

type goList struct {
	ImportPath string
	Deps       []string
}

func main() {
	flag.Parse()

	pkgRoot := *pkgF
	if len(pkgRoot) == 0 {
		log.Fatal("you must specify the package path")
	}

	root := getGoList(pkgRoot)

	m := make(map[string]bool)
	for _, dep := range root.Deps {
		depRoot := getGoList(dep)
		for _, depDep := range depRoot.Deps {
			m[depDep] = true
		}
	}

	fmt.Println("\nDependencies of sub dependency are:")
	for k := range m {
		fmt.Println(k)
	}
}

func getGoList(pkg string) *goList {
	log.Printf("running command: `go list %s`", pkg)

	cmd := exec.Command("go", "list", "-json", pkg)
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("get out pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd start: %v", err)
	}

	var l goList
	if err := json.NewDecoder(out).Decode(&l); err != nil {
		log.Fatalf("json decode: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("cmd wait: %v", err)
	}

	return &l
}
