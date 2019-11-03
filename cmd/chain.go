package cmd

import (
	"fmt"
	"log"

	"github.com/jiangjincc/islands/block"

	"github.com/spf13/cobra"
)

var (
	data string

	chainCmd = &cobra.Command{
		Use:   "chain",
		Short: "区块链命令行工具",
		Long:  ``,
	}

	chainAddCmd = &cobra.Command{
		Use:   "add",
		Short: "添加区块",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			block := block.GetBlockchain()
			err := block.AddBlockToBlockChain([]byte(data))
			if err != nil {
				log.Panic(err)
			}
		},
	}

	chainListCmd = &cobra.Command{
		Use:   "list",
		Short: "列出所有区块",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			block := block.GetBlockchain()
			fmt.Println("打印区块")
			block.PrintBlocks()
		},
	}

	chainInitCmd = &cobra.Command{
		Use:   "init",
		Short: "初始化区块链",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			block.CreateBlockchainWithGenesisBlock()
		},
	}
)

func chainCmdExecute(rootCmd *cobra.Command) {
	// 解析参数
	chainAddCmd.Flags().StringVarP(&data, "data", "d", "", "区块数据 (required)")
	_ = chainAddCmd.MarkFlagRequired("data")

	rootCmd.AddCommand(chainCmd)
	chainCmd.AddCommand(chainAddCmd)
	chainCmd.AddCommand(chainListCmd)
	chainCmd.AddCommand(chainInitCmd)
}