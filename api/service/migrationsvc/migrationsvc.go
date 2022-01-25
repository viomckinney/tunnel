package migrationsvc

import (
	"database/sql"
	"errors"

	"violet.wtf/tunnel/api/appctx"
)

const errVersion = -1

func EnsureMigrationTableExists() error {
	_, err := appctx.DB.Exec(
		"CREATE TABLE IF NOT EXISTS migration_status (version INT)",
	)
	return err
}

func GetVersion() (int, error) {
	row := appctx.DB.QueryRow("SELECT version FROM migration_status")

	if err := row.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return errVersion, err
	}

	var version int
	row.Scan(&version)

	return version, nil
}

func SetVersion(version int) error {
	currentVersion, err := GetVersion()
	if err != nil {
		return err
	}

	// Insert if not stored
	if currentVersion == 0 {
		appctx.DB.Exec("INSERT INTO migration_status (version) VALUES ($1)", 0)
	}

	_, err = appctx.DB.Exec(
		"UPDATE migration_status SET version = $1",
		version,
	)

	return err
}
