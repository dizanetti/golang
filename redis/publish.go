package main

func publisher(key string, value interface{}) {
	rdb := connectionRedis()
	err := rdb.Set(ctx, key, value, 0).Err()

	if err != nil {
		panic(err)
	}

	defer rdb.Close()
}
