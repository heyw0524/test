package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/exec"
)

func init() {
	viper.SetConfigName("config") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(".")    // 设置配置文件和可执行二进制文件在用一个目录
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err)// 读取配置文件失败致命错误
	}
}

func Runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func main() {
	fmt.Println("获取配置文件的string",viper.GetString(`app.name`))
	fmt.Println("获取配置文件的string",viper.GetInt(`app.foo`))
	fmt.Println("获取配置文件的string",viper.GetBool(`app.bar`))
	fmt.Println("获取配置文件的map[string]string",viper.GetStringMapString(`app`))

	v := string(Runcmd("git describe --abbrev=0 --tags",false))
	fmt.Println(v)
}

