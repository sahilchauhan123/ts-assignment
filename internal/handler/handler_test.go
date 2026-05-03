package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"course-curriculum/internal/types"

	"github.com/gin-gonic/gin"
)

type fakeDB struct {
	createCourseName          string
	createCourseErr           error
	courses                   []types.Course
	createSubjectName         string
	createSubjectCourseID     int
	createSubjectErr          error
	subjects                  []types.Subject
	createLessonName          string
	createLessonSubjectID     int
	createLessonErr           error
	lessons                   []types.Lesson
	createStudentName         string
	createStudentCourseID     int
	createStudentErr          error
	students                  []types.Student
	createLessonCompletionSID int
	createLessonCompletionLID int
	createLessonCompletionErr error
	lessonProgressStudentID   int
	lessonProgress            []types.SubjectProgress
	lessonProgressErr         error
}

func (f *fakeDB) CreateCourse(name string) error {
	f.createCourseName = name
	return f.createCourseErr
}

func (f *fakeDB) GetCourses() ([]types.Course, error) {
	return f.courses, nil
}

func (f *fakeDB) CreateSubject(name string, courseID int) error {
	f.createSubjectName = name
	f.createSubjectCourseID = courseID
	return f.createSubjectErr
}

func (f *fakeDB) GetSubjects(subjectId int) ([]types.Subject, error) {
	return f.subjects, nil
}

func (f *fakeDB) CreateLesson(name string, subjectID int) error {
	f.createLessonName = name
	f.createLessonSubjectID = subjectID
	return f.createLessonErr
}

func (f *fakeDB) GetLessons() ([]types.Lesson, error) {
	return f.lessons, nil
}

func (f *fakeDB) CreateStudent(name string, courseID int) error {
	f.createStudentName = name
	f.createStudentCourseID = courseID
	return f.createStudentErr
}

func (f *fakeDB) CreateLessonCompletion(studentID int, lessonID int) error {
	f.createLessonCompletionSID = studentID
	f.createLessonCompletionLID = lessonID
	return f.createLessonCompletionErr
}

func (f *fakeDB) GetStudents() ([]types.Student, error) {
	return f.students, nil
}

func (f *fakeDB) GetLessonsProgressOfStudent(studentID int) ([]types.SubjectProgress, error) {
	f.lessonProgressStudentID = studentID
	return f.lessonProgress, f.lessonProgressErr
}

func init() {
	gin.SetMode(gin.TestMode)
}

func runTestRequest(t *testing.T, handler gin.HandlerFunc, method, target, body string, params ...gin.Param) (*httptest.ResponseRecorder, *gin.Context) {
	t.Helper()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(method, target, strings.NewReader(body))
	if body != "" {
		c.Request.Header.Set("Content-Type", "application/json")
	}
	if len(params) > 0 {
		c.Params = params
	}
	handler(c)
	return w, c
}

func decodeJSON[T any](t *testing.T, body *httptest.ResponseRecorder) T {
	t.Helper()
	var value T
	if err := json.Unmarshal(body.Body.Bytes(), &value); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}
	return value
}

func TestCreateCourseHandler(t *testing.T) {
	f := &fakeDB{}
	w, _ := runTestRequest(t, CreateCourse(f), http.MethodPost, "/course", `{"name":"Math"}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]any
	resp = decodeJSON[map[string]any](t, w)
	if resp["message"] != "Course created successfully" {
		t.Fatalf("unexpected response message: %v", resp["message"])
	}
	if f.createCourseName != "Math" {
		t.Fatalf("expected CreateCourse name Math, got %q", f.createCourseName)
	}
}

func TestGetCoursesHandler(t *testing.T) {
	f := &fakeDB{courses: []types.Course{{ID: 1, Name: "Math"}}}
	w, _ := runTestRequest(t, GetCourses(f), http.MethodGet, "/course", "")

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp []types.Course
	resp = decodeJSON[[]types.Course](t, w)
	if len(resp) != 1 || resp[0].Name != "Math" {
		t.Fatalf("unexpected response body: %#v", resp)
	}
}

func TestCreateSubjectHandler(t *testing.T) {
	f := &fakeDB{}
	w, _ := runTestRequest(t, CreateSubject(f), http.MethodPost, "/subject", `{"name":"Algebra","course_id":2}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]any
	resp = decodeJSON[map[string]any](t, w)
	if resp["message"] != "Subject created successfully" {
		t.Fatalf("unexpected response message: %v", resp["message"])
	}
	if f.createSubjectName != "Algebra" || f.createSubjectCourseID != 2 {
		t.Fatalf("CreateSubject called with wrong args: %v %v", f.createSubjectName, f.createSubjectCourseID)
	}
}

