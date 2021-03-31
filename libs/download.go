package libs

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url string, save_path string) {
	fmt.Printf("%s %s", url, save_path)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	dest, err := os.Create("storage/node-v14.16.0.pkg")
	if err != nil {
		panic(err)
	}
	defer dest.Close()
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		panic(err)
	}
}
