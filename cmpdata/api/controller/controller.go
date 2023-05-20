package controller

import (
	"cmpdata/config"
	"cmpdata/consts"
	"cmpdata/logger"
	"cmpdata/utils"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ReqData struct {
	Env   string `uri:"env" binding:"required"`
	Table string `uri:"table" binding:"required"`
}

func ReciveData(c *gin.Context) {
	var data ReqData
	if err := c.ShouldBindUri(&data); err != nil {
		c.Status(http.StatusBadGateway)
		return
	}

	// get env and table
	env := data.Env
	table := data.Table
	logger.InfoLn(fmt.Sprintf("env:[%s], table:[%s]", env, table))

	bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
	bodyStr := fmt.Sprintf("%s|%s|%s", env, table, string(bodyAsByteArray))
	logger.InfoLn(bodyStr)

	// connection to rmq
	envMode := os.Getenv("RUN_MODE")
	connectionString := config.Appconfig.GetString(fmt.Sprintf("%s.server.rmqurl", envMode))
	rmqProducer := utils.RMQProducer{
		Queue:            consts.EXAMPLE_QUEUE,
		ConnectionString: connectionString,
	}

	rmqProducer.PublishMessage("text/plain", []byte(bodyStr))

	c.JSON(http.StatusOK, gin.H{
		"response": "Message received",
	})
}
