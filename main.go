/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"net/url"
	"os"
	"path/filepath"
	"time"

	utilsmetadata "github.com/Aryaman6492/utils-k8s-go/armometadata"
	"github.com/go-logr/zapr"
	"github.com/kubescape/go-logger"
	"github.com/kubescape/go-logger/helpers"
	"github.com/Aryaman6492/storage/pkg/cleanup"
	"github.com/Aryaman6492/storage/pkg/cmd/server"
	"github.com/Aryaman6492/storage/pkg/registry/file"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"
	"k8s.io/klog/v2"
)

func main() {
	flag.Parse()

	if logger, err := zap.NewProduction(); err == nil {
		logger = logger.WithOptions(zap.IncreaseLevel(zap.FatalLevel))
		klog.SetLogger(zapr.NewLogger(logger))
	}

	ctx := context.Background()
	clusterData, err := utilsmetadata.LoadConfig("/etc/config/clusterData.json")
	if err != nil {
		logger.L().Ctx(ctx).Fatal("load config error", helpers.Error(err))
	}
	// to enable otel, set OTEL_COLLECTOR_SVC=otel-collector:4317
	if otelHost, present := os.LookupEnv("OTEL_COLLECTOR_SVC"); present {
		ctx = logger.InitOtel("storage",
			os.Getenv("RELEASE"),
			os.Getenv("ACCOUNT_ID"),
			clusterData.ClusterName,
			url.URL{Host: otelHost})
		defer logger.ShutdownOtel(ctx)
	}

	// setup storage components
	osFs := afero.NewOsFs()
	pool := file.NewPool(filepath.Join(file.DefaultStorageRoot, "metadata.sq3"), 0) // If less than 1, a reasonable default is used.

	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewWardleServerOptions(os.Stdout, os.Stderr, osFs, pool, clusterData.Namespace)
	cmd := server.NewCommandStartWardleServer(options, stopCh)

	// cleanup task
	client, disco, err := cleanup.NewKubernetesClient()
	kubernetesAPI := cleanup.NewKubernetesAPI(client, disco)
	if err != nil {
		panic(err.Error())
	}
	interval := os.Getenv("CLEANUP_INTERVAL")
	intervalDuration, err := time.ParseDuration(interval)
	if err != nil {
		intervalDuration = time.Hour * 24
		logger.L().Info("failed to parse cleanup interval, falling back to default", helpers.Error(err), helpers.String("interval", intervalDuration.String()))
	}

	relevancyEnabled := clusterData.RelevantImageVulnerabilitiesEnabled != nil && *clusterData.RelevantImageVulnerabilitiesEnabled

	cleanupHandler := cleanup.NewResourcesCleanupHandler(osFs, file.DefaultStorageRoot, pool, intervalDuration, kubernetesAPI, relevancyEnabled)
	go cleanupHandler.StartCleanupTask(ctx)

	logger.L().Info("APIServer started")
	code := cli.Run(cmd)
	os.Exit(code)
}
