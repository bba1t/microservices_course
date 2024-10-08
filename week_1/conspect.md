### week_1 - базовая настройка обоих серваков 

#### http
1. `cmd/http_server` - получение данных с тела `http-запроса`, `Decode/Encode`, `request параметры`, `Lock()`
2. `cmd/http_client` - `HTTP POST-запрос`, `HTTP GET-запрос`, `Marshal`

#### grpc
1. В папку `api` помещаю и заполняю декларацию protobuf. Для работы с proto-файлами устанавливаю protoc: `brew install protobuf`
2. После того как декларация готова, из нее можно сгенерировать код на любом языке программирования. Для этого в makefile 
прописываю цели `install-deps` и `generate-note-api`
3. `cmd/grpc_server` - имплементирую метод `Get()` 
4. `cmd/grpc_client` - 