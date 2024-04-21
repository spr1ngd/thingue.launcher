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
	"thingue-launcher/client/global"
	"thingue-launcher/client/initialize"
	"thingue-launcher/client/service"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/model"
)

var (
	removeId uint
	modId    uint
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
		_ = survey.AskOne(prompt0, &instance.Name)

		prompt1 := &survey.Input{
			Message: "选择ThingUE启动位置:",
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*")
				return files
			},
		}
		_ = survey.AskOne(prompt1, &instance.ExecPath)

		var selected []string
		prompt2 := &survey.MultiSelect{
			Message:  "实例设置:",
			PageSize: 10,
			Options:  []string{"渲染控制", "多用户同时操作", "故障恢复", "自动启停", "使用WebRTC中继", "根据视口大小调整分辨率", "隐藏控制UI", "无操作关闭连接"},
			Default:  []string{"使用WebRTC中继", "渲染控制"},
		}
		if survey.AskOne(prompt2, &selected) == nil {
			setSelected(instance, selected)
		}

		if instance.AutoControl {
			prompt := &survey.Input{
				Message: "自动启停关闭延迟时间(秒):",
			}
			_ = survey.AskOne(prompt, &instance.StopDelay)
		}

		if instance.PlayerConfig.IdleDisconnect {
			prompt := &survey.Input{
				Message: "无操作关闭连接等待时间(分钟):",
			}
			_ = survey.AskOne(prompt, &instance.PlayerConfig.IdleTimeout)
		}

		launchArguments := ""
		prompt3 := &survey.Editor{
			Message:       "启动参数配置",
			Default:       constants.DEFAULT_THINGUE_LAUNCH_ARGUMENTS,
			AppendDefault: true,
		}
		if survey.AskOne(prompt3, &launchArguments) == nil {
			instance.LaunchArguments = strings.Split(launchArguments, "\n")
		}

		prompt4 := &survey.Editor{
			Message:       "元数据配置",
			Default:       instance.Metadata,
			AppendDefault: true,
		}
		_ = survey.AskOne(prompt4, &instance.Metadata)

		prompt5 := &survey.Editor{
			Message:       "Pak资源配置",
			Default:       instance.PaksConfig,
			AppendDefault: true,
		}
		_ = survey.AskOne(prompt5, &instance.PaksConfig)

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

var instanceModCmd = &cobra.Command{
	Use: `mod`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client.Init()
		instance := service.InstanceManager.GetById(modId)
		if instance == nil {
			fmt.Println("实例不存在")
			return nil
		}

		prompt0 := &survey.Input{
			Message: "实例名称:",
			Default: instance.Name,
		}
		_ = survey.AskOne(prompt0, &instance.Name)

		prompt1 := &survey.Input{
			Message: "选择ThingUE启动位置:",
			Default: instance.ExecPath,
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*")
				return files
			},
		}
		_ = survey.AskOne(prompt1, &instance.ExecPath)

		var selected []string
		prompt2 := &survey.MultiSelect{
			Message:  "实例设置:",
			PageSize: 10,
			Options:  []string{"渲染控制", "多用户同时操作", "故障恢复", "自动启停", "使用WebRTC中继", "根据视口大小调整分辨率", "隐藏控制UI", "无操作关闭连接"},
			Default:  reverseSelected(instance),
		}
		if survey.AskOne(prompt2, &selected) == nil {
			setSelected(instance, selected)
		}

		if instance.AutoControl {
			prompt := &survey.Input{
				Message: "自动启停关闭延迟时间(秒):",
				Default: strconv.Itoa(instance.StopDelay),
			}
			_ = survey.AskOne(prompt, &instance.StopDelay)
		}

		if instance.PlayerConfig.IdleDisconnect {
			prompt := &survey.Input{
				Message: "无操作关闭连接等待时间(分钟):",
				Default: strconv.Itoa(int(instance.PlayerConfig.IdleTimeout)),
			}
			_ = survey.AskOne(prompt, &instance.PlayerConfig.IdleTimeout)
		}

		launchArguments := ""
		prompt3 := &survey.Editor{
			Message:       "启动参数配置",
			Default:       strings.Join(instance.LaunchArguments, "\n"),
			AppendDefault: true,
		}
		if survey.AskOne(prompt3, &launchArguments) == nil {
			instance.LaunchArguments = strings.Split(launchArguments, "\n")
		}

		prompt4 := &survey.Editor{
			Message:       "元数据配置",
			Default:       instance.Metadata,
			AppendDefault: true,
		}
		_ = survey.AskOne(prompt4, &instance.Metadata)

		prompt5 := &survey.Editor{
			Message:       "Pak资源配置",
			Default:       instance.PaksConfig,
			AppendDefault: true,
		}
		_ = survey.AskOne(prompt5, &instance.PaksConfig)

		global.APP_DB.Save(instance)
		return nil
	},
}

func init() {
	instanceCmd.AddCommand(instanceListCmd)
	instanceCmd.AddCommand(instanceAddCmd)
	instanceRemoveCmd.Flags().UintVarP(&removeId, "id", "i", 0, "要删除的实例ID")
	instanceCmd.AddCommand(instanceRemoveCmd)
	instanceModCmd.Flags().UintVarP(&modId, "id", "i", 0, "要修改的实例ID")
	instanceCmd.AddCommand(instanceModCmd)
	rootCmd.AddCommand(instanceCmd)
}

func setSelected(instance *model.ClientInstance, selected []string) {
	instance.EnableRenderControl = false
	instance.EnableMultiuserControl = false
	instance.FaultRecover = false
	instance.AutoControl = false
	instance.EnableRelay = false
	instance.PlayerConfig.MatchViewportRes = false
	instance.PlayerConfig.HideUI = false
	instance.PlayerConfig.IdleDisconnect = false
	for _, s := range selected {
		if "渲染控制" == s {
			instance.EnableRenderControl = true
		} else if "多用户同时操作" == s {
			instance.EnableMultiuserControl = true
		} else if "故障恢复" == s {
			instance.FaultRecover = true
		} else if "自动启停" == s {
			instance.AutoControl = true
		} else if "使用WebRTC中继" == s {
			instance.EnableRelay = true
		} else if "根据视口大小调整分辨率" == s {
			instance.PlayerConfig.MatchViewportRes = true
		} else if "隐藏控制UI" == s {
			instance.PlayerConfig.HideUI = true
		} else if "无操作关闭连接" == s {
			instance.PlayerConfig.IdleDisconnect = true
		}
	}
}

func reverseSelected(instance *model.ClientInstance) []string {
	var selected []string
	if instance.EnableRenderControl {
		selected = append(selected, "渲染控制")
	}
	if instance.EnableMultiuserControl {
		selected = append(selected, "多用户同时操作")
	}
	if instance.FaultRecover {
		selected = append(selected, "故障恢复")
	}
	if instance.AutoControl {
		selected = append(selected, "自动启停")
	}
	if instance.EnableRelay {
		selected = append(selected, "使用WebRTC中继")
	}
	if instance.PlayerConfig.MatchViewportRes {
		selected = append(selected, "根据视口大小调整分辨率")
	}
	if instance.PlayerConfig.HideUI {
		selected = append(selected, "隐藏控制UI")
	}
	if instance.PlayerConfig.IdleDisconnect {
		selected = append(selected, "无操作关闭连接")
	}
	fmt.Println(selected)
	return selected
}
