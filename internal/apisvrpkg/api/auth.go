package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/model"
	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/chanprogo/somemodule/pkg/util/jwtutil"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if !ok {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}

	isExist, err := model.CheckAuth(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, nil)
		return
	}

	token, err := jwtutil.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, constant.RESPONSE_CODE_OK, map[string]string{
		"token": token,
	})
}
