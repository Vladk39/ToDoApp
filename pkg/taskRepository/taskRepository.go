package taskRepository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todoapi/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

const GetAllTasks = `select * from tasks`
const DeleteUser = `delete from tasks where id = $1`
const UpdateTask = `update tasks set title = $1, description = $2, status = $3, updated_at = NOW() where id = $4`
const AddTask = `insert into tasks (title, description, status) values ($1, $2, $3)`

// TaskResponse - структура ответа после создания или обновления задачи
// @Description Структура ответа для возвращаемых данных о задаче
type Task struct {
	// ID задачи
	// @example 1
	ID int `db:"id" json:"id,omitempty"`

	// Заголовок задачи
	// @example "Купить хлеб"
	Title string `db:"title" json:"title"`

	// Описание задачи
	// @example "Пойти в магазин и купить хлеб"
	Description string `db:"description" json:"description"`

	// Статус задачи
	// @example "new"
	Status string `db:"status" json:"status"`

	// Время создания задачи
	// @example "2023-04-01T12:00:00Z"
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`

	// Время последнего обновления задачи
	// @example "2023-04-02T14:00:00Z"
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type Repository struct {
	db *pgxpool.Pool
}

func ConnectDB(c *config.Config) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), c.DBconfig)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка подключения к БД")
	}

	return db, nil
}

func NewRepository(c *config.Config) (*Repository, error) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка создания репозитория")
	}

	if err := applyMigrations(c); err != nil {
		return nil, errors.Wrap(err, "ошибка применения миграций")
	}

	return &Repository{
		db: db,
	}, nil
}

func applyMigrations(c *config.Config) error {
	fmt.Println(c.DBconfig)
	db, err := sql.Open("postgres", c.DBconfig)
	if err != nil {
		return errors.Wrap(err, "не удалось открыть БД для миграций")
	}
	defer db.Close()
	goose.SetDialect("postgres")
	goose.Up(db, "../migrations")

	return nil
}

func (r *Repository) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := r.db.Query(ctx, GetAllTasks)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка выполнения запроса")
	}
	defer rows.Close()

	var Tasks []Task

	for rows.Next() {
		var t Task
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "ошибка чтения строк")
		}
		Tasks = append(Tasks, t)
	}

	return Tasks, nil
}

func (r *Repository) DeleteTask(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, DeleteUser, id)
	if err != nil {
		return errors.Wrap(err, "ошибка удаления задачи")
	}

	return err
}

func (r *Repository) UpdateTask(ctx context.Context, t Task) error {
	_, err := r.db.Exec(ctx, UpdateTask, t.Title, t.Description, t.Status, t.ID)
	if err != nil {
		return errors.Wrap(err, "ошибка при обновлении данных")
	}

	return nil
}

func (r *Repository) AddTask(ctx context.Context, t Task) error {
	_, err := r.db.Exec(ctx, AddTask, t.Title, t.Description, t.Status)
	if err != nil {
		return errors.Wrap(err, "ошибка добавления задачи")
	}

	return nil
}
