package main
import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)
func main() {
	resp, err := http.Get("http://localhost:8081/?id=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	resp2, err := http.PostForm("http://localhost:8081/", url.Values{"name": {"Marry"}, "age": {"18"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()
	io.Copy(os.Stdout, resp2.Body)
}

