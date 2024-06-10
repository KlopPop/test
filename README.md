# Сервис товаров для марктеплейса

для старта подключаем бд
psql -h localhost -U postgres -w -c "create database postgres;"

накатываем миграции
 из файла db/migrations/{version}_create_tables.up.sql

ставим модули
  go get github.com/BurntSushi/toml
  go get github.com/go-chi
  go get github.com/joho/godotenv
  go get github.com/kelseyhightower/envconfig
  go get github.com/lib/pq
  go get gopkg.in/yaml.v3
  go get olympos.io/encoding/edn