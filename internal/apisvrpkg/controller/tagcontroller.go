package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/common"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/service"
	"github.com/chanprogo/somemodule/pkg/conf/iconf"
	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type TagController struct {
	app.Controller
}

func (t *TagController) Router(e *gin.Engine) {
	group := e.Group("/api")
	{
		group.POST("/tags", t.AddTag)          // 新建标签
		group.DELETE("/tags/:id", t.DeleteTag) // 删除指定标签
		group.PUT("/tags/:id", t.EditTag)      // 更新指定标签
		group.GET("/tags", t.GetTags)          // 获取标签列表
	}
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Router /api/tags [post]
func (t *TagController) AddTag(c *gin.Context) {

	var form AddTagForm
	_, errCode := common.BindAndValid(c, &form)
	if errCode != constant.RESPONSE_CODE_OK {
		t.RespErr(c, errCode)
		return
	}

	tagService := service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}

	exists, err := tagService.ExistByName()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "获取已存在标签失败")
		return
	}
	if exists {
		t.RespErr(c, "已存在该标签名称")
		return
	}

	err = tagService.Add()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "新增标签失败")
		return
	}
	t.RespOK(c)
}

// @Produce  json
// @Param id path int true "ID"
// @Router /api/tags/{id} [delete]
func (t *TagController) DeleteTag(c *gin.Context) {

	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		common.MarkErrors(valid.Errors)
		// http.StatusBadRequest
		t.RespErr(c, constant.INVALID_PARAMS)
	}

	tagService := service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "获取已存在标签失败")
		return
	}
	if !exists {
		t.RespErr(c, "该标签不存在")
		return
	}

	if err := tagService.Delete(); err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "删除标签失败")
		return
	}
	t.RespOK(c)
}

type EditTagForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Router /api/tags/{id} [put]
func (t *TagController) EditTag(c *gin.Context) {

	form := EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}

	_, errCode := common.BindAndValid(c, &form)
	if errCode != constant.RESPONSE_CODE_OK {
		t.RespErr(c, errCode)
		return
	}

	tagService := service.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := tagService.ExistByID()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "获取已存在标签失败")
		return
	}
	if !exists {
		t.RespErr(c, "该标签不存在")
		return
	}

	err = tagService.Edit()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "修改标签失败")
		return
	}

	t.RespOK(c)
}

// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Router /api/tags [get]
func (t *TagController) GetTags(c *gin.Context) {

	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * iconf.AppSetting.PageSize
	}
	tagService := service.Tag{
		Name:     name,
		State:    state,
		PageNum:  result,
		PageSize: iconf.AppSetting.PageSize,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "获取所有标签失败")
		return
	}

	count, err := tagService.Count()
	if err != nil {
		// http.StatusInternalServerError
		t.RespErr(c, "统计标签失败")
		return
	}

	t.Put(c, "lists", tags)
	t.Put(c, "total", count)
	t.RespOK(c)
}
