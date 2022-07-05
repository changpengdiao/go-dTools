/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var exts = []string{"java", "jsp", "css", "js", "xml", "properties", "txt", "lua"}

// retainFileCmd represents the retainFile command
var retainFileCmd = &cobra.Command{
	Use:   "retainFile",
	Short: "遍历一个文件夹，保留后缀为" + strings.Join(exts, "") + ",其他后缀的文件都删除",
	Long:  "遍历一个文件夹，保留后缀为" + strings.Join(exts, "") + ",其他后缀的文件都删除",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("retainFile called")
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			fmt.Println("path参数不能为空")
			return
		}
		//
		walk_path(path)
	},
}

func init() {
	retainFileCmd.Flags().StringP("path", "p", "", "文件夹路径")
	rootCmd.AddCommand(retainFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// retainFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

func walk_path(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, fi := range files {
		filepath := path + string(os.PathSeparator) + fi.Name()
		if fi.IsDir() {
			walk_path(filepath)
		} else {
			lastindex := strings.LastIndex(fi.Name(), ".")
			if lastindex == -1 {
				continue
			}
			ext := string(fi.Name()[lastindex+1:])
			hasExt := false
			for _, v := range exts {
				if v == ext {
					hasExt = true
					break
				}
			}
			if !hasExt {
				//fmt.Println(filepath)
				os.Remove(filepath)
			}
		}
	}
}
