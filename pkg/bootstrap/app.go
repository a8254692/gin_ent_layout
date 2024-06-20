package bootstrap

import (
	"back-end/merchant/internal/data"
)

type App struct {
	Data *data.Data
}

func (s *App) AddData(data *data.Data) {
	if s.Data == nil {
		s.Data = data
	}
	return
}
