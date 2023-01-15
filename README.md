# golang-basic-structure

.
├── Flow.drawio
├── README.md
├── go.mod
├── go.sum
├── main.go
├── pkg
│   ├── build
│   │   └── docker-compose.yaml
│   ├── client
│   │   ├── add-book.go
│   │   ├── get-books.go
│   │   ├── init.go
│   │   └── mock_ClientAPI.go
│   ├── config
│   │   ├── load.go
│   │   └── resource
│   │       ├── config.dev.yml
│   │       └── config.local.yml
│   ├── db
│   │   ├── mongo
│   │   ├── postgresql
│   │   │   ├── init.go
│   │   │   └── script
│   │   │       └── create-table.sql
│   │   └── redis
│   ├── handler
│   │   ├── book
│   │   │   ├── add-book.go
│   │   │   ├── get-books.go
│   │   │   ├── httpinfo
│   │   │   │   └── data.go
│   │   │   └── init.go
│   │   ├── init.go
│   │   └── middleware
│   │       └── middleware.go
│   ├── model
│   │   └── book.go
│   ├── repository
│   │   ├── book
│   │   │   ├── add-book.go
│   │   │   ├── get-books.go
│   │   │   ├── init.go
│   │   │   ├── mock_BookRepository.go
│   │   │   └── repo.go
│   │   └── init.go
│   ├── service
│   │   ├── book
│   │   │   ├── add-book.go
│   │   │   ├── get-books.go
│   │   │   ├── init.go
│   │   │   └── mock_BookInterfaces.go
│   │   └── init.go
│   └── utils
│       ├── do-request.go
│       ├── init.go
│       └── mock_UtilsInterface.go
└── project-structure

21 directories, 37 files
