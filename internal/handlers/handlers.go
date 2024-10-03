package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"project01/internal/models"
	"strconv"
)

// 静态示例数据
var questions []models.Question

// 从文件加载问题列表
func LoadQuestionsFromFile(filePath string) error {
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

// 保存问题列表到文件
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

// 登录成功
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "登陆成功！")
}

// 获取问题列表
func ListQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}

// 创建新的问题
func CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.BindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	question.ID = len(questions) + 1
	questions = append(questions, question)
	c.JSON(http.StatusCreated, question)
}

// 获取编号为id的问题
func GetQuestion(c *gin.Context) {
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
}

// 修改编号为id的问题
func UpdateQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var question models.Question
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
}

// 创建答案
func CreateAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var answer models.Answer
	if err := c.BindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	answer.ID = len(questions[id-1].Answers) + 1
	answer.QuestionId = id
	questions[id-1].Answers = append(questions[id-1].Answers, answer)
	c.JSON(http.StatusCreated, answer)
}

// 获取编号为id的问题的答案列表
func GetAnswers(c *gin.Context) {
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
}

// 获取编号为id的问题的答案
func GetAnswer(c *gin.Context) {
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
					c.JSON(http.StatusOK, answer)
					return
				}
			}
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Question or Answer not found"})
}

// 修改编号为id的问题的答案
func UpdateAnswer(c *gin.Context) {
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
	var answer models.Answer
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
}

// 保存数据到文件
func Save(c *gin.Context) {
	if err := saveQuestionsToFile("questions.json"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save questions", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
