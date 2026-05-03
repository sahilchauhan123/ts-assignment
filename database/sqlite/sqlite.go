package sqlite

import (
	"course-curriculum/internal/types"
	"database/sql"
	"fmt"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func NewSqlite() *Sqlite {

	sqliteDb, err := sql.Open("sqlite3", "file:mydb.sqlite?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}

	type table struct {
		name        string
		createQuery string
	}

	tables := []table{
		{
			name: "courses",
			createQuery: `CREATE TABLE IF NOT EXISTS courses (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL UNIQUE
			)`,
		},
		{
			name: "subjects",
			createQuery: `CREATE TABLE IF NOT EXISTS subjects (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				course_id INTEGER NOT NULL,
				FOREIGN KEY (course_id) REFERENCES courses(id)
			);`,
		},
		{
			name: "lessons",
			createQuery: `CREATE TABLE IF NOT EXISTS lessons (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				subject_id INTEGER NOT NULL,
				FOREIGN KEY (subject_id) REFERENCES subjects(id)
			);`,
		},
		{
			name: "lesson_completions",
			createQuery: `CREATE TABLE IF NOT EXISTS lesson_completions (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				student_id INTEGER NOT NULL,
				lesson_id INTEGER NOT NULL,
				FOREIGN KEY (student_id) REFERENCES students(id),
				FOREIGN KEY (lesson_id) REFERENCES lessons(id)
			);`,
		},
		{
			name: "students",
			createQuery: `CREATE TABLE IF NOT EXISTS students (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				course_id INTEGER NOT NULL,
				FOREIGN KEY (course_id) REFERENCES courses(id)
			);`,
		},
	}

	for _, table := range tables {
		_, err = sqliteDb.Exec(table.createQuery)
		if err != nil {
			panic(fmt.Errorf("error in creating %s table: %w", table.name, err))
		}
	}

	return &Sqlite{
		db: sqliteDb,
	}
}

func (s *Sqlite) CreateCourse(name string) error {
	_, err := s.db.Exec("INSERT INTO courses (name) VALUES (?)", name)
	return err
}

func (s *Sqlite) GetCourses() ([]types.Course, error) {
	rows, err := s.db.Query("SELECT id, name FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []types.Course
	for rows.Next() {
		var course types.Course
		if err := rows.Scan(&course.ID, &course.Name); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (s *Sqlite) CreateSubject(name string, courseID int) error {
	_, err := s.db.Exec("INSERT INTO subjects (name, course_id) VALUES (?, ?)", name, courseID)
	return err
}

func (s *Sqlite) GetSubjects(courseId int) ([]types.Subject, error) {
	rows, err := s.db.Query("SELECT id, name FROM subjects WHERE course_id = ?", courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []types.Subject
	for rows.Next() {
		var subject types.Subject
		if err := rows.Scan(&subject.ID, &subject.Name); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

func (s *Sqlite) CreateLesson(name string, subjectID int) error {
	_, err := s.db.Exec("INSERT INTO lessons (name, subject_id) VALUES (?, ?)", name, subjectID)
	return err
}

func (s *Sqlite) GetLessons(subjectId int) ([]types.Lesson, error) {
	rows, err := s.db.Query("SELECT id, name, subject_id FROM lessons WHERE subject_id = ?", subjectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []types.Lesson
	for rows.Next() {
		var lesson types.Lesson
		if err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.SubjectID); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func (s *Sqlite) CreateStudent(name string, courseID int) error {
	_, err := s.db.Exec("INSERT INTO students (name, course_id) VALUES (?, ?)", name, courseID)
	return err
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	rows, err := s.db.Query("SELECT id, name, course_id FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.CourseID); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *Sqlite) CreateLessonCompletion(studentID int, lessonID int) error {
	_, err := s.db.Exec("INSERT INTO lesson_completions (student_id, lesson_id) VALUES (?, ?)", studentID, lessonID)
	return err
}

func (s *Sqlite) GetLessonsProgressOfStudent(studentID int) ([]types.SubjectProgress, error) {
	rows, err := s.db.Query(`
		SELECT subjects.id AS subject_id,
		       lessons.id AS lesson_id,
		       lessons.name AS lesson_name,
		       lc.id AS completion_id
		FROM students
		JOIN subjects ON subjects.course_id = students.course_id
		LEFT JOIN lessons ON lessons.subject_id = subjects.id 
		LEFT JOIN lesson_completions lc ON lc.lesson_id = lessons.id AND lc.student_id = ?
		WHERE students.id = ?
		ORDER BY subjects.id, lessons.id
	`, studentID, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	progressMap := make(map[int]*types.SubjectProgress)
	for rows.Next() {
		var subjectID int
		var lessonID sql.NullInt64
		var lessonName sql.NullString
		var completionID sql.NullInt64
		if err := rows.Scan(&subjectID, &lessonID, &lessonName, &completionID); err != nil {
			return nil, err
		}

		sp, ok := progressMap[subjectID]
		if !ok {
			sp = &types.SubjectProgress{
				SubjectID:      subjectID,
				StudentID:      studentID,
				LessonProgress: []types.LessonCompletion{},
			}
			progressMap[subjectID] = sp
		}

		if !lessonID.Valid {
			continue
		}

		lessonCompletion := types.LessonCompletion{
			ID:          0,
			StudentID:   studentID,
			LessonID:    int(lessonID.Int64),
			LessonName:  lessonName.String,
			IsCompleted: completionID.Valid && completionID.Int64 != 0,
		}
		if completionID.Valid {
			lessonCompletion.ID = int(completionID.Int64)
		}
		sp.LessonProgress = append(sp.LessonProgress, lessonCompletion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	progress := make([]types.SubjectProgress, 0, len(progressMap))
	for _, sp := range progressMap {
		progress = append(progress, *sp)
	}
	sort.Slice(progress, func(i, j int) bool {
		return progress[i].SubjectID < progress[j].SubjectID
	})

	return progress, nil
}
