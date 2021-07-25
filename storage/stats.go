package storage

import (
	"fmt"

	"github.com/fathisiddiqi/go-mini-kraicklist/cmd/rest/models"
)

func (s *Storage) GetStatistics() (*models.Statistic, error) {
	var stat models.Statistic
	
	query := `SELECT COUNT(id) FROM ads`
	row := s.db.QueryRow(query)
	err := row.Scan(&stat.TotalAds)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &stat, nil
}