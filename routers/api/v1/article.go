package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/huangzhuo492008824/go-gin-example/models"
	"github.com/huangzhuo492008824/go-gin-example/pkg/app"
	"github.com/huangzhuo492008824/go-gin-example/pkg/e"
	"github.com/huangzhuo492008824/go-gin-example/pkg/setting"
	"github.com/huangzhuo492008824/go-gin-example/pkg/util"
	"github.com/huangzhuo492008824/go-gin-example/service/article_service"
	"github.com/unknwon/com"
)

// @Summary 获取文章列表
// @Produce  json
// @Security ApiKeyAuth
// @Param title query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	title := c.Query("title")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if title != "" {
		maps["title"] = title
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 创建文章
// @Produce  json
// @Security ApiKeyAuth
// @Param article body models.Article true "Article"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {

	// @Param title body string true "Title"
	// @Param desc body string true "Desc"
	// @Param tag_id body int false "TagId"
	// @Param state body int false "State"
	// @Param content body string false "Content"
	// @Param created_by body string false "CreatedBy"
	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Printf("%v", &json)
	log.Printf("%T", int(json["tag_id"].(float64)))
	tagId := int(json["tag_id"].(float64))
	title := json["title"].(string)
	desc := json["desc"].(string)
	content := json["content"].(string)
	state := int(json["tag_id"].(float64))
	createdBy := json["created_by"].(string)

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签id必须大于1")
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题长度最长为100个字符")
	valid.Required(desc, "desc").Message("描述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "name").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0/1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			code = e.SUCCESS
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			data["tag_id"] = tagId
			models.AddArticle(data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 修改文章
// @Produce  json
// @Security ApiKeyAuth
// @Param title query string false "Title"
// @Param desc query string false "Desc"
// @Param content query string false "Content"
// @Param id path int false "ID"
// @Param tag_id query int false "TagID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许为0/1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "name").Message("修改人最长为100个字符")
	valid.MaxSize(title, 100, "name").Message("名称最长为100个字符")
	valid.MaxSize(desc, 100, "name").Message("描述最长为100个字符")
	valid.MaxSize(content, 100, "name").Message("内容最长为100个字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if tagId > 0 {
				data["tag_id"] = tagId
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}
			if state != -1 {
				data["state"] = state
			}
			data["modified_by"] = modifiedBy
			models.EditArticle(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// @Summary 删除文章标签
// @Security ApiKeyAuth
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 获取文章详情
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}
