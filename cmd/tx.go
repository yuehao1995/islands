package cmd

import (
	"github.com/jiangjincc/islands/block"

	"github.com/jiangjincc/islands/utils"
	"github.com/spf13/cobra"
)

var (
	// 参数
	from   string
	to     string
	amount string

	address string

	txCmd = &cobra.Command{
		Use:   "tx",
		Short: "交易相关命令",
		Long:  ``,
	}

	txSendCmd = &cobra.Command{
		Use:   "send",
		Short: "发送交易",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			b := block.GetBlockchain()
			b.MineNewBlock(utils.JsonToArray(from), utils.JsonToArray(to), utils.JsonToArray(amount))
		},
	}

	getBalanceCmd = &cobra.Command{
		Use:   "balance",
		Short: "查询月余额",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			b := block.GetBlockchain()
			b.GetBalance(address)

		},
	}
)

func txCmdExecute(rootCmd *cobra.Command) {
	// 解析参数
	txSendCmd.Flags().StringVarP(&from, "from", "f", "", "发送人")
	_ = txSendCmd.MarkFlagRequired("from")

	txSendCmd.Flags().StringVarP(&to, "to", "t", "", "接受人")
	_ = txSendCmd.MarkFlagRequired("to")

	txSendCmd.Flags().StringVarP(&amount, "amount", "a", "", "金额")
	_ = txSendCmd.MarkFlagRequired("amount")

	getBalanceCmd.Flags().StringVarP(&address, "address", "a", "", "地址")
	_ = getBalanceCmd.MarkFlagRequired("address")

	rootCmd.AddCommand(txCmd)
	txCmd.AddCommand(txSendCmd)
	txCmd.AddCommand(getBalanceCmd)
}
