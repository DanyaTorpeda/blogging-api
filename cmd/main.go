package main

import (
	blogging "blogging_app"
	"blogging_app/pkg/handler"
	"blogging_app/pkg/repository"
	"blogging_app/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/*
	1. создание структуры проекта
	2. объявление структуры сервера
	3. объявление сущностей
	4. имплементация хэндлера(типа) и инициализация эндпоинтов
	5. реализация обработчиков для работы хэндлера
	6. объявления слоя сервисов и перечисление интерфейсов для реализации
	7. объявление слоя хранилища и перечисления интерфейсов для реализации
	8. добавить файл конфигурации
	9. инициализация бд
	10. создания файлов миграций
	11. подключение бд
	12. добавить переменные окружения
	13. добавление логирования
	14. реализации регистрации
	15. функция для обработки ошибок
	16. реализация слоя сервисов
	17. реализация слоя репозитория
	18. реализация авторизации
	19. создание middleware
	20. реализация методов эндпоинтов для объектов
	21. если нужно то создать функцию привидение типов с контекста (типа интерфейс) к нужному целевому типу
	22. для возвращения всех использовать дополнительную структуру
	23. graceful shutdown
*/

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error occured reading environment variables: %s", err.Error())
	}

	if err := InitConfigs(); err != nil {
		logrus.Errorf("error occured setting configuration variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Errorf("error occured connecting to db: %s", err.Error())
	}
	if db == nil {
		logrus.Errorf("db is nil in main func")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(blogging.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Errorf("error occured server didn't start working: %s", err.Error())
	}
}

func InitConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
