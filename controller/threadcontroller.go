/*
|--------------------------------------------------------------------------
| Thread Controller
|--------------------------------------------------------------------------
|
| This controller is Thread(add, edit,delete)
|
*/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/model"
)

//get All Threads
func ThreadControllerActionIndex(c *gin.Context) {

	db := forumDB(c)

	var threads []Thread

	db.Debug().Find(&threads)

	for i, v := range threads {
		v.WithUser(db)
		threads[i] = v
	}

	if "application/json" == c.ContentType() {
		c.JSON(http.StatusOK, gin.H{
			"data": threads,
		})
		return
	}

	c.HTML(http.StatusOK, "thread/index.html", gin.H{
		"host":       "http://" + c.Request.Host,
		"css":        "http://" + c.Request.Host + "/assets/css/app.css",
		"js":         "http://" + c.Request.Host + "/assets/js/app.js",
		"threads":    threads,
		"ginContext": c,
	})
}

func ThreadControllerActionShow(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	thread, err := (&Thread{}).FindById(c, forumDB(c), "User", "Reply")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	if "application/json" == c.ContentType() {
		c.JSON(http.StatusOK, gin.H{
			"data": thread,
		})
		return
	}

	c.HTML(http.StatusOK, "thread/show.html", gin.H{
		"host":       "http://" + c.Request.Host,
		"css":        "http://" + c.Request.Host + "/assets/css/app.css",
		"js":         "http://" + c.Request.Host + "/assets/js/app.js",
		"thread":     thread,
		"ginContext": c,
	})
}

//ShowCreatePage
func ThreadControllerActionShowCreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "thread/create.html", gin.H{
		"host":       "http://" + c.Request.Host,
		"css":        "http://" + c.Request.Host + "/assets/css/app.css",
		"js":         "http://" + c.Request.Host + "/assets/js/app.js",
		"title":      "Welcome go forum",
		"content":    "You are logged in!",
		"ginContext": c,
	})
}

//ShowEditPage
func ThreadControllerActionShowEditPage(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	thread, err := (&Thread{}).FindById(c, forumDB(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "thread/edit.html", gin.H{
		"host":       "http://" + c.Request.Host,
		"css":        "http://" + c.Request.Host + "/assets/css/app.css",
		"js":         "http://" + c.Request.Host + "/assets/js/app.js",
		"thread":     thread,
		"ginContext": c,
	})
}

//Store
func ThreadControllerActionStore(c *gin.Context) {

	if err := ValidatePostFromParams(c, "title", "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	(&Thread{}).Create(c, forumDB(c))

	c.Redirect(http.StatusFound, "/t")
}

//Edit
func ThreadControllerActionEdit(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := ValidatePostFromParams(c, "title", "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := (&Thread{}).Edit(c, forumDB(c)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/t")
}

func ThreadControllerActionDestroy(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := (&Thread{}).DestroyById(c, forumDB(c)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/t")
}
