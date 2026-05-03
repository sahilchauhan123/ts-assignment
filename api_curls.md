# Course Curriculum API - Curl Examples

This document provides curl commands to test all API endpoints. Assuming the server is running on `http://localhost:8080`.

## Create Course
```bash
curl -X POST http://localhost:8080/course \
  -H "Content-Type: application/json" \
  -d '{"name": "Mathematics"}'
```

## Get All Courses
```bash
curl -X GET http://localhost:8080/course
```

## Create Subject
```bash
curl -X POST http://localhost:8080/subject \
  -H "Content-Type: application/json" \
  -d '{"name": "Algebra", "course_id": 1}'
```

## Get Subjects by Course ID
```bash
curl -X GET http://localhost:8080/subjects/1
```

## Create Lesson
```bash
curl -X POST http://localhost:8080/lesson \
  -H "Content-Type: application/json" \
  -d '{"name": "Linear Equations", "subject_id": 1}'
```

## Get Lessons by Subject ID
```bash
curl -X GET http://localhost:8080/lesson/1
```

## Create Student
```bash
curl -X POST http://localhost:8080/student \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "course_id": 1}'
```

## Get All Students
```bash
curl -X GET http://localhost:8080/student
```

## Create Lesson Completion
```bash
curl -X POST http://localhost:8080/lesson-completion \
  -H "Content-Type: application/json" \
  -d '{"student_id": 1, "lesson_id": 1}'
```

## Get Student Lessons Progress
```bash
curl -X GET http://localhost:8080/student/1/lessons-progress
```