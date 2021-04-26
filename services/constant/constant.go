package constant

import "fmt"

const (
	DateFormat          = "2006-01-02"
	TimeFormatShort     = "15:04"
	TimeFormatLong      = TimeFormatShort + ":05"
	DateTimeFormatShort = DateFormat + " " + TimeFormatShort
	DateTimeFormatLong  = DateFormat + " " + TimeFormatLong
)

const (
	AllowedHeaders     = "Authorization;Content-Type"
	HeaderCacheControl = "Cache-Control"
	CacheNoCache       = "no-cache"
	CacheNoStore       = "no-store"
	HeaderUsername     = "X-Username"
)

const (
	MongoConnect          = "mongodb: connecting"
	MongoDisconnect       = "mongodb: disconnecting"
	MongoFailedConnect    = "mongodb: failed to connect"
	MongoNoResponse       = "mongodb: no response received"
	MongoGenericError     = "mongodb: unknown error"
	PostgreConnect        = "postgresql: connecting"
	PostgreDisconnect     = "postgresql: disconnecting"
	PostgreFailedConnect  = "postgresql: failed to connect"
	PostgreNoResponse     = "postgresql: no response received"
	PostgreInitializeDb   = "postgresql: initialize db, auto migration"
	PostgreGenericError   = "postgresql: unknown error"
	RedisConnect          = "redis: connecting"
	RedisDisconnect       = "redis: disconnecting"
	RedisFailedConnect    = "redis: failed to connect"
	RedisNoResponse       = "redis: no response received"
	RedisFailedToFlush    = "redis: failed to flush"
	RedisSetDataFailed    = "redis: failed to set data"
	RedisGetDataFailed    = "redis: failed to get data"
	RedisDeleteDataFailed = "redis: failed to delete data"
	RedisGenericError     = "redis: unknown error"
	LogOpenFileFailed     = "log: failed to open file"
	LogClose              = "log: closing log file"
	UpdateSwaggerHost     = "swagger: updating host"
	ServeSwagger          = "swagger: serving"
	GinInitialize         = "gin: initialize"
	GinInitializeCors     = "gin: initialize cors"
	GinInitializeRouter   = "gin: initialize router"
	GinListenServer       = "gin: listen and serve"
)

const (
	ModelMigration = "model: auto migrate all models"
	ModelInit      = "model: initialize model with default values"
)

const (
	SuccesCreateOrder         = "Success create order"
	ApiSuccess                = "SUCCESS"
	ApiError                  = "ERROR"
	MaximalAddress            = "Jumlah address ada sudah maximal, silahkan hapus salah satu dan masukan address baru"
	ApiUnknownError           = "UNKNOWN ERROR"
	ApiRecordNotFound         = "Record Not Found"
	ApiInvalidParameter       = "parameter is not valid"
	ApiInvalidHeader          = "http header is not valid"
	ApiInvalidJson            = "json not in correct format"
	ApiAlreadyExists          = "object already exists"
	ApiObjectRequired         = "object required"
	ApiQueryFailed            = "query failed"
	ApiCreateFailed           = "create object failed"
	ApiUpdateFailed           = "update object failed"
	ApiDeleteFailed           = "delete object failed"
	ApiObjectNotFound         = "object not found"
	TotalMaximalAddress int64 = 5
	TotalDistance       int64 = 3000
	ApiSellerNotFOUND         = "Opps tidak ada seller di area anda"
	LimitSearchNear           = 35
	OrderStatusWaiting        = "waiting"
	OrderStatusDone           = "done"
	OrderStatusOtw            = "otw"
	OrderStatusProcess        = "process"
	ErrorGetData              = "Erro get data from DB !"
)

func ComposeMessage(msg, detail string) string {
	if detail == "" {
		return fmt.Sprintf(msg)
	} else {
		return fmt.Sprintf(msg+", detail: '%s'", detail)
	}
}
