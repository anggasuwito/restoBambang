package repositories

import "gomux/main/apimaster/models"

//MenuRepository MenuRepository
type MenuRepository interface {
	AddMenu(newMenu models.Menu) error
	GetMenuByID(id string) (models.Menu, error)
	UpdateMenusByID(id string, changeMenu models.Menu) error
	DeleteDataMenuByID(id string) error
	GetAllMenus() ([]*models.Menu, error)
}
