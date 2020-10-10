package dao

import (
	"encoding/json"
	"fmt"
	"go-gin-sample/models"
	"io/ioutil"
	"time"
)

const sourceFile = "json/member.json"

type MemberRecord struct {
	Key          string
	Name         string
	EmailAddress string
	CreatedAt    string
}

func (r *MemberRecord) toEntity() (*models.Member, error) {
	createdAt, err := time.Parse("2006-01-02 15:04:05", r.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &models.Member{
		Key:          r.Key,
		Name:         r.Name,
		EmailAddress: r.EmailAddress,
		CreatedAt:    createdAt,
	}, nil
}

type MemberJsonDao struct {
	Source string
}

func NewMemberJsonDao() *MemberJsonDao {
	return &MemberJsonDao{}
}

func (dao *MemberJsonDao) Get() (*[]models.Member, error) {
	records, err := getRecords()
	if err != nil {
		return nil, err
	}

	var results []models.Member
	for _, record := range *records {
		entity, err := record.toEntity()
		if err != nil {
			return nil, err
		}
		results = append(results, *entity)
	}

	return &results, nil
}

func (dao *MemberJsonDao) GetById(id string) (*models.Member, error) {
	records, err := getRecords()
	if err != nil {
		return nil, err
	}

	for _, record := range *records {
		if record.Key == id {
			member, err := record.toEntity()
			if err != nil {
				return nil, err
			}
			return member, nil
		}
	}

	return nil, fmt.Errorf("No such entity")
}

func getRecords() (*[]MemberRecord, error) {
	fileData, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return nil, err
	}

	var records []MemberRecord
	err = json.Unmarshal(fileData, &records)
	if err != nil {
		return nil, err
	}

	return &records, nil
}
