package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fathisiddiqi/go-mini-kraicklist/cmd/rest/models"
)

func (s *Storage) GetList(limit int) ([]models.List, error) {
	// execute sql query
	queryGet := fmt.Sprintf("SELECT id, title, body, tags, created_at, updated_at FROM ads ORDER BY created_at DESC LIMIT %d", limit)

	rows, err := s.db.Query(queryGet)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due: %v", err)
	}
	defer rows.Close()
	// iterate rows
	var lists []models.List
	for rows.Next() {
		var list models.List
		var tagsJSON []byte
		err = rows.Scan(&list.ID, &list.Title, &list.Body, &tagsJSON, &list.CreatedAt, &list.UpdatedAt)
		if err != nil {
			continue
		}
		if len(tagsJSON) > 0 {
			err = json.Unmarshal(tagsJSON, &list.Tags)
			if err != nil {
				continue
			}
		}
		lists = append(lists, list)
	}
	// check error in rows iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate rows due: %v", err)
	}

	return lists, nil
}

func (s *Storage) AddList(list models.List) (*models.List, error) {
	// encode tags
	tagsJSON, _ := json.Marshal(list.Tags)
	// get current timestamp for `created_at` value
	nowTs := time.Now().Unix()

	// execute sql query
	query := `INSERT INTO ads (title, body, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, list.Title, list.Body, string(tagsJSON), nowTs, nowTs)
	if err != nil {
		return nil, fmt.Errorf("unable to execute insert due: %v", err)
	}
	// get record id
	recordID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("unable to get last inserted id due: %v", err)
	}
	// assign id & created at value to list
	list.ID = int(recordID)
	list.CreatedAt = nowTs
	list.UpdatedAt = nowTs

	return &list, nil
}

func (s *Storage) FindListById(id int) (int, error) {
	rows := s.db.QueryRow("SELECT id FROM ads WHERE id=?", id)
	err := rows.Scan(&id)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("data is not found")
	}

	return id, nil
}

func (s *Storage) UpdateList(list models.List) (*models.List, error) {
	tagsJSON, _ := json.Marshal(list.Tags)

	nowTs := time.Now().Unix()

	query := `UPDATE ads SET title=?, body=?, tags=?, updated_at=? WHERE id=?`
	_, err := s.db.Exec(query, list.Title, list.Body, tagsJSON, nowTs, list.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to execute update due: %v", err)
	}

	list.UpdatedAt = nowTs

	return &list, nil
}

func (s *Storage) DeleteList(id int) (*models.List, error) {
	query := `DELETE FROM ads WHERE id=?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return nil, fmt.Errorf("unable to execute delete due: %v", err)
	}

	return &models.List{ID: id}, nil
}