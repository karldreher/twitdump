package cmd

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func Test_downloadFile(t *testing.T) {
	downloadFile("https://raw.githubusercontent.com/karldreher/twitdump/main/LICENSE")
	if _, err := os.Stat("LICENSE"); err != nil {
		t.Errorf("File was not downloaded")
	}
	file, err := os.Open("LICENSE")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}
	sum := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Println(sum)
	if sum != "4708c809a9e53325e77934783843db2547ee06073119ef51f13fc155c43966fa" {
		t.Error("SHA256 of downloaded file does not match expected")
	}
	// This is done a second time so that we can verify that the file is not downloaded twice.
	downloadFile("https://raw.githubusercontent.com/karldreher/twitdump/main/LICENSE")

}
