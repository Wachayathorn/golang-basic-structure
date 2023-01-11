package utils

type UtilsInterface interface {
	DoRequest(url, httpMethod, path string, data interface{}) ([]byte, error)
}

type Utils struct {
}

func Init() UtilsInterface {
	return &Utils{}
}
