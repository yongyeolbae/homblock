package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/baeyongyeol/explorer"
	"github.com/baeyongyeol/rest"
)

func usage() {
	fmt.Printf("발명 코인에 오신것을 환영합니다.\n\n")
	fmt.Printf("다음과 같은 명령어를 입력해 주세요.\n\n")
	fmt.Printf("-port:	서버의 포트를 설정하세요.\n")
	fmt.Printf("-mode:	    html과 rest 중 하나를 선택하세요.\n\n")
	runtime.Goexit()

}
func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	//rest.Start(4000)

	port := flag.Int("port", 4000, "서버의 포트를 설정하세요.")
	mode := flag.String("mode", "rest", "html과 rest 중 하나를 선택하세요.")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
