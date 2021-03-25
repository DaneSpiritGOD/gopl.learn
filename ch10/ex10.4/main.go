package main

import (
	"encoding/json"
	"flag"
	"log"
	"os/exec"
)

var pkg = flag.String("p", "go/build", "package path")

/* go list -json
   {
   	"Dir": "C:\\Program Files\\Go\\src\\go\\build",
   	"ImportPath": "go/build",
   	"Name": "build",
   	"Doc": "Package build gathers information about Go packages.",
   	"Target": "C:\\Program Files\\Go\\pkg\\windows_amd64\\go\\build.a",
   	"Root": "C:\\Program Files\\Go",
   	"Match": [
   		"go/build"
   	],
   	"Goroot": true,
   	"Standard": true,
   	"GoFiles": [
   		"build.go",
   		"doc.go",
   		"gc.go",
   		"read.go",
   		"syslist.go",
   		"zcgo.go"
   	],
   	"IgnoredGoFiles": [
   		"gccgo.go"
   	],
   	"Imports": [
   		"bufio",
   		"bytes",
   		"errors",
   		"fmt",
   		"go/ast",
   		"go/doc",
   		"go/parser",
   		"go/token",
   		"internal/execabs",
   		"internal/goroot",
   		"internal/goversion",
   		"io",
   		"io/fs",
   		"io/ioutil",
   		"os",
   		"path",
   		"path/filepath",
   		"runtime",
   		"sort",
   		"strconv",
   		"strings",
   		"unicode",
   		"unicode/utf8"
   	],
   	"Deps": [
   		"bufio",
   		"bytes",
   		"context",
   		"errors",
   		"fmt",
   		"go/ast",
   		"go/doc",
   		"go/parser",
   		"go/scanner",
   		"go/token",
   		"internal/bytealg",
   		"internal/cpu",
   		"internal/execabs",
   		"internal/fmtsort",
   		"internal/goroot",
   		"internal/goversion",
   		"internal/lazyregexp",
   		"internal/oserror",
   		"internal/poll",
   		"internal/race",
   		"internal/reflectlite",
   		"internal/syscall/execenv",
   		"internal/syscall/windows",
   		"internal/syscall/windows/registry",
   		"internal/syscall/windows/sysdll",
   		"internal/testlog",
   		"internal/unsafeheader",
   		"io",
   		"io/fs",
   		"io/ioutil",
   		"math",
   		"math/bits",
   		"net/url",
   		"os",
   		"os/exec",
   		"path",
   		"path/filepath",
   		"reflect",
   		"regexp",
   		"regexp/syntax",
   		"runtime",
   		"runtime/internal/atomic",
   		"runtime/internal/math",
   		"runtime/internal/sys",
   		"sort",
   		"strconv",
   		"strings",
   		"sync",
   		"sync/atomic",
   		"syscall",
   		"text/template",
   		"text/template/parse",
   		"time",
   		"unicode",
   		"unicode/utf16",
   		"unicode/utf8",
   		"unsafe"
   	],
   	"TestGoFiles": [
   		"build_test.go",
   		"deps_test.go",
   		"read_test.go",
   		"syslist_test.go"
   	],
   	"TestImports": [
   		"bytes",
   		"flag",
   		"fmt",
   		"go/token",
   		"internal/testenv",
   		"io",
   		"io/fs",
   		"os",
   		"path/filepath",
   		"reflect",
   		"runtime",
   		"sort",
   		"strings",
   		"testing"
   	]
   }
*/

type goList struct {
}

func main() {
	flag.Parse()

	pkgPath := *pkg
	if len(pkgPath) == 0 {
		log.Fatal("you must specify the package path")
	}

	cmd := exec.Command("go", "list", pkgPath)
	stdin, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(stdin).Decode(&person); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
