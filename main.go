package main

import (
	"fmt"
	"net/http"
	"os"

	"abtest-server/db"
	"abtest-server/service"
)

func init() {

	// DBPath 实验配置中心
	service.DBPath = os.Getenv("GOPATH") + "/src" + "/abtest-server/db/"
}

func main() {

	// PM设计实验，生成实验配置
	fmt.Println("PM 通过 ABT Server 设计实验，生成实验配置，入库")
	db.DataInit()

	http.HandleFunc("/DescribeConfig", describeConfig)
	http.HandleFunc("/CreateConfig", createConfig)
	http.HandleFunc("/UpdateConfig", updateConfig)

	fmt.Println("AB Test Server Listening ...")
	http.ListenAndServe(":9527", nil)
}

func describeConfig(w http.ResponseWriter, r *http.Request) {
	projectID := r.FormValue("project")
	w.Write(service.DescribeABTestConfig(projectID))
}

func createConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateConfig Not Implement!"))
}

func updateConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateConfig Not Implement!"))
}
