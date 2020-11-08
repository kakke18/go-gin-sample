package dao

import (
	"encoding/json"
	"fmt"
	"go-gin-sample/models"
	"io/ioutil"
	"time"
)

const sourceFile = "json/user.json"

type UserRecord struct {
	Key          string
	Name         string
	EmailAddress string
	CreatedAt    string
}

func (r *UserRecord) toEntity() (*models.User, error) {
	createdAt, err := time.Parse("2006-01-02 15:04:05", r.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Key:          r.Key,
		Name:         r.Name,
		EmailAddress: r.EmailAddress,
		CreatedAt:    createdAt,
	}, nil
}

type UserJsonDao struct {
	Source string
}

func NewUserJsonDao() *UserJsonDao {
	return &UserJsonDao{}
}

func (dao *UserJsonDao) List() (*[]models.User, error) {
	records, err := getRecords()
	if err != nil {
		return nil, err
	}

	var results []models.User
	for _, record := range *records {
		entity, err := record.toEntity()
		if err != nil {
			return nil, err
		}
		results = append(results, *entity)
	}

	return &results, nil
}

func (dao *UserJsonDao) Get(id string) (*models.User, error) {
	records, err := getRecords()
	if err != nil {
		return nil, err
	}

	for _, record := range *records {
		if record.Key == id {
			user, err := record.toEntity()
			if err != nil {
				return nil, err
			}
			return user, nil
		}
	}

	return nil, fmt.Errorf("No such entity")
}

func getRecords() (*[]UserRecord, error) {
	fileData, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return nil, err
	}

	var records []UserRecord
	err = json.Unmarshal(fileData, &records)
	if err != nil {
		return nil, err
	}

	return &records, nil
}
