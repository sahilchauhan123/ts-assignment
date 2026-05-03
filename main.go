package main

import (
	"course-curriculum/database/sqlite"
	"course-curriculum/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	sqlite := sqlite.NewSqlite()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.POST("/course", handler.CreateCourse(sqlite))
	r.GET("/course", handler.GetCourses(sqlite))
	r.POST("/subject", handler.CreateSubject(sqlite))
	r.GET("/subject", handler.GetSubjects(sqlite))
	r.POST("/lesson", handler.CreateLesson(sqlite))
	r.GET("/lesson", handler.GetLessons(sqlite))
	r.POST("/student", handler.CreateStudent(sqlite))
	r.GET("/student", handler.GetStudents(sqlite))
	r.POST("/lesson-completion", handler.CreateLessonCompletion(sqlite))
	r.GET("/student/:id/lessons-progress", handler.GetLessonsProgressOfStudent(sqlite))

	r.Run(":8080")
}
