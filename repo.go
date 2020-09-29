package main

type ReposController interface {
	databaseGet(url *Urls) error
	databaseSave(url *Urls, id string) error
	databaseCreate(url *Urls) error
	databaseDelete(url *Urls, id string) error
}
type monitorRepo struct{}

var repo ReposController

func setRepoController(repoType ReposController) {
	repo = repoType
}

func (rp *monitorRepo) databaseGet(url *Urls) error {
	return db.Where("id = ?", url.ID).First(&url).Error
}

func (rp *monitorRepo) databaseSave(url *Urls, id string) error {
	return db.Save(&url).Error
}

func (rp *monitorRepo) databaseCreate(url *Urls) error {
	return db.Create(&url).Error
}

func (rp *monitorRepo) databaseDelete(url *Urls, id string) error {
	return db.Delete(&url).Error
}
