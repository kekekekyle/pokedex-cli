package pokecache

import (
  "time"
  "sync"
)

type cacheEntry struct {
  createdAt time.Time
  val       []byte
}

type Cache struct {
  entry    map[string]cacheEntry
  mu       *sync.RWMutex
  interval time.Duration
  ticker   *time.Ticker
}

func (c *Cache) Add (key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()
  c.entry[key] = cacheEntry{
    createdAt: time.Now(),
    val: val,
  }
}

func (c *Cache) Get (key string) ([]byte, bool) {
  c.mu.RLock()
  defer c.mu.RUnlock()
  if entry, ok := c.entry[key]; ok {
    return entry.val, true
  }
  return nil, false
}

func (c *Cache) reapLoop () {
  c.mu.Lock()
  defer c.mu.Unlock()

  for key, val := range c.entry {
    if time.Now().Sub(val.createdAt) > c.interval {
      delete(c.entry, key)
    }
  }
}

func (c *Cache) Stop () {
  c.ticker.Stop()
}

func NewCache (interval time.Duration) *Cache {
  ticker := time.NewTicker(interval)

  cache := &Cache{
    entry: map[string]cacheEntry{},
    mu: &sync.RWMutex{},
    interval: interval,
    ticker: ticker,
  }

  go func () {
    for {
      <- ticker.C
      cache.reapLoop()
    }
  }()

  return cache
}
