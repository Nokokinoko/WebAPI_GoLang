package handler

type DevelopController struct {
	Controller
}

func (ctrl DevelopController) urlBase(prm JsonParameter) string {
	return urlStaging
}

func (ctrl DevelopController) fqgn() string {
	return "com.example_studio.example_qgn"
}

func (ctrl DevelopController) identification() string {
	return "dev"
}
