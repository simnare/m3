	"github.com/m3db/m3/src/dbnode/x/tchannel"
	xconfig "github.com/m3db/m3x/config"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"
	"github.com/m3db/m3x/instrument"
	xlog "github.com/m3db/m3x/log"
	"github.com/m3db/m3x/pool"
	xsync "github.com/m3db/m3x/sync"
		logger.Fatalf("could not parse new file mode: %v", err)
		logger.Fatalf("could not parse new directory mode: %v", err)
		logger.Fatalf("could not acquire lock on %s: %v", lockPath, err)
		logger.Fatalf("could not connect to metrics: %v", err)
		logger.Fatalf("could not resolve local host ID: %v", err)
				logger.Fatalf("unable to create etcd clusters: %v", err)
			logger.Infof("using seed nodes etcd cluster: zone=%s, endpoints=%v", zone, endpoints)
		logger.WithFields(
			xlog.NewField("hostID", hostID),
			xlog.NewField("seedNodeHostIDs", fmt.Sprintf("%v", seedNodeHostIDs)),
		).Info("resolving seed node configuration")
				logger.Fatalf("unable to create etcd config: %v", err)
				logger.Fatalf("could not start embedded etcd: %v", err)
		SetMetricsSamplingRate(cfg.Metrics.SampleRate())
		logger.Warnf("max index query IDs concurrency was not set, falling back to default value")
		logger.Fatalf("unable to start build reporter: %v", err)
		logger.Fatalf("could not construct query cache: %s", err.Error())
		logger.Fatalf("could not set initial runtime options: %v", err)
			logger.Fatalf("could not determine if host supports HugeTLB: %v", err)
			logger.Warnf("host doesn't support HugeTLB, proceeding without it")
		logger.Fatalf("unknown commit log queue size type: %v",
			cfg.CommitLog.Queue.CalculationType)
			logger.Fatalf("unknown commit log queue channel size type: %v",
				cfg.CommitLog.Queue.CalculationType)
	// Set the series cache policy
	seriesCachePolicy := cfg.Cache.SeriesConfiguration().Policy
	opts = opts.SetSeriesCachePolicy(seriesCachePolicy)

	// Apply pooling options
	opts = withEncodingAndPoolingOptions(cfg, logger, opts, cfg.PoolingPolicy)

		logger.Fatalf("could not create persist manager: %v", err)
			logger.Fatalf("could not initialize dynamic config: %v", err)
			logger.Fatalf("could not initialize static config: %v", err)
		logger.Fatalf("could not initialize m3db topology: %v", err)
		})
		logger.Fatalf("could not create m3db client: %v", err)
		logger.Fatalf("could not create bootstrap process: %v", err)
				logger.Errorf("updated bootstrapper list is empty")
				logger.Errorf("updated bootstrapper list failed: %v", err)
		logger.Fatalf("could not create cluster topology watch: %v", err)
		logger.Fatalf("could not construct database: %v", err)
		logger.Fatalf("could not open database: %v", err)
		logger.Fatalf("could not open tchannelthrift interface on %s: %v",
			cfg.ListenAddress, err)
	logger.Infof("node tchannelthrift: listening on %v", cfg.ListenAddress)
		logger.Fatalf("could not open tchannelthrift interface on %s: %v",
			cfg.ClusterListenAddress, err)
	logger.Infof("cluster tchannelthrift: listening on %v", cfg.ClusterListenAddress)
		logger.Fatalf("could not open httpjson interface on %s: %v",
			cfg.HTTPNodeListenAddress, err)
	logger.Infof("node httpjson: listening on %v", cfg.HTTPNodeListenAddress)
		logger.Fatalf("could not open httpjson interface on %s: %v",
			cfg.HTTPClusterListenAddress, err)
	logger.Infof("cluster httpjson: listening on %v", cfg.HTTPClusterListenAddress)
				logger.Errorf("debug server could not listen on %s: %v", cfg.DebugListenAddress, err)
			logger.Fatalf("could not bootstrap database: %v", err)
		logger.Infof("bootstrapped")
	logger.Warnf("interrupt: %v", interruptErr)
			logger.Errorf("close database error: %v", err)
		logger.Infof("server closed")
		logger.Errorf("server closed after %s timeout", closeTimeout.String())
func bgValidateProcessLimits(logger xlog.Logger) {
		logger.Warnf(`cannot validate process limits: invalid configuration found [%v]`, message)
		logger.WithFields(
			xlog.NewField("url", xdocs.Path("operational_guide/kernel_configuration")),
		).Warnf(`invalid configuration found [%v], refer to linked documentation for more information`, err)
	logger xlog.Logger,
			logger.Warnf("error resolving cluster new series insert limit: %v", err)
		logger.Warnf("unable to set cluster new series insert limit: %v", err)
		logger.Errorf("could not watch cluster new series insert limit: %v", err)
					logger.Warnf("unable to parse new cluster new series insert limit: %v", err)
				logger.Warnf("unable to set cluster new series insert limit: %v", err)
	logger xlog.Logger,
	logger xlog.Logger,
		logger.Errorf("could not resolve KV key %s: %v", key, err)
			logger.Errorf("could not unmarshal KV key %s: %v", key, err)
			logger.Errorf("could not process value of KV key %s: %v", key, err)
			logger.Infof("set KV key %s: %v", key, protoValue.Value)
		logger.Errorf("could not watch KV key %s: %v", key, err)
					logger.Warnf("could not set default for KV key %s: %v", key, err)
				logger.Warnf("could not unmarshal KV key %s: %v", key, err)
				logger.Warnf("could not process change for KV key %s: %v", key, err)
			logger.Infof("set KV key %s: %v", key, protoValue.Value)
	logger xlog.Logger,
		logger.Fatalf("could not watch value for key with KV: %s",
			kvconfig.BootstrapperKey)
				logger.WithFields(
					xlog.NewField("key", kvconfig.BootstrapperKey),
					xlog.NewErrField(err),
				).Error("error converting KV update to string array")
	logger xlog.Logger,
		logger.Infof("bytes pool registering bucket capacity=%d, size=%d, "+
			bucket.RefillLowWaterMark, bucket.RefillHighWaterMark)
		logger.Fatalf("unrecognized pooling type: %s", policy.Type)
	logger.Infof("bytes pool %s init", policy.Type)
		SetWriteBatchPool(writeBatchPool)
	resultsPool := index.NewResultsPool(
		poolOptions(policy.IndexResultsPool, scope.SubScope("index-results-pool")))
		SetResultsPool(resultsPool)
	resultsPool.Init(func() index.Results {
		return index.NewResults(nil, index.ResultsOptions{}, indexOpts)