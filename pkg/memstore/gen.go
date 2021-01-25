package memstore

//go:generate go run ./generate/gen.go -fileOut="generated--userstore.go" -name=User
//go:generate go run ./generate/gen.go -fileOut="generated--messagestore.go" -name=Message
//go:generate go run ./generate/gen.go -fileOut="generated--nodestore.go" -name=Node
