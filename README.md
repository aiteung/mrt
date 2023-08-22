# Module Handling Iteung
Iteung module routing

Usage
```go
Modulename,Pesan:=IteungModuleCall(Info *types.MessageInfo, Message *waProto.Message, waclient *whatsmeow.Client, MongoConn *mongo.Database, TypoCollection string, ModuleCollection string)
if Modulename != ""{
    Caller(Modulename,Pesan)
}
```

List of Mongo Collection :

1. Typo : to replace typo word
2. Module : return module name from keyword 

## Dev
```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v0.0.1                          #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/aiteung/module@v0.0.1   #publish to pkg dev, replace ORG/URL with your repo URL
```
