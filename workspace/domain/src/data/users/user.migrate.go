package users

import (
	"fmt"
	"time"

	"dealls.test/material/src/client/postgresql/executor"
	"dealls.test/material/src/client/postgresql/options"
	"dealls.test/material/src/contract"
	"gorm.io/gorm"
)

type migrator struct {
	*gorm.DB
}

func Migrate(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&migrator{
		dbi,
	})
}

// Instance implements execution.
func (um *migrator) Instance() *gorm.DB {
	return um.DB
}

// Table implements migrator
func (*migrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(EntityModel{}); !ok {
		if err := m.CreateTable(EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*migrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(EntityModel{}); ok {
		if err := m.DropTable(EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*migrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*migrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*migrator) Index(m *executor.Migrator) error {
	if ok := m.HasIndex(EntityModel{}, "un_email"); !ok {

		if err := m.CreateIndex(EntityModel{}, options.UNIQUE, "un_email", "email"); err != nil {
			return err
		}

	}

	if ok := m.HasIndex(EntityModel{}, "idx_username_email_search"); !ok {

		if err := m.CreateIndex(EntityModel{}, options.INDEX, "idx_username_email_search", "email", "username", "deleted_at"); err != nil {
			return err
		}

	}

	return nil
}

// Seeder implements executor.Actions.
func (*migrator) Seeder(m *executor.Migrator) error {
	go func() {
		m.Insert(&[]EntityModel{
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLH",
					},
				},
				Entity: Entity{
					Email:    "admin@admin.id",
					Password: "useradmin",
					Username: "useradmin",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLZ",
					},
				},
				Entity: Entity{
					Email:    "admin1@admin.id",
					Password: "useradmin1",
					Username: "useradmin1",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLp",
					},
				},
				Entity: Entity{
					Email:    "admin2@admin.id",
					Password: "useradmin2",
					Username: "useradmin2",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLM",
					},
				},
				Entity: Entity{
					Email:    "admin3@admin.id",
					Password: "useradmin3",
					Username: "useradmin3",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLB",
					},
				},
				Entity: Entity{
					Email:    "admin4@admin.id",
					Password: "useradmin4",
					Username: "useradmin4",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLS",
					},
				},
				Entity: Entity{
					Email:    "admin5@admin.id",
					Password: "useradmin5",
					Username: "useradmin5",
				},
			},
		})
	}()

	userss := []EntityModel{}
	for i := 0; i < 100; i++ {
		userss = append(userss, EntityModel{
			Entity: Entity{
				Email:    fmt.Sprintf("user-%d-%d@user.co.id", i, time.Now().Unix()),
				Password: "userpassword",
				Username: fmt.Sprintf("user-%d-%d", i, time.Now().Unix()),
			},
		})
	}

	if err := m.Insert(&userss); err != nil {
		return err
	}

	return nil
}
