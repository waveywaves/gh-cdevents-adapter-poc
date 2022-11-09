package main

import (
    "encoding/json"
    "fmt"
    "os"
	"context"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
    gtihubcontext := os.Getenv("GITHUB_CONTEXT")

    var ghevent map[string]interface{}
    json.Unmarshal([]byte(gtihubcontext), &ghevent)

    if ghevent["event_name"] == "pull_request" {
		eventinfo := ghevent["event"].(map[string]interface{})
		if eventinfo["action"] == "opened" {
			event, err := cdevents.NewChangeCreatedEvent()
			if err != nil {
				fmt.Printf("could not create a cdevent, %v", err)
			}
			event.SetSubjectId("actor_id")
			event.SetSource("repositoryUrl")
			ce, err := cdevents.AsCloudEvent(event)
			ctx := cloudevents.ContextWithTarget(context.Background(), "http://cloudevents-player.default.10.108.12.246.sslip.io/")
			ctx = cloudevents.WithEncodingBinary(ctx)

			c, err := cloudevents.NewClientHTTP()
			if err != nil {
				fmt.Printf("failed to create client, %v", err)
			}

			if result := c.Send(ctx, *ce); cloudevents.IsUndelivered(result) {
				fmt.Printf("failed to send, %v", result)
			}
		}
	}
}
