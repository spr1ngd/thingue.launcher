package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"thingue-launcher/client"
	"thingue-launcher/client/initialize"
	"thingue-launcher/client/service"
	"thingue-launcher/common/model"
)

var (
	removeId uint
)

var instanceCmd = &cobra.Command{
	Use:   `instance`,
	Short: "管理实例",
	Long:  "Commands for manage ue instances",
	RunE: func(cmd *cobra.Command, args []string) error {
		//初始化Gorm
		fmt.Println("初始化Gorm112")
		initialize.InitGorm()
		return nil
	},
}

var instanceListCmd = &cobra.Command{
	Use: `list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client.Init()
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "实例名称", "启动位置"})
		list := service.InstanceManager.List()
		for _, instance := range list {
			table.Append([]string{strconv.Itoa(int(instance.CID)), instance.Name, instance.ExecPath})
		}
		table.Render()
		return nil
	},
}

var instanceAddCmd = &cobra.Command{
	Use: `add`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client.Init()
		var instance = &model.ClientInstance{}

		prompt0 := &survey.Input{
			Message: "实例名称:",
		}
		survey.AskOne(prompt0, &instance.Name)

		prompt1 := &survey.Input{
			Message: "选择ThingUE启动位置:",
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*")
				return files
			},
		}
		survey.AskOne(prompt1, &instance.ExecPath)

		selected := []string{}
		prompt2 := &survey.MultiSelect{
			Message:  "实例设置:",
			PageSize: 10,
			Options:  []string{"渲染控制", "多用户同时操作", "故障恢复", "自动启停", "使用WebRTC中继", "根据视口大小调整分辨率", "隐藏控制UI", "无操作关闭连接"},
			Default:  []string{"使用WebRTC中继", "渲染控制"},
		}
		survey.AskOne(prompt2, &selected)
		for _, s := range selected {
			if "渲染控制" == s {
				instance.EnableRenderControl = true
			} else if "多用户同时操作" == s {
				instance.EnableMultiuserControl = true
			} else if "故障恢复" == s {
				instance.FaultRecover = true
			} else if "自动启停" == s {
				instance.AutoControl = true
				prompt := &survey.Input{
					Message: "自动启停->关闭延迟时间(秒):",
				}
				survey.AskOne(prompt, &instance.StopDelay)
			} else if "使用WebRTC中继" == s {
				instance.EnableRelay = true
			} else if "根据视口大小调整分辨率" == s {
				instance.PlayerConfig.MatchViewportRes = true
			} else if "隐藏控制UI" == s {
				instance.PlayerConfig.HideUI = true
			} else if "无操作关闭连接" == s {
				instance.PlayerConfig.IdleDisconnect = true
				prompt := &survey.Input{
					Message: "无操作关闭连接->等待时间(分钟):",
				}
				survey.AskOne(prompt, &instance.PlayerConfig.IdleTimeout)
			}
		}

		launchArguments := `-AudioMixer
-RenderOffScreen
-ForceRes
-ResX=1920
-ResY=1080`
		prompt3 := &survey.Editor{
			Message:       "启动参数配置",
			Default:       launchArguments,
			AppendDefault: true,
			Help:          "Enter进入编辑器修改，Ctrl+C使用当前值",
		}
		survey.AskOne(prompt3, &launchArguments)
		instance.LaunchArguments = strings.Split(launchArguments, "\n")

		metadata := ""
		prompt4 := &survey.Editor{
			Message:       "元数据配置",
			Default:       metadata,
			AppendDefault: true,
			Help:          "Enter进入编辑器修改，Ctrl+C使用当前值",
		}
		survey.AskOne(prompt4, &metadata)
		instance.Metadata = metadata

		paksConfig := ""
		prompt5 := &survey.Editor{
			Message:       "Pak资源配置",
			Default:       paksConfig,
			AppendDefault: true,
			Help:          "Enter进入编辑器修改，Ctrl+C使用当前值",
		}
		survey.AskOne(prompt5, &paksConfig)
		instance.PaksConfig = paksConfig

		service.InstanceManager.Create(instance)
		return nil
	},
}

var instanceRemoveCmd = &cobra.Command{
	Use: `remove`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client.Init()
		service.InstanceManager.Delete(removeId)
		return nil
	},
}

func init() {
	instanceCmd.AddCommand(instanceListCmd)
	instanceCmd.AddCommand(instanceAddCmd)
	instanceRemoveCmd.Flags().UintVar(&removeId, "id", 0, "要删除的实例ID")
	instanceCmd.AddCommand(instanceRemoveCmd)
	rootCmd.AddCommand(instanceCmd)
}
