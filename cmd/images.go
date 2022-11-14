package cmd

import (
	"errors"
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
	"github.com/spf13/viper"
)

var countValue int
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Download images from your own timeline.",
	Long: `example:
			twitdump images --config config.yaml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config := oauth1.NewConfig(viper.GetViper().GetString("consumerKey"), viper.GetViper().GetString("consumerSecret"))
		token := oauth1.NewToken(viper.GetViper().GetString("accessToken"), viper.GetViper().GetString("accessSecret"))
		httpClient := config.Client(oauth1.NoContext, token)
		// Twitter client
		t := twitter.NewClient(httpClient)

		// We need to keep the user from specifying a value higher than twitter allows.  Twitter will actually honor whatever it's sent, but I think will actually keep the max at 200.
		if countValue > 200 {
			log.Fatal("Error: Value specified for --count cannot be higher than 200.")
		}

		search, resp, err := t.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName: viper.GetViper().GetString("screenName"),
			Count:      countValue,
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
	imagesCmd.PersistentFlags().IntVar(&countValue, "count", 5, "Number of images to search in timeline.  Max=200")
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

	if _, err := os.Stat(fileName); err == nil {
		//fileName exists and we should not download it again.
		fmt.Printf("File %s already exists and will not be downloaded again. \n", fileName)
	} else if errors.Is(err, os.ErrNotExist) {
		//fileName does *not* exist, so proceed with downloading.

		client := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := client.Get(fileURL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Create blank file
		// TODO: Create path based on config
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		// Put content on file
		size, err := io.Copy(file, resp.Body)

		defer file.Close()

		fmt.Printf("Downloaded file %s with size %d \n", fileName, size)
	}

}
