Чтобы запустить, обновите путь импорта в файле main.go

Было: 
myStack "projects/bracket-balance/myStack"

Должно стать: 
myStack "ВашаПапка/bracket-balance/myStack"


После чего:

go mod tidy 
go run main.go