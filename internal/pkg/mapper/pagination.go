package mapper

import (
	"goProject/internal/config"
	"goProject/internal/dto"
)

func ToPaginationDto(p dto.PaginationDto ,total int) dto.PaginationResponseDto {
	limit := p.LimitOr()
	if limit > config.PaginationLimitMax {
		limit = config.PaginationLimitMax
	}
	offset := p.OffsetOr()

	return dto.PaginationResponseDto{
		Offset: offset,
		Limit: limit,
		Total: total,
	}
}