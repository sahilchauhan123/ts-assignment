package database

import "course-curriculum/internal/types"

type Db interface {
	CreateCourse(name string) error
	GetCourses() ([]types.Course, error)
	CreateSubject(name string, courseID int) error
	GetSubjects(courseId int) ([]types.Subject, error)
	CreateLesson(name string, subjectID int) error
	GetLessons(subjectId int) ([]types.Lesson, error)
	CreateStudent(name string, courseID int) error
	CreateLessonCompletion(studentID int, lessonID int) error
	GetStudents() ([]types.Student, error)
	GetLessonsProgressOfStudent(studentID int) ([]types.SubjectProgress, error)
}