func TestGetSubjectsHandler(t *testing.T) {
	f := &fakeDB{subjects: []types.Subject{{ID: 5, Name: "Algebra"}}}
	w, _ := runTestRequest(t, GetSubjects(f), http.MethodGet, "/subject/5", "", gin.Param{Key: "id", Value: "5"})

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp []types.Subject
	resp = decodeJSON[[]types.Subject](t, w)
	if len(resp) != 1 || resp[0].Name != "Algebra" {
		t.Fatalf("unexpected response body: %#v", resp)
	}
}

func TestCreateLessonHandler(t *testing.T) {
	f := &fakeDB{}
	w, _ := runTestRequest(t, CreateLesson(f), http.MethodPost, "/lesson", `{"name":"Lesson 1","subject_id":7}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]any
	resp = decodeJSON[map[string]any](t, w)
	if resp["message"] != "Lesson created successfully" {
		t.Fatalf("unexpected response message: %v", resp["message"])
	}
	if f.createLessonName != "Lesson 1" || f.createLessonSubjectID != 7 {
		t.Fatalf("CreateLesson called with wrong args: %v %v", f.createLessonName, f.createLessonSubjectID)
	}
}

func TestGetLessonsHandler(t *testing.T) {
	f := &fakeDB{lessons: []types.Lesson{{ID: 11, Name: "Lesson 1", SubjectID: 7}}}
	w, _ := runTestRequest(t, GetLessons(f), http.MethodGet, "/lesson", "")

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp []types.Lesson
	resp = decodeJSON[[]types.Lesson](t, w)
	if len(resp) != 1 || resp[0].Name != "Lesson 1" {
		t.Fatalf("unexpected response body: %#v", resp)
	}
}

func TestCreateStudentHandler(t *testing.T) {
	f := &fakeDB{}
	w, _ := runTestRequest(t, CreateStudent(f), http.MethodPost, "/student", `{"name":"Alice","course_id":3}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]any
	resp = decodeJSON[map[string]any](t, w)
	if resp["message"] != "Student created successfully" {
		t.Fatalf("unexpected response message: %v", resp["message"])
	}
	if f.createStudentName != "Alice" || f.createStudentCourseID != 3 {
		t.Fatalf("CreateStudent called with wrong args: %v %v", f.createStudentName, f.createStudentCourseID)
	}
}

func TestGetStudentsHandler(t *testing.T) {
	f := &fakeDB{students: []types.Student{{ID: 4, Name: "Alice", CourseID: 3}}}
	w, _ := runTestRequest(t, GetStudents(f), http.MethodGet, "/student", "")

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp []types.Student
	resp = decodeJSON[[]types.Student](t, w)
	if len(resp) != 1 || resp[0].Name != "Alice" {
		t.Fatalf("unexpected response body: %#v", resp)
	}
}

func TestCreateLessonCompletionHandler(t *testing.T) {
	f := &fakeDB{}
	w, _ := runTestRequest(t, CreateLessonCompletion(f), http.MethodPost, "/lesson-completion", `{"student_id":4,"lesson_id":13}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]any
	resp = decodeJSON[map[string]any](t, w)
	if resp["message"] != "Lesson completion recorded successfully" {
		t.Fatalf("unexpected response message: %v", resp["message"])
	}
	if f.createLessonCompletionSID != 4 || f.createLessonCompletionLID != 13 {
		t.Fatalf("CreateLessonCompletion called with wrong args: %d %d", f.createLessonCompletionSID, f.createLessonCompletionLID)
	}
}

func TestGetLessonsProgressOfStudentHandler(t *testing.T) {
	f := &fakeDB{lessonProgress: []types.SubjectProgress{{SubjectID: 1, StudentID: 4, LessonProgress: []types.LessonCompletion{{LessonID: 7, IsCompleted: true}}}}}
	w, _ := runTestRequest(t, GetLessonsProgressOfStudent(f), http.MethodGet, "/student/4/lessons-progress", "", gin.Param{Key: "id", Value: "4"})

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp []types.SubjectProgress
	resp = decodeJSON[[]types.SubjectProgress](t, w)
	if len(resp) != 1 || resp[0].SubjectID != 1 || resp[0].StudentID != 4 {
		t.Fatalf("unexpected response body: %#v", resp)
	}
}
