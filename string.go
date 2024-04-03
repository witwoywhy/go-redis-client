package gedis

import "strconv"

type Stringer interface {
	Append(key string, value any) (int, error)
	Decr(key string) (int, error)
	DecrBy(key string, value int) (int, error)
	Get(key string) (string, error)
	GetDel(key string) (string, error)
	GetEx(key string, options ...TTL) (string, error)
	GetRange(key string, start, end int) (string, error)
	Incr(key string) (int, error)
	IncrBy(key string, value int) (int, error)
	IncrByFloat(key string, value float64) (float64, error)

	Set(key string, value any) error
	MGet(keys ...string) ([]string, error)
	MSet(multiple []Multiple) error
	GetSet(key string, value any) (string, error)
}

func (g *Gedis) Append(key string, value any) (int, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(APPEND)
	w.addString(key)
	err := w.add(value)
	if err != nil {
		return 0, err
	}

	_, err = pool.conn.Write(w.toBytes())
	if err != nil {
		return 0, err
	}

	return readIntegers(<-ch)
}

func (g *Gedis) Decr(key string) (int, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(DECRBY)
	w.addString(key)

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return 0, err
	}

	return readIntegers(<-ch)
}

func (g *Gedis) DecrBy(key string, value int) (int, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(DECRBY)
	w.addString(key)
	w.add(value)

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return 0, err
	}

	return readIntegers(<-ch)
}

func (g *Gedis) get(cmd, key string) (string, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(cmd)
	w.addString(key)

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return "", err
	}

	ls, err := readStrings(<-ch)
	if err != nil {
		return "", err
	}

	if len(ls) < 1 {
		return "", nil
	}

	return ls[0], nil
}

func (g *Gedis) Get(key string) (string, error) {
	return g.get(GET, key)
}

func (g *Gedis) GetDel(key string) (string, error) {
	return g.get(GETDEL, key)
}

func (g *Gedis) GetEx(key string, options ...TTL) (string, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(GETEX)
	w.addString(key)

	if len(options) > 0 {
		option := options[0]
		w.addString(string(option.Option))
		if option.Option != PERSIST {
			w.add(option.Time)
		}
	}

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return "", err
	}

	ls, err := readStrings(<-ch)
	if err != nil {
		return "", err
	}

	if len(ls) < 1 {
		return "", nil
	}

	return ls[0], nil
}

func (g *Gedis) GetRange(key string, start int, end int) (string, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(GETRANGE)
	w.addString(key)
	err := w.add(start)
	if err != nil {
		return "", err
	}

	err = w.add(end)
	if err != nil {
		return "", err
	}

	_, err = pool.conn.Write(w.toBytes())
	if err != nil {
		return "", err
	}

	ls, err := readStrings(<-ch)
	if err != nil {
		return "", err
	}

	if len(ls) < 1 {
		return "", nil
	}

	return ls[0], nil
}

func (g *Gedis) GetSet(key string, value any) (string, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(GETSET)
	w.addString(key)
	err := w.add(value)
	if err != nil {
		return "", err
	}

	_, err = pool.conn.Write(w.toBytes())
	if err != nil {
		return "", err
	}

	ls, err := readStrings(<-ch)
	if err != nil {
		return "", err
	}

	if len(ls) < 1 {
		return "", nil
	}

	return ls[0], nil
}

func (g *Gedis) Incr(key string) (int, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(INCR)
	w.addString(key)

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return 0, err
	}

	return readIntegers(<-ch)
}

func (g *Gedis) IncrBy(key string, value int) (int, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(INCRBY)
	w.addString(key)
	err := w.add(value)
	if err != nil {
		return 0, err
	}

	_, err = pool.conn.Write(w.toBytes())
	if err != nil {
		return 0, err
	}

	return readIntegers(<-ch)
}

func (g *Gedis) IncrByFloat(key string, value float64) (float64, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(INCRBYFLOAT)
	w.addString(key)
	err := w.add(value)
	if err != nil {
		return 0, err
	}

	_, err = pool.conn.Write(w.toBytes())
	if err != nil {
		return 0, err
	}

	ls, err := readStrings(<-ch)
	if err != nil {
		return 0, err
	}

	if len(ls) < 1 {
		return 0, nil
	}

	return strconv.ParseFloat(ls[0], 64)
}

func (g *Gedis) Set(key string, value any) error {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(SET)
	w.add(key)
	err := w.add(value)
	if err != nil {
		return err
	}

	_, err = pool.conn.Write(w.toBytes())
	if err != nil {
		return err
	}

	resp := <-ch
	err = isSimpleError(resp)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gedis) MGet(keys ...string) ([]string, error) {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(MGET)
	for _, k := range keys {
		err := w.add(k)
		if err != nil {
			return nil, err
		}
	}

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return nil, err
	}

	ls, err := readStrings(<-ch)
	if err != nil {
		return nil, err
	}

	if len(ls) < 1 {
		return nil, nil
	}

	return ls, nil
}

func (g *Gedis) MSet(multiple []Multiple) error {
	pool := g.getPool()
	defer g.returnToPool(pool)

	ch := makeResponsChAndRead(pool)

	w := &writer{lists: make([]string, 0)}
	w.cmd(MSET)
	for _, m := range multiple {
		w.addString(m.Key)
		err := w.add(m.Value)
		if err != nil {
			return err
		}
	}

	_, err := pool.conn.Write(w.toBytes())
	if err != nil {
		return err
	}

	err = isSimpleError(<-ch)
	if err != nil {
		return err
	}

	return nil
}
