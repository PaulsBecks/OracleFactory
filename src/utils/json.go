package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetJSONFromCtx(ctx *gin.Context) (bodyData map[string]interface{}, err error) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	err = json.Unmarshal(data, &bodyData)
	return bodyData, err
}

func GetMapInterfaceFromJson(data []byte) (bodyData map[string]interface{}, err error) {
	err = json.Unmarshal(data, &bodyData)
	return bodyData, err
}
