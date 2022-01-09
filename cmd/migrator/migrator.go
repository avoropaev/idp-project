package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	pg_sugar "git.vseinstrumenti.net/golang/pg-sugar"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	//postgres driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//Register for file migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	migrationsDir = "migrations"

	argFilesCreate = "new"
	argTo          = "to"
	argForce       = "force"
	argUp          = "up"
	argVersion     = "version"
)

type logger struct {
	z *zerolog.Logger
}

func (l logger) Printf(format string, v ...interface{}) {
	l.z.Printf(format, v...)
}

func (logger) Verbose() bool {
	return true
}

func NewCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "migrate",
		Short:        "Run database migrations",
		SilenceUsage: true,
		Args:         cobra.ExactValidArgs(1),
		ValidArgs:    []string{argFilesCreate, argTo, argForce, argUp, argVersion},

		RunE: func(cmd *cobra.Command, args []string) error {
			l := log.Ctx(ctx).With().Str("component", "migrate").Logger()

			c := pg_sugar.FromEnv()

			l.Info().Msg(fmt.Sprintf("Postgres DSN: %s", c.FormatDSN()))

			if args[0] == argFilesCreate {
				if err := handleArgFilesCreate(); err != nil {
					l.Error().Err(err).Send()
					return err
				}

				return nil
			}

			start := time.Now()

			migrant, err := getMigrant(ctx)
			if err != nil {
				return err
			}

			defer func() {
				mErr := &multierror.Error{}

				if err != nil {
					mErr = multierror.Append(mErr, err)
				}

				sourceErr, dbErr := migrant.Close()

				if sourceErr != nil {
					mErr = multierror.Append(mErr, sourceErr)
				}

				if dbErr != nil {
					mErr = multierror.Append(mErr, dbErr)
				}

				err = mErr.ErrorOrNil()

				l.Info().Dur("duration: ", time.Since(start)).Send()
			}()

			switch args[0] {
			case argTo:
				if err := handleToVersion(migrant); err != nil {
					return err
				}
			case argForce:
				if err := handleForceVersion(migrant); err != nil {
					return err
				}
			case argUp:
				if err := up(migrant); err != nil {
					return err
				}
			case argVersion:
				version, dirty, err := migrant.Version()
				if err != nil {
					return err
				}

				fmt.Printf("Current version: %d\n", version)

				if dirty {
					fmt.Println("dirty")
				} else {
					fmt.Println("clean")
				}
			}

			return nil
		},
	}

	return cmd
}

func getMigrant(ctx context.Context) (*migrate.Migrate, error) {
	c := pg_sugar.FromEnv()

	m, err := migrate.New("file://"+migrationsDir, c.FormatDSN())
	if err != nil {
		return nil, err
	}

	m.Log = logger{log.Ctx(ctx)}

	return m, nil
}

func up(m *migrate.Migrate) (err error) {
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func handleToVersion(m *migrate.Migrate) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter version which you want use")

	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println("Migrating...")

	version, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return err
	}

	return toVersion(m, uint(version))
}

func toVersion(m *migrate.Migrate, version uint) (err error) {
	if err := m.Migrate(version); err != nil {
		return err
	}

	return nil
}

func handleForceVersion(m *migrate.Migrate) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter version to force use")

	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println("Migrating...")

	version, err := strconv.Atoi(strings.TrimSpace(text))

	if err != nil {
		return err
	}

	return forceVersion(m, version)
}

func forceVersion(m *migrate.Migrate, version int) (err error) {
	if err := m.Force(version); err != nil {
		return err
	}

	return nil
}

func handleArgFilesCreate() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter migration name")

	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println("Creating migrations...")

	name := strings.TrimSpace(text)

	err = createMigrationFiles(migrationsDir, name, "sql")

	if err != nil {
		return err
	}

	return nil
}

func createMigrationFiles(migrationsDir, name, extension string) error {
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(migrationsDir, os.FileMode(0755)); err != nil {
			return err
		}
	}

	filename := fmt.Sprintf("%s_%s.", time.Now().Format("20060102150405"), name)

	for _, value := range []string{filename + "up." + extension, filename + "down." + extension} {
		if _, err := os.Create(path.Join(migrationsDir, value)); err != nil {
			return err
		}
	}

	return nil
}
