package handler

import (
	"course-curriculum/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type Db interface {
// 	CreateCourse(name string) error
// 	GetCourses() ([]types.Course, error)
// 	CreateSubject(name string, courseID int) error
// 	GetSubjects() ([]types.Subject, error)
// 	CreateLesson(name string, subjectID int) error
// 	GetLessons() ([]types.Lesson, error)
// 	CreateStudent(name string, courseID int) error
// 	CreateLessonCompletion(studentID int, lessonID int) error
// 	GetStudents() ([]types.Student, error)
// 	GetLessonsProgressOfStudent(studentID int) ([]types.LessonProgress, error)
// }

func CreateCourse(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RequestBody struct {
			Name string `json:"name"`
		}
		var reqBody RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
			return
		}
		err := db.CreateCourse(reqBody.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]any{"message": "Course created successfully"})
	}
}

func GetCourses(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := db.GetCourses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

func CreateSubject(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RequestBody struct {
			Name     string `json:"name"`
			CourseID int    `json:"course_id"`
		}
		var reqBody RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
			return
		}
		err := db.CreateSubject(reqBody.Name, reqBody.CourseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]any{"message": "Subject created successfully"})
	}
}

func GetSubjects(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId, err := strconv.Atoi(c.Param("course_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid Course ID"})
			return
		}
		subjects, err := db.GetSubjects(courseId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, subjects)
	}
}

func CreateLesson(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RequestBody struct {
			Name      string `json:"name"`
			SubjectID int    `json:"subject_id"`
		}
		var reqBody RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
			return
		}
		err := db.CreateLesson(reqBody.Name, reqBody.SubjectID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]any{"message": "Lesson created successfully"})
	}
}

func GetLessons(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		subjectId, err := strconv.Atoi(c.Param("subject_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid Subject ID"})
			return
		}
		lessons, err := db.GetLessons(subjectId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}

func CreateStudent(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RequestBody struct {
			Name     string `json:"name"`
			CourseID int    `json:"course_id"`
		}
		var reqBody RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
			return
		}
		err := db.CreateStudent(reqBody.Name, reqBody.CourseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]any{"message": "Student created successfully"})
	}
}

func GetStudents(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := db.GetStudents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, students)
	}
}

func CreateLessonCompletion(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RequestBody struct {
			StudentID int `json:"student_id"`
			LessonID  int `json:"lesson_id"`
		}
		var reqBody RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
			return
		}
		err := db.CreateLessonCompletion(reqBody.StudentID, reqBody.LessonID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]any{"message": "Lesson completion recorded successfully"})
	}
}

func GetLessonsProgressOfStudent(db database.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid student ID"})
			return
		}
		progress, err := db.GetLessonsProgressOfStudent(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, progress)
	}
}
