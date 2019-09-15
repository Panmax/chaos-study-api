package settings

type SettingSerializer struct {
	SettingModel
}

type SettingResponse struct {
	Count     uint8 `json:"count"`
	NotRepeat bool  `json:"not_repeat"`
}

func (s *SettingSerializer) Response() SettingResponse {
	response := SettingResponse{
		Count:     s.Count,
		NotRepeat: s.NotRepeat,
	}
	return response
}
