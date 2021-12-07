package app

// TraceApplication struct
type TraceApplication struct {
	router *AppRouter
	logger *AppLogger
}

// Create function
func Create() (a *TraceApplication, err error) {
	a = new(TraceApplication)

	// if a.logger, err = CreateLogger(); err != nil {
	// 	return
	// }

	if a.router, err = CreateRouter(); err != nil {
		return
	}

	return
}

func (a *TraceApplication) Run() (err error) {
	err = a.router.Run()
	return
}
