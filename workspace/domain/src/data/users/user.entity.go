package users

import (
	"strings"

	"dealls.test/domain/src/data/interactions"
	"dealls.test/domain/src/data/privilages"
	"dealls.test/material/src/contract"
	"dealls.test/material/src/helper"
	"gorm.io/gorm"
)

type Entity struct {
	Username string `json:"username,omitempty" gorm:"size:124;not null;"`
	Email    string `json:"email,omitempty" gorm:"size:124;not null;"`
	Password string `json:"password,omitempty" gorm:"size:255;not null;"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity

	// Relation Entity
	OwnerInteraction  []interactions.EntityModel `json:"owner_interaction" gorm:"foreignKey:OwnerID"`
	TargetInteraction []interactions.EntityModel `json:"target_interaction" gorm:"foreignKey:TargetID"`

	Privilages []privilages.EntityModel `json:"privilages" gorm:"foreignKey:UserID"`
}

func (EntityModel) TableName() string {
	return "users"
}
func (m *EntityModel) IsExist() bool {
	return m.ID != 0
}

func (m *EntityModel) BeforeCreate(cx *gorm.DB) error {
	m.Username = strings.ToLower(m.Username)
	m.Email = strings.ToLower(m.Email)

	hash, err := helper.Hash(m.Password)

	if err != nil {
		return err.Origin
	}

	m.Password = *hash

	return m.MetaEntity.BeforeCreate(cx)
}
