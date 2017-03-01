package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// API 数据
type API interface{}

// Projects 所有项目
type Projects map[string]API

// GetProjects 获取已创建 API 列表
// @Accept  json
// @Produce json
// @Success 200 {object} Projects
// @Router /admin/projects [get]
func GetProjects(c *gin.Context) {
	var apis []APIS
	db := GetDB()
	db.Find(&apis)

	projects := make(Projects)

	for _, api := range apis {
		var a API
		json.Unmarshal([]byte(api.Value), &a)
		projects[api.Method+" "+api.Path] = a
	}

	c.JSON(200, &projects)
}

func CreateORUpdateProject(c *gin.Context) {
	var payload struct {
		Url  string
		Data API
	}
	c.BindJSON(&payload)

	method, path := getMethodAndPath(payload.Url)
	api, _ := json.Marshal(payload.Data)
	apiModel := APIS{
		Method: method,
		Path:   path,
		Value:  string(api)}

	db := GetDB()
	var exitApi []APIS
	db.Where("method = ? AND path = ?", method, path).Find(&exitApi)
	if len(exitApi) > 0 {
		db.Model(&exitApi[0]).Update("value", apiModel.Value)
	} else {
		db.Create(&apiModel)
	}
	c.JSON(200, gin.H{
		"result": "success",
	})
}

func DeleteProject(c *gin.Context) {
	var payload struct {
		Url string
	}
	c.BindJSON(&payload)
	method, path := getMethodAndPath(payload.Url)
	db := GetDB()
	db.Where("method = ? AND path = ?", method, path).Delete(&APIS{})
	c.JSON(200, gin.H{
		"result": "delete success",
	})
}

// GetAPI 获取已创建 API 列表
// @Accept  json
// @Produce json
// @Success 200 {object} Api
// @Router /admin/projects [get]
func GetAPI(method string, path string) API {
	db := GetDB()
	var apis []APIS
	db.Where("path = ? AND method in (?)", path, []string{method, "any"}).Find(&apis)
	var api API
	switch len(apis) {
	case 2:
		for _, a := range apis {
			if a.Method == method {
				json.Unmarshal([]byte(a.Value), &api)
				break
			}
		}
	case 1:
		json.Unmarshal([]byte(apis[0].Value), &api)
	case 0:
		return nil
	}
	return api
}
