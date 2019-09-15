package courses

type CourseSerializer struct {
	CourseModel
}

type CourseResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Chapters  uint16 `json:"chapters"`
	Url       string `json:"url"`
	Pick      uint8  `json:"pick"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func (s *CourseSerializer) Response() CourseResponse {
	response := CourseResponse{
		ID:        s.ID,
		Name:      s.Name,
		Chapters:  s.Chapters,
		Url:       s.Url,
		Pick:      s.Pick,
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
	return response
}
