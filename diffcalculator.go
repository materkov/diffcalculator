package diffcalculator

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var snsClient *sns.SNS

func init() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("[ERROR] error creating AWS session: %s", err)
	}
	snsClient = sns.New(sess)
}

// Calculate calculate diff from previous version, and sends new items to SNS
func Calculate(sourceID string, items []Item) error {
	oldItems, err := StdStore.Get(sourceID)
	if err != nil {
		return fmt.Errorf("error getting from storage: %s", err)
	}

	err = StdStore.Save(sourceID, items)
	if err != nil {
		return fmt.Errorf("error saving to storage: %s", err)
	}

	if len(oldItems) == 0 {
		return nil // First run (no old items)
	}

	itemsAdded := make([]Item, 0)

	for _, item := range items {
		found := false
		for _, oldItem := range oldItems {
			if item.ID == oldItem.ID {
				found = true
				break
			}
		}

		if !found {
			itemsAdded = append(itemsAdded, item)
		}
	}

	if len(itemsAdded) > 0 {
		for _, item := range itemsAdded {
			itemMarshalled, err := json.Marshal(item.Data)
			if err != nil {
				return fmt.Errorf("error marshaling item: %s", err)
			}

			input := sns.PublishInput{}
			input.SetTopicArn("arn:aws:sns:eu-central-1:563473344515:message")
			input.SetMessage(string(itemMarshalled))
			if _, err = snsClient.Publish(&input); err != nil {
				return fmt.Errorf("error sending message: %s", err)
			}
		}
	}

	return nil
}
