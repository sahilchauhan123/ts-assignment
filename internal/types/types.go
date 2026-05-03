package types

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Lesson struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SubjectID int    `json:"subject_id"`
}

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CourseID int    `json:"course_id"`
}

type LessonCompletion struct {
	ID          int  `json:"id"`
	StudentID   int  `json:"student_id"`
	LessonID    int  `json:"lesson_id"`
	IsCompleted bool `json:"is_completed,omitempty"`
}

type SubjectProgress struct {
	SubjectID      int                `json:"subject_id"`
	StudentID      int                `json:"student_id"`
	LessonProgress []LessonCompletion `json:"lesson_progress"`
}
