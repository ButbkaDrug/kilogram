package client

import(
    "path/filepath"
    "log"
    "os"
    "os/signal"
    "syscall"
    tdlib "github.com/zelenin/go-tdlib/client"
)

const (
    apiId   = 25850438
    apiHash = "b38df0acdac68ace28ea63cd70b95d67"
)

func NewClient() *tdlib.Client {

	// client authorizer
	authorizer := tdlib.ClientAuthorizer()

	go tdlib.CliInteractor(authorizer)


	authorizer.TdlibParameters <- &tdlib.SetTdlibParametersRequest{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    false,
		UseMessageDatabase:     false,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

    _, err := tdlib.SetLogVerbosityLevel(&tdlib.SetLogVerbosityLevelRequest{
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
