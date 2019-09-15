package courses

import (
	"database/sql/driver"
	"encoding/json"
)

type CourseSerializer struct {
	CourseModel
}

type CourseResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	TotalChapter uint16 `json:"total_chapter"`
	Url          string `json:"url"`
	Pick         uint8  `json:"pick"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

func (s *CourseSerializer) Response() CourseResponse {
	response := CourseResponse{
		ID:           s.ID,
		Name:         s.Name,
		TotalChapter: s.TotalChapter,
		Url:          s.Url,
		Pick:         s.Pick,
		CreatedAt:    s.CreatedAt.Unix(),
		UpdatedAt:    s.UpdatedAt.Unix(),
	}
	return response
}

type CoursePickResponse struct {
	Course   CourseResponse `json:"course"`
	Chapters []int          `json:"chapters"`
}

type CoursePickResults []CoursePickResponse

func (c CoursePickResults) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CoursePickResults) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
