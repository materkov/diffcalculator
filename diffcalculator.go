package diffcalculator

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const exchangeName = "sender"

var snsClient *sns.SNS

func init() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("[ERROR] error creating AWS session: %s", err)
	}
	snsClient = sns.New(sess)
}

// Calculate calculate diff from previous version, and sends new posts to SNS
func Calculate(sourceID string, posts []Post) error {
	oldPosts, err := StdStore.Get(sourceID)
	if err != nil {
		return fmt.Errorf("error getting from storage: %s", err)
	}

	err = StdStore.Save(sourceID, posts)
	if err != nil {
		return fmt.Errorf("error saving to storage: %s", err)
	}

	if len(oldPosts) == 0 {
		return nil // First run (no old posts)
	}

	postsAdded := make([]Post, 0)

	for _, post := range posts {
		found := false
		for _, oldPost := range oldPosts {
			if post.ExternalID == oldPost.ExternalID {
				found = true
				break
			}
		}

		if !found {
			postsAdded = append(postsAdded, post)
		}
	}

	if len(postsAdded) > 0 {
		for _, post := range postsAdded {
			postMarshalled, err := json.Marshal(post)
			if err != nil {
				return fmt.Errorf("error marshaling post: %s", err)
			}

			input := sns.PublishInput{}
			input.SetTopicArn("arn:aws:sns:eu-central-1:563473344515:message")
			input.SetMessage(string(postMarshalled))
			if _, err = snsClient.Publish(&input); err != nil {
				return fmt.Errorf("error sending message: %s", err)
			}
		}
	}

	return nil
}
