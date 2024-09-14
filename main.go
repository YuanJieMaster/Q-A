package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Question struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Detail      string   `json:"detail"`
	Author      string   `json:"author"`
	CreatedAt   string   `json:"created_at"`
	AuthorEmail string   `json:"author_email"`
	Answers     []Answer `json:"answers"`
}

type Answer struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
	QuestionId  int    `json:"question_id,omitempty"`
	IsBest      bool   `json:"is_best,omitempty"`
}

// 静态示例数据
var questions []Question

// 从文件加载问题列表
func loadQuestionsFromFile(filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &questions)
	if err != nil {
		return err
	}

	return nil
}

func saveQuestionsToFile(filePath string) error {
	file, err := json.MarshalIndent(questions, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := gin.Default()

	// 从文件加载问题列表
	if err := loadQuestionsFromFile("questions.json"); err != nil {
		panic("Failed to load questions from file: " + err.Error())
	} else {
		println("Load questions from file success")
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "登陆成功！")
	})

	// 获取问题列表
	r.GET("/api/question", func(c *gin.Context) {
		c.JSON(http.StatusOK, questions)
	})

	// 创建新的问题
	r.POST("/api/question", func(c *gin.Context) {
		var question Question
		if err := c.BindJSON(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		question.ID = len(questions) + 1
		questions = append(questions, question)
		c.JSON(http.StatusCreated, question)
	})

	// 获取编号为id的问题
	r.GET("/api/question/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for _, question := range questions {
			if question.ID == id {
				c.JSON(http.StatusOK, question)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Question not found"})
	})

	// 修改编号为id的问题
	r.PUT("/api/question/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var question Question
		if err := c.BindJSON(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		for _, q := range questions {
			if q.ID == id {
				//questions[i] = question
				c.JSON(http.StatusOK, "question updated successfully")
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Question not found"})
	})

	// 创建答案
	r.POST("/api/question/:id/answer", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var answer Answer
		if err := c.BindJSON(&answer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		answer.ID = len(questions[id-1].Answers) + 1
		answer.QuestionId = id
		questions[id-1].Answers = append(questions[id-1].Answers, answer)
		c.JSON(http.StatusCreated, answer)
	})

	// 获取编号为id的问题的答案列表
	r.GET("/api/question/:id/answer", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for _, question := range questions {
			if question.ID == id {
				c.JSON(http.StatusOK, question.Answers)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Question not found"})
	})

	// 获取编号为id的问题的答案
	r.GET("/api/question/:id/answer/:answerID", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		answerID, err := strconv.Atoi(c.Param("answerID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Answer ID"})
			return
		}
		for _, question := range questions {
			if question.ID == id {
				for _, answer := range question.Answers {
					if answer.ID == answerID {
						c.JSON(http.StatusOK, "answer updated successfully")
						return
					}
				}
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Question or Answer not found"})
	})

	// 修改编号为id的问题的答案
	r.PUT("/api/question/:id/answer/:answerID", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		answerID, err := strconv.Atoi(c.Param("answerID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Answer ID"})
			return
		}
		var answer Answer
		if err := c.BindJSON(&answer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		for _, question := range questions {
			if question.ID == id {
				for _, a := range question.Answers {
					if a.ID == answerID {
						//questions[i].Answers[j] = answer
						c.JSON(http.StatusOK, "answer updated successfully")
						return
					}
				}
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Question or Answer not found"})
	})

	// 保存数据到文件
	r.POST("/api/save", func(c *gin.Context) {
		if err := saveQuestionsToFile("questions.json"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save questions", "details": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
	})

	panic(r.Run(":8080")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
