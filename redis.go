package gedis

import (
	"sync"
)

type Rediser interface {
	Stringer
}

type Gedis struct {
	pools map[int]pool
	mu    sync.Mutex
}

func NewGedis(config *Config) (Rediser, error) {
	if config.ConnectionPool <= 0 {
		config.ConnectionPool = DefaultConnectionPool
	}

	gedis := &Gedis{
		pools: make(map[int]pool),
		mu:    sync.Mutex{},
	}
	for i := 0; i < config.ConnectionPool; i++ {
		conn, err := newConn(config)
		if err != nil {
			return nil, err
		}

		if config.Password != "" {
			err := authen(config.Password, conn)
			if err != nil {
				return nil, err
			}
		}

		err = selectDB(config.DB, conn)
		if err != nil {
			return nil, err
		}

		gedis.pools[i] = pool{
			id:   i,
			conn: conn,
		}
	}

	return gedis, nil
}

func (g *Gedis) getPool() *pool {
	var p *pool
	var key int
	for {
		g.mu.Lock()
		for k, v := range g.pools {
			p = &v
			key = k
		}

		if p != nil {
			delete(g.pools, key)
			g.mu.Unlock()
			return p
		}

		g.mu.Unlock()
	}
}

func (g *Gedis) returnToPool(pool *pool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.pools[pool.id] = *pool
}
