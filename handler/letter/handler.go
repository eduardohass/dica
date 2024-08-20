package letter

import (
	"fmt"

	"github.com/eduardohass/dica/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	fmt.Println("DBG==Inicializando o handler letter")
	logger = config.GetLogger("handler")
	fmt.Println("DBG==logger: ", logger)
	db = config.GetDatabase()
	fmt.Println("DBG==db: ", db)
}