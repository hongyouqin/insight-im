package main

import (
	"context"
	msggate "insight/internal/msg-gate"
	"insight/pkg/common/config"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator"
	"github.com/natefinch/lumberjack"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	fx.New(
		fx.Provide(newLogger),
		fx.Provide(newConfig),
		fx.Provide(msggate.NewWsServer),
		fx.Provide(newValidator),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			//optional 使得fx框架里的日志输出到指定的logger
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(Server),
	).Run()
}

func Server(lc fx.Lifecycle, log *zap.Logger, cfg *config.GateConfig, wsSvr *msggate.WsServer, validate *validator.Validate) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					//启动websocket
					go func() {
						wsSvr.StartWs()
					}()
					//启动rpc
					startRpc(log)
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				log.Info("server exiting")
				return nil
			},
		})
}

func startRpc(log *zap.Logger) {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(time.Second * 60),
		MaxConnectionAgeGrace: time.Duration(time.Second * 20),
		Time:                  time.Duration(time.Second * 60),
		Timeout:               time.Duration(time.Second * 60),
		MaxConnectionAge:      time.Duration(time.Hour * 2),
	})
	server := grpc.NewServer(keepParams)
	defer server.GracefulStop()
	listen, err := net.Listen("tcp", ":7748")
	if err != nil {
		panic("listening err:" + err.Error())
	}
	defer listen.Close()
	log.Info("msg-gate rpc listen success", zap.String("address", ":7748"))
	err = server.Serve(listen)
	if err != nil {
		log.Error("rpc listening err", zap.String("err", err.Error()))
	}
}

func newConfig() *config.GateConfig {
	var cfg config.GateConfig
	if _, err := toml.DecodeFile("../../configs/msg-gate/msg-gate.toml", &cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func newLogger() (*zap.Logger, error) {
	//return zap.NewProduction()
	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//文件writeSyncerí
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/msg-gate.log", //日志文件存放目录
		MaxSize:    10,                    //文件大小限制,单位MB
		MaxBackups: 20,                    //最大保留日志文件数量
		MaxAge:     30,                    //日志文件保留天数
		Compress:   false,                 //是否压缩处理
	})
	fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	logger := zap.New(fileCore, zap.AddCaller()) //AddCaller()为显示文件名和行号
	return logger, nil
}

func newValidator() *validator.Validate {
	validate := validator.New()
	return validate
}
