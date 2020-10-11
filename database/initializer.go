package database

/* var repository RepositoryController */

func init() {

	SetRepoController(&MonitorRepo{})
	repository = GetRepoController()
	SetHTTPController(&MonitorHttp{})
	httpCalls = GetHTTPController()

}
