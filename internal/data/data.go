package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"why-sys/internal/conf"
	"why-sys/internal/data/ent"
	"why-sys/internal/data/ent/migrate"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAuthRepo, NewRoleRepo, NewMenuRepo, NewPermissionRepo)

var rbacModel = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
`

// Data .
type Data struct {
	db *ent.Client

	casbinAdapter *casbin.SyncedCachedEnforcer
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to postgresql: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true), migrate.WithForeignKeys(false)); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	model, err := model.NewModelFromString(rbacModel)
	if err != nil {
		log.Errorf("字符串%s加载模型失败!", rbacModel)
		return nil, nil, err
	}
	adapter, err := entadapter.NewAdapter(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("casbin adapter 创建失败!")
		return nil, nil, err
	}
	syncedCachedEnforcer, _ := casbin.NewSyncedCachedEnforcer(model, adapter)
	syncedCachedEnforcer.SetExpireTime(60 * 60)

	return &Data{
			db:            client,
			casbinAdapter: syncedCachedEnforcer,
		}, func() {
			log.Info("message", "closing the data resources")
			if err := client.Close(); err != nil {
				log.Error(err)
			}
		}, nil
}
