module db

go 1.12

replace Utils/wlog => ../wlog

require (
	Utils/wlog v0.0.0-00010101000000-000000000000
	github.com/BurntSushi/toml v0.3.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/square/go-jose.v1 v1.1.2 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
