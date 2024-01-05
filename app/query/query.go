package query

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var ServiceQuery queryList

var filterList []string

type queryList struct{}

func (c *queryList) QueryListTopic(srcip []string, srcusername string, srcpassword string) ([]string, error) {
	testAdmin, _ := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(srcip)),
		admin.WithCredentials(primitive.Credentials{
			AccessKey: srcusername,
			SecretKey: srcpassword,
		}),
	)
	result, err := testAdmin.FetchAllTopicList(context.Background())
	if err != nil {
		fmt.Println("FetchAllTopicList error:", err.Error())
	}
	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Query Shutdown admin error: %s", err.Error())
	}
	return result.TopicList, nil
}

func (c *queryList) FilterSystemTopic(srcTopicList []string) ([]string, error) {
	for _, value := range srcTopicList {
		switch value {
		case "rocketmq-cluster-a":
		case "broker-gworker03":
		case "SCHEDULE_TOPIC_XXXX":
		case "RMQ_SYS_TRANS_HALF_TOPIC":
		case "rocketmq-cluster-a_REPLY_TOPIC":
		case "BenchmarkTest":
		case "OFFSET_MOVED_EVENT":
		case "%RETRY%TOOLS_CONSUMER":
		case "SELF_TEST_TOPIC":
		case "TBW102":
		case "broker-gworker01":
		default:
			filterList = append(filterList, value)
		}
	}
	return filterList, nil
}
