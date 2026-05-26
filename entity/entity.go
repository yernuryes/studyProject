package entity

import "time"

type Course struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Chapter struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	CourseID    int64     `json:"course_id"`
	Course      *Course   `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Lesson struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Order       int       `json:"order"`
	ChapterID   int64     `json:"chapter_id"`
	Chapter     *Chapter  `json:"chapter,omitempty" gorm:"foreignKey:ChapterID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
