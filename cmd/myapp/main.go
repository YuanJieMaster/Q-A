package main

import (
	"github.com/gin-gonic/gin"
	"project01/internal/handlers"
)

func main() {
	r := gin.Default()

	// 从文件加载问题列表
	if err := handlers.LoadQuestionsFromFile("questions.json"); err != nil {
		panic("Failed to load questions from file: " + err.Error())
	} else {
		println("Load questions from file success")
	}

	// 登录成功
	r.GET("/", handlers.Login)

	// 获取问题列表
	r.GET("/api/question", handlers.ListQuestions)

	// 创建新的问题
	r.POST("/api/question", handlers.CreateQuestion)

	// 获取编号为id的问题
	r.GET("/api/question/:id", handlers.GetQuestion)

	// 修改编号为id的问题
	r.PUT("/api/question/:id", handlers.UpdateQuestion)

	// 创建答案
	r.POST("/api/question/:id/answer", handlers.CreateAnswer)

	// 获取编号为id的问题的答案列表
	r.GET("/api/question/:id/answer", handlers.GetAnswers)

	// 获取编号为id的问题的答案
	r.GET("/api/question/:id/answer/:answerID", handlers.GetAnswer)

	// 修改编号为id的问题的答案
	r.PUT("/api/question/:id/answer/:answerID", handlers.UpdateAnswer)

	// 保存数据到文件
	r.POST("/api/save", handlers.Save)

	panic(r.Run(":8080")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
