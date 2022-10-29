package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Download images from your own timeline.",
	Long: `example:  
			twitdump images --config config.yaml		
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Wire this up with the config - these do nothing now
		config := oauth1.NewConfig("consumerKey", "consumerSecret")
		token := oauth1.NewToken("accessToken", "accessSecret")
		httpClient := config.Client(oauth1.NoContext, token)
		// Twitter client
		t := twitter.NewClient(httpClient)
		// TODO: paging?  It's supposed to be 100, but in functional testing this only got about 24 posts.
		search, resp, err := t.Timelines.UserTimeline(&twitter.UserTimelineParams{
			// TODO: Wire this up in configuration
			ScreenName: "MyCoolTwitterNameWithoutAmpersand",
		})

		if (err == nil) && (resp != nil) {
			for _, element := range search {
				if element.Entities.Media != nil {
					// TODO Param large by default?
					downloadUrl := element.Entities.Media[0].MediaURLHttps + "?name=large"
					downloadFile(downloadUrl)
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}

func downloadFile(fileURL string) {
	var fileName string

	f, err := url.Parse(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	path := f.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	// Create file
	// TODO: Create path based on config
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded file %s with size %d", fileName, size)

}
