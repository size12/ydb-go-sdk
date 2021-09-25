package main

import (
	"context"
	"fmt"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/errors"
	"os"
	"path"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	connectParams, err := ydb.ConnectionString(os.Getenv("YDB"))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot create connect params from connection string env['YDB'] = '%s': %v\n", os.Getenv("YDB"), err)
		os.Exit(1)
	}

	options := []ydb.Option{
		ydb.WithDialTimeout(3 * time.Second),
		ydb.WithCertificatesFromFile("~/.ydb/CA.pem"),
		ydb.WithSessionPoolIdleThreshold(time.Second * 5),
		ydb.WithSessionPoolKeepAliveMinSize(-1),
		ydb.WithDiscoveryInterval(5 * time.Second),
	}
	if token, has := os.LookupEnv("YDB_ACCESS_TOKEN_CREDENTIALS"); has {
		options = append(options, ydb.WithAccessTokenCredentials(token))
	}
	if v, has := os.LookupEnv("YDB_ANONYMOUS_CREDENTIALS"); has && v == "1" {
		options = append(options, ydb.WithAnonymousCredentials())
	}

	db, err := ydb.New(
		ctx,
		connectParams,
		options...,
	)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "connect error: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = db.Close() }()

	err = db.Scheme().CleanupDatabase(ctx, connectParams.Database(), "series", "episodes", "seasons")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cleaunup database failed: %v\n", err)
		os.Exit(1)
	}

	err = db.Scheme().EnsurePathExists(ctx, connectParams.Database())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ensure path exists failed: %v\n", err)
		os.Exit(1)
	}

	err = describeTableOptions(ctx, db.Table())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "describe table options error: %v\n", err)
		os.Exit(1)
	}

	err = createTables(ctx, db.Table(), connectParams.Database())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "create tables error: %v\n", err)
		os.Exit(1)
	}

	err = describeTable(ctx, db.Table(), path.Join(
		connectParams.Database(), "series",
	))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "describe table error: %v\n", err)
		os.Exit(1)
	}

	err = fillTablesWithData(ctx, db.Table(), connectParams.Database())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fill tables with data error: %v\n", err)
		os.Exit(1)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err = selectSimple(ctx, db.Table(), connectParams.Database())
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "select simple error: %v\n", err)
			}

			err = scanQuerySelect(ctx, db.Table(), connectParams.Database())
			if err != nil {
				if !errors.IsTransportError(err, errors.TransportErrorUnimplemented) {
					_, _ = fmt.Fprintf(os.Stderr, "scan query select error: %v\n", err)
				}
			}

			err = readTable(ctx, db.Table(), path.Join(
				connectParams.Database(), "series",
			))
			if err != nil {
				fmt.Printf("read table error: %v\n", err)
			}
		}()
	}
	wg.Wait()

	for e, s := range db.Stats() {
		fmt.Printf("endpoint '%v' stats: %v\n", e, s)
	}

	whoAmI, err := db.Discovery().WhoAmI(ctx)
	fmt.Printf("whoAmI: %v ,%v\n", whoAmI, err)
}
