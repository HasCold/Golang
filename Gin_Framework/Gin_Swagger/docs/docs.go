package docs

// Add these gin-swagger packages
// go get -u github.com/swaggo/swag/cmd/swag
// go get -u github.com/swaggo/gin-swagger
// go get -u github.com/swaggo/files

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// Go templating

var doc = `{

}`
