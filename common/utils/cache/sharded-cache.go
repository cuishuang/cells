// +build !arm

/*
 * Copyright (c) 2019-2021. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

// Package cache provides ready-to-use in-memory cache mechanisms
package cache

import (
"fmt"
"os"
"strconv"
"sync"
"time"

bigcache "github.com/allegro/bigcache/v3"
"github.com/pydio/cells/v4/common/service/metrics"
)

var (
	defaultConfig bigcache.Config
	o             = sync.Once{}
)

// NewSharded creates a cache instance instrumented with a regular report of statistics
func NewSharded(identifier string, opts ...Option) Cache {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	bcConf := DefaultBigCacheConfig()
	if opt.EvictionTime > 0 {
		bcConf.LifeWindow = opt.EvictionTime
	}
	if opt.CleanWindow > 0 {
		bcConf.CleanWindow = opt.CleanWindow
	}
	return newInstrumentedCache(identifier, bcConf)
}

// newInstrumentedCache creates a BigCache instance with a regular report of statistics
func newInstrumentedCache(serviceName string, cacheConfig ...bigcache.Config) *InstrumentedCache {
	scope := metrics.GetMetricsForService(serviceName)
	conf := DefaultBigCacheConfig()
	if len(cacheConfig) > 0 {
		conf = cacheConfig[0]
	}
	c, _ := bigcache.NewBigCache(conf)
	i := &InstrumentedCache{}
	i.Cache = &bigCache{c}
	i.scope = scope
	i.ticker = time.NewTicker(30 * time.Second)
	go func() {
		for range i.ticker.C {
			s := c.Stats()
			i.scope.Gauge("bigcache_capacity").Update(float64(c.Capacity()))
			i.scope.Gauge("bigcache_hits").Update(float64(s.Hits))
			i.scope.Gauge("bigcache_collisions").Update(float64(s.Collisions))
			i.scope.Gauge("bigcache_delHits").Update(float64(s.DelHits))
			i.scope.Gauge("bigcache_delMisses").Update(float64(s.DelMisses))
			i.scope.Gauge("bigcache_misses").Update(float64(s.Misses))
		}
	}()
	return i
}

func DefaultBigCacheConfig() bigcache.Config {
	o.Do(func() {
		defaultConfig = bigcache.DefaultConfig(30 * time.Minute)
		defaultConfig.Shards = 64
		defaultConfig.MaxEntriesInWindow = 10 * 60 * 64
		defaultConfig.MaxEntrySize = 200
		defaultConfig.HardMaxCacheSize = 8
		if limit := os.Getenv("CELLS_CACHES_HARD_LIMIT"); limit != "" {
			if l, e := strconv.ParseInt(limit, 10, 64); e == nil {
				if l < 8 {
					fmt.Println("[ENV] ## WARNING ## CELLS_CACHES_HARD_LIMIT cannot use a value lower than 8 (MB).")
				} else {
					defaultConfig.HardMaxCacheSize = int(l)
				}
			}
		}
		// Disabled as cleanUp may seem to be able to lock the cache
		// c.CleanWindow = 10 * time.Minute
	})
	return defaultConfig
}

