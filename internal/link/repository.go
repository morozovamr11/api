// Паттерн "Репозиторий" (Repository) — это слой абстракции между бизнес-логикой приложения и хранилищем данных (БД, API, файловая система). Он инкапсулирует логику работы с данными, предоставляя чистый интерфейс для их чтения, записи, обновления и удаления (CRUD).
package link

import "api/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
