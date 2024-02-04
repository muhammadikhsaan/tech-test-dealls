package privilages

import (
	"dealls.test/material/src/client/postgresql/executor"
	"dealls.test/material/src/client/postgresql/options"
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
func (m *migrator) Instance() *gorm.DB {
	return m.DB
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
	if ok := m.HasIndex(EntityModel{}, "idx_search_by_user"); !ok {
		if err := m.CreateIndex(EntityModel{}, options.INDEX, "idx_search_by_user", "user_id"); err != nil {
			return err
		}
	}

	return nil
}

// Seeder implements executor.Actions.
func (*migrator) Seeder(m *executor.Migrator) error {
	return nil
}
