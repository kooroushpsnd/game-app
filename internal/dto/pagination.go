package dto

import "goProject/internal/config"

type PaginationDto struct {
	Offset *int `query:"offset" validate:"omitempty,min=0"`
	Limit  *int `query:"limit"  validate:"omitempty,min=1,max=100"`
}

type PaginationResponseDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

func (p PaginationDto) OffsetOr() int {
	if p.Offset == nil {
		return config.PaginationOffsetDefault
	}
	return *p.Offset
}

func (p PaginationDto) LimitOr() int {
	if p.Limit == nil {
		return config.PaginationLimitDefault
	}
	return *p.Limit
}