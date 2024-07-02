package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

func main() {

	// 设置配置位置
	viper.AddConfigPath("./viper-example/viper")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		// 读取配置失败
		panic(err)
	}

	// 读取配置
	cfg := struct {
		Web struct {
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			IdleTimeout     time.Duration
			ShutdownTimeout time.Duration
			APIHost         string
			DebugHost       string
		}
	}{} // 定义配置结构体

	// 设置配置默认值
	viper.SetDefault("Web.ReadTimeout", "5s")
	viper.SetDefault("Web.WriteTimeout", "10s")
	viper.SetDefault("Web.IdleTimeout", "120s")
	viper.SetDefault("Web.ShutdownTimeout", "20s")
	viper.SetDefault("Web.APIHost", "0.0.0.0:3000")
	viper.SetDefault("Web.DebugHost", "0.0.0.0:4000")

	// rootCmd 是一个命令行工具
	rootCmd := &cobra.Command{
		Use: "viper-example",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	// 添加命令行参数
	rootCmd.PersistentFlags().StringVar(&cfg.Web.APIHost, "api-host", "0.0.0.0:3000", "API host")
	// 能够根据命令行改变配置
	viper.BindPFlags(rootCmd.Flags())

	err := viper.Unmarshal(&cfg)
	if err != nil {
		// 解析配置失败
		panic(err)
	}

	// 打印配置cfg
	println(cfg.Web.APIHost)
	println(cfg.Web.DebugHost)
	println(cfg.Web.ReadTimeout)
	println(cfg.Web.WriteTimeout)
	println(cfg.Web.IdleTimeout)
	println(cfg.Web.ShutdownTimeout)

	err = rootCmd.Execute()
	if err != nil {
		// 执行命令失败
		panic(err)
	}
}
