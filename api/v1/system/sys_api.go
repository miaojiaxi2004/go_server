package system

import (
	"github.com/gin-gonic/gin"
	"github.com/miaojiaxi2004/go_server/model/common/response"
)

type SystemApiApi struct{}

func (s *SystemApiApi) CreateApi(c *gin.Context) {
	response.OkWithData("测试一下", c)
}

func (s *SystemApiApi) DeleteApi(c *gin.Context) {

}

func (s *SystemApiApi) GetApiList(c *gin.Context) {

}

func (s *SystemApiApi) GetApiById(c *gin.Context) {

}

func (s *SystemApiApi) UpdateApi(c *gin.Context) {

}

func (s *SystemApiApi) GetAllApis(c *gin.Context) {

}

func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {

}
