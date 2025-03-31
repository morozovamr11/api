// Паттерн "Репозиторий" (Repository) — это слой абстракции между бизнес-логикой приложения и хранилищем данных (БД, API, файловая система). Он инкапсулирует логику работы с данными, предоставляя чистый интерфейс для их чтения, записи, обновления и удаления (CRUD).
package link

import (
	"api/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

// можно было сделать здесь проверку ошибки если  ошибка создания то пересоздать хэш для проблемы генерации дубликата хэша
func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash) //DB.First возвращает первое совпадение
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, id) //DB.First возвращает первое совпадение
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link) //обогащает ответ существующими данными Clauses(clause.Returning{})
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id) //удаляет из таблицы по id
	if result.Error != nil {
		return result.Error
	}
	return nil
}
