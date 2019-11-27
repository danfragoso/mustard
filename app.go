package mustard

//NewApp - Creates and returns a new mustard app
func NewApp(name string) *App {
	app := App{
		Name: name,
	}

	return &app
}
