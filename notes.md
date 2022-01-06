OBS: NÃO MEXER EM:
Dentro da pasta pb
`user_grpc.pb.go`
`user.pb.go`
pois esses arquivos foram gerados automaticamente com esse comando
`protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb`

Comando pra rodar gRPC client evans
`evans -r repl --host localhost --port 50051`
