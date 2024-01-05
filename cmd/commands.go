package cmd

import (
	"errors"
	"fmt"
	"os"
	"rocketmqSync/app/query"
	"rocketmqSync/app/sync"

	"github.com/spf13/cobra"
)

var (
	srcipAddress []string
	srcusername  string
	srcpassword  string
	dstipAddress []string
	dstusername  string
	dstpassword  string
	dstbrokerip  string
	enable       bool
)

var RootCmd = &cobra.Command{
	Use:   "Source Rocketmq cluster service",
	Short: "Source Rocketmq cluster NameServer IpAddress IP地址与端口",
	Long:  `源集群地址与端口`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if enable {
			fmt.Println("源集群IP地址:", srcipAddress)
			topiclist, err := query.ServiceQuery.QueryListTopic(srcipAddress, srcusername, srcpassword)
			if err != nil {
				fmt.Println("源集群信息查询失败 !", err.Error())
			}
			// fmt.Println("源集群Topic信息:", topiclist)
			filterTopicList, err := query.ServiceQuery.FilterSystemTopic(topiclist)
			if err != nil {
				fmt.Println("Filter System Topic error", err.Error())
			}
			fmt.Println("Filter User Create Topic Name: ", filterTopicList)
			fmt.Println("目标集群IP地址:", dstipAddress)
			result, err := sync.ServiceDst.SyncTopicInfo(dstipAddress, dstusername, dstpassword, filterTopicList, dstbrokerip)
			if err != nil {
				fmt.Println("Sync Error", err.Error())
			}
			fmt.Println(result)
		}
		return errors.New("no flags find 请填写相关数据参数 !")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringArrayVarP(&srcipAddress, "Src IpAddress", "s", []string{}, "源rocketmq集群nameserverIP与端口地址")
	RootCmd.PersistentFlags().StringVarP(&srcusername, "src UserName", "u", "src cluster name", "源rocketmq集群用户认证用户名称")
	RootCmd.PersistentFlags().StringVarP(&srcpassword, "src Password", "p", "src cluster passwd", "源rocketmq集群用户认证密码")
	RootCmd.PersistentFlags().StringArrayVarP(&dstipAddress, "Dst IpAddress", "d", []string{}, "目标rocketmq集群nameserverIP与端口地址")
	RootCmd.PersistentFlags().StringVarP(&dstusername, "Dst UserName", "i", "Dst cluster name", "目标rocketmq集群用户认证用户名称")
	RootCmd.PersistentFlags().StringVarP(&dstpassword, "Dst Password", "o", "Dst cluster name", "目标rocketmq集群用户认证密码")
	RootCmd.PersistentFlags().StringVarP(&dstbrokerip, "Dst BrokerIP", "r", "Dst BrokerIP", "目标集群rocketmq brokerIP地址与端口")
	RootCmd.PersistentFlags().BoolVarP(&enable, "运行请选择true", "b", false, "运行请选择true,否则为false !")
	// RootCmd.AddCommand(RootCmd)
}
