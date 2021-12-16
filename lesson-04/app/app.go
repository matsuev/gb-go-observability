package app

// TracingApp
type TracingApp struct {
	tracer  *AppTracer
	config  *AppConfig
	router  *AppRouter
	storage *AppStorage
}

// Create function
func Create() (a *TracingApp, err error) {
	a = new(TracingApp)

	a.config = CreateAppConfig()

	if a.tracer, err = CreateAppTracer(a.config); err != nil {
		return
	}

	if a.router, err = CreateAppRouter(a.config); err != nil {
		return
	}

	if a.storage, err = CreateAppStorage(a.config); err != nil {
		return
	}

	a.initRoutes()

	return
}

// initRoutes function
func (a *TracingApp) initRoutes() {
	a.router.router.GET("/", a.indexHandler)
	a.router.router.GET("/users", a.GetAllUsers)
	a.router.router.GET("/users/:uid", a.GetUserByUid)
}

// Run function
func (a *TracingApp) Run() (err error) {
	err = a.router.Run()
	return
}
