package plans

type PlanSerializer struct {
	PlanModel
}

type PlanResponse struct {
	Count     uint8 `json:"count"`
	NotRepeat bool  `json:"not_repeat"`
}

func (s *PlanSerializer) Response() PlanResponse {
	response := PlanResponse{
		Count:     s.Count,
		NotRepeat: s.NotRepeat,
	}
	return response
}
