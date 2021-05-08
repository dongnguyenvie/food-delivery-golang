package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessReponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessReponse(data interface{}) *successRes {
	return NewSuccessReponse(data, nil, nil)
}
