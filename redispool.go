package main

import  "menteslibres.net/gosexy/redis"

type redisPool struct {
	connections chan *redis.Client
	connFn      func() (*redis.Client, error) // function to create new connection.
}

func (this *redisPool) Get() (*redis.Client, bool) {
	var conn *redis.Client
	select {
	case conn = <-this.connections:
	default:
		conn, err := this.connFn()
		if err != nil {
			return nil, false
		}

		return conn, true
	}

	if err := this.testConn(conn); err != nil {
		return this.Get() // if connection is bad, get the next one in line until base case is hit, then create new client
	}

	return conn, true
}

func (this *redisPool) Close(conn *redis.Client) {
	select {
	case this.connections <- conn:
		return
	default:
		conn.Quit()
	}
}

func (this *redisPool) testConn(conn *redis.Client) error {
	if _, err := conn.Ping(); err != nil {
		conn.Quit()
		return err
	}

	return nil
}
func newcon()( *redis.Client,error) {
    var client *redis.Client
    client = redis.New()
    err := client.Connect(host, port)
    return client, err
}
