// пакет работы с БД postgres
package postgres
/* 
import (
	"context"
	"errors"
	
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	checkComment string = `select count("Id")  from censor
	where  position( lower(censor."Banned") in lower( $1) ) <>0   ; `
	addWord string = `INSERT INTO censor
	("Banned")
	VALUES( lower($1))  ON CONFLICT DO nothing;`
)

// объект для работы с  База данных новостей.
type DB struct {
	pool *pgxpool.Pool
}

// Конструктор  создания и подключения к БД.
func New() (*DB, error) {
	
	connstr := os.Getenv("CENSORDB")
	if connstr == "" {
		return nil, errors.New("не указано подключение к БД в  переменной CENSORDB, формат  postgres://[user]:[passwd]@[host]:[port]/[database]")
	}
	pool, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	db := DB{
		pool: pool,
	}
	return &db, nil
}

// закрытие соединения с БД.
func (db *DB) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}

// Проверка  комментария на запрещенные слова.
func (db *DB) CheckComment(c *pb.Comment) (bool, error) {
	var id int64
	//fmt.Println(comment)
	err := db.pool.QueryRow(context.Background(), checkComment, // 	(id_news, "Content", parent, "time",  "authorId")
		c.GetContent(),
	).Scan(&id)
	if err != nil {
		return true, err
	}
	return id == 0, nil
}

// Проверка  комментария на запрещенные слова.
func (db *DB) AddWord(word string) error {

	_, err := db.pool.Exec(context.Background(), addWord,
		word,
	)
	if err != nil {
		return err
	}
	return nil
}
 */