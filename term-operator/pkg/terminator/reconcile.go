package terminator

import (
	"github.com/knabben/terminator/term-operator/pkg/apis/app/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

// Reconcile new terminator CR
func Reconcile(term *v1alpha1.Terminator, event sdk.Event) error {
	if event.Deleted {
		return nil
	}

	memcacheDep := deploymentForMemcached(term)
	redisDep := deploymentForRedis(term)
	rabbitDep := deploymentForRabbit(term)
	elasticDep := deploymentForElastic(term)

	var memcacheRep int32
	if term.Spec.Memcache {
		memcacheRep = 1
	}

	var redisRep int32
	if term.Spec.Redis {
		redisRep = 1
	}

	var rabbitRep int32
	if term.Spec.Rabbitmq {
		rabbitRep = 1
	}

	var elasticRep int32
	if term.Spec.Elastic {
		elasticRep = 1
	}

	go setReplica(memcacheDep, memcacheRep)
	go setReplica(redisDep, redisRep)
	go setReplica(rabbitDep, rabbitRep)
	go setReplica(elasticDep, elasticRep)

	return nil
}
