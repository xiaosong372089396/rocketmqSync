package sync

import (
	"context"
	"errors"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var ServiceDst dstSyncTopic

type dstSyncTopic struct{}

func (d *dstSyncTopic) SyncTopicInfo(dstip []string, dstusername string, dstpassword string, topicList []string, brokerAddr string) (string, error) {
	syncAdmin, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(dstip)),
		admin.WithCredentials(primitive.Credentials{
			AccessKey: dstusername,
			SecretKey: dstpassword,
		}),
	)
	if err != nil {
		fmt.Println("sync 集群认证错误, 请检查用户名密码是否正确 !", err.Error())
	}
	for _, topicInfo := range topicList {
		err = syncAdmin.CreateTopic(
			context.Background(),
			admin.WithTopicCreate(topicInfo),
			admin.WithBrokerAddrCreate(brokerAddr),
			admin.WithReadQueueNums(16),
			admin.WithWriteQueueNums(16),
			admin.WithPerm(6),
		)
		if err != nil {
			return "Sync Topic Error:", errors.New(err.Error())
		}
	}
	err = syncAdmin.Close()
	if err != nil {
		fmt.Printf("Sync Shutdown admin error: %s", err.Error())
	}
	return "Sync Topic Successfuly", nil
}
