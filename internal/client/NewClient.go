package client

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
    "strconv"

	"github.com/joho/godotenv"
	tdlib "github.com/zelenin/go-tdlib/client"
)

func NewClient() *tdlib.Client {

    err := godotenv.Load()
    if err != nil{
        log.Fatal(err)
    }
    var apiHash = os.Getenv("KILOGRAM_API_HASH")

    id, err := strconv.Atoi(os.Getenv("KILOGRAM_API_ID"))

    if err != nil {
        log.Fatal(err)
    }

    var apiId = int32(id)

	// client authorizer
	authorizer := tdlib.ClientAuthorizer()

	go tdlib.CliInteractor(authorizer)


	authorizer.TdlibParameters <- &tdlib.SetTdlibParametersRequest{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        false,
		UseChatInfoDatabase:    false,
		UseMessageDatabase:     false,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: false,
		IgnoreFileNames:        false,
	}

    _, err = tdlib.SetLogVerbosityLevel(&tdlib.SetLogVerbosityLevelRequest{
        NewVerbosityLevel: 0,
    })

    if err != nil {
        log.Fatal(err)
    }


    c, err := tdlib.NewClient(authorizer)

	if err != nil {
        log.Fatal("Could not start client!")
	}

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		c.Stop()
		os.Exit(1)
	}()

    return c
}
