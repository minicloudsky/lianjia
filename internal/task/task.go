package task

import (
	"context"
	"github.com/go-co-op/gocron"
	redislock "github.com/go-co-op/gocron-redis-lock"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minicloudsky/lianjia/internal/biz"
	"github.com/minicloudsky/lianjia/internal/conf"
	"github.com/minicloudsky/lianjia/internal/service"
	"github.com/minicloudsky/lianjia/pkg/system_stat"
	"github.com/redis/go-redis/v9"
	"time"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewTaskScheduler)

type TaskScheduler struct {
	confData      *conf.Data
	repo          biz.LianjiaRepo
	log           *log.Helper
	redisClient   *redis.Client
	svc           *service.Service
	cronScheduler *gocron.Scheduler
	uc            *biz.LianjiaUsecase
}

func NewTaskScheduler(confData *conf.Data, uc *biz.LianjiaUsecase, svc *service.Service, redisClient *redis.Client,
	repo biz.LianjiaRepo, logger log.Logger) *TaskScheduler {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.NewHelper(logger).Errorf("Error loading location:", err)
		return nil
	}
	taskScheduler := &TaskScheduler{
		confData:      confData,
		log:           log.NewHelper(log.With(logger, "module", "task")),
		svc:           svc,
		cronScheduler: gocron.NewScheduler(loc),
		redisClient:   redisClient,
		uc:            uc,
		repo:          repo,
	}
	locker, err := redislock.NewRedisLocker(redisClient, redislock.WithTries(1))
	if err != nil {
		return nil
	}
	taskScheduler.cronScheduler.SingletonModeAll()
	taskScheduler.cronScheduler.TagsUnique()
	taskScheduler.cronScheduler.WithDistributedLocker(locker)
	taskScheduler.cronScheduler.StartAsync()
	err = taskScheduler.StartAllTask()
	if err != nil {
		return nil
	}

	return taskScheduler
}

func (taskScheduler *TaskScheduler) StartAllTask() error {
	info := system_stat.StatCpuMemory()
	taskScheduler.log.Infof("---current system info--- CPU: %d cores, Memory: %d MB",
		info.NumCpu, info.Memory)
	for _, task := range taskScheduler.confData.Tasks {
		if task.Enable {
			taskFunc, ok := FuncMap[task.Name]
			if ok {
				ctx := context.Background()
				job, err := taskScheduler.cronScheduler.Cron(task.CronExpr).Tag(task.Name).Do(taskFunc, ctx, taskScheduler)
				taskScheduler.log.Infof("%s next run time: %v", job.GetName(), job.NextRun().Format(time.DateTime))
				if err != nil {
					taskScheduler.log.Errorf("fail to do job %v err: %v", job, err)
					return err
				}
			}
		}
	}
	taskScheduler.log.Infof("total jobs %d", len(taskScheduler.cronScheduler.Jobs()))
	return nil
}

var FuncMap = map[string]func(ctx context.Context, ts *TaskScheduler){
	"fetch_lianjia_house": FetchLianjiaHouse,
}

func (taskScheduler *TaskScheduler) Start(ctx context.Context) error {
	return nil
}

func (taskScheduler *TaskScheduler) Stop(ctx context.Context) error {
	taskScheduler.cronScheduler.Stop()
	return nil
}

func FetchLianjiaHouse(ctx context.Context, ts *TaskScheduler) {

	if !ts.uc.IsErshouFangTaskFinish(ctx) &&
		!ts.uc.IsLoupanTaskFinish(ctx) &&
		!ts.uc.IsCommercialTaskFinish(ctx) {
		ts.log.Infof("---task is running,passed")
		return
	}
	ts.uc.FetchCityList(ctx)
	ts.uc.InitLianjiaChan(ctx)
	go func() {
		err := ts.uc.Receive(ctx)
		if err != nil {
			ts.log.Errorf("Receive Message err! err: %v", err)
		}
	}()
	go ts.uc.FetchErShouFang(ctx)
	go ts.uc.FetchLoupan(ctx)
	//go ts.uc.FetchZufang(ctx)
	go ts.uc.FetchCommercial(ctx)

}
