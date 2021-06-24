// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 112.

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/lhbelfanti/goal-book/chapter4/examples/11github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/*
//text output
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
60 issues:
#33416   bserdar encoding/json: This CL adds Decoder.InternKeys
#43716 ggaaooppe encoding/json: increment byte counter when using decode
#45628 pgundlach encoding/xml: add Decoder.InputPos
#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
#11046     kurin encoding/json: Decoder internally buffers full input
#32779       rsc encoding/json: memoize strings during decode
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
#29035    jaswdr proposal: encoding/json: add error var to compare  the
#43401  opennota proposal: encoding/csv: add Reader.InputOffset method
#31701    lr1980 encoding/json: second decode after error impossible
#14750 cyberphon encoding/json: parser ignores the case of member names
#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
#40982   Segflow encoding/json: use different error type for unknown fie
#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
#28923     mvdan encoding/json: speed up the decoding scanner
#40983   Segflow encoding/json: return a different error type for unknow
#16212 josharian encoding/json: do all reflect work before decoding
#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#43513 Alexander encoding/json: add line number to SyntaxError
#45512 colin-sit encoding/json: cannot unmarshal custom interface value
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#29750  jacoelho cmd/vet: stdmethods check gets confused if run on a pac
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
#30301     zelch encoding/xml: option to treat unknown fields as an erro
#26946    deuill encoding/json: clarify what happens when unmarshaling i
#33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
#22752  buyology proposal: encoding/json: add access to the underlying d
*/
