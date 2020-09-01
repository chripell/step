package main
import (
	"context"
	"log"
	"fmt"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	chat "google.golang.org/api/chat/v1"
)

// To run locally:
// * Enable the chat API in the GCP console for the project
// * Select the project
//  gcloud config set project google.com:utility-emblem-286918
// * Login with default credentials (follow instructions for the web browser)
//  gcloud auth application-default login

func main() {
  // Setup client to write messages to chat.google.com
	ctx := context.Background()
	client := getOauthClient(ctx)
	service, err := chat.New(client)
	if err != nil {
		log.Fatalf("failed to create chat service: %s", err)
	}
	spaceSrv := chat.NewSpacesService(service)
	token := ""
	for {
		l := spaceSrv.List().Context(ctx)
		if token != "" {
			l = l.PageToken(token)
		}
		resp, err := l.Do()
		if err != nil {
			log.Fatalf("Space listing failed: %v", err)
		}
		token = resp.NextPageToken
		for _, s := range resp.Spaces {
			fmt.Printf("%s %s\n", s.Name, s.DisplayName)
		}
	}
}

func getOauthClient(ctx context.Context) *http.Client {
	creds, err := google.FindDefaultCredentials(
		ctx, 
		"https://www.googleapis.com/auth/chat.bot",
	)
	if err != nil {
		log.Fatal(err)
	}
	return oauth2.NewClient(ctx, creds.TokenSource)
}
