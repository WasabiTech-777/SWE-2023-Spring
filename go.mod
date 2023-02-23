module mymodule

go 1.19

require (
	github.com/gorilla/mux v1.8.0 // direct
	github.com/rs/cors v1.8.3 // direct
	gorm.io/gorm v1.24.5
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1
	gorm.io/driver/postgres v1.4.8
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)
require "github.com/WasabiTech-777/SWE-2023-Spring/initialize" v0.0.0
replace "github.com/WasabiTech-777/SWE-2023-Spring/initialize" v0.0.0 => "./src/server/initialize"
require "github.com/WasabiTech-777/SWE-2023-Spring/routes" v0.0.0
replace "github.com/WasabiTech-777/SWE-2023-Spring/routes" v0.0.0 => "./src/server/routes"
require "github.com/WasabiTech-777/SWE-2023-Spring/models" v0.0.0
replace "github.com/WasabiTech-777/SWE-2023-Spring/models" v0.0.0 => "./src/server/models"