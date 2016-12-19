# Redlot
Redlot is a KV database.



## Implemented

### INFO

​	This command will return the server informations.

### KV

- `SET key value` Set value by the key.
- `GET key` Get value by the key.
- `INCR key number`
- `DEL key` Delete value by the key.
- `EXISTS key` Check key is exists.
- `SETX tempKey value TTL` Set value by the key and expire it after timeout.
- SETEX (Alias of SETX)
- `TTL key` Return the lifetime of the key.
- `EXPIRE key` Expire the key after timeout.
- `KEYS start end limit` List keys in the range.
- `RKEYS start end limit` Reverse list keys in the range.
- `SCAN start end limit` List KV pair that keys in the range.
- `RSCAN start end limit` Reverse list KV pair that keys in the range.
- `MULTI_GET key1 key2 ...` Batch read data from db.
- `MULTI_SET key1 value1 key2 value2 ...` Batch write data to db.
- `MULTI_DEL key1 key2 ...` Batch delete value by keys.

### Hash

- `HSET key field value` Hset will set a hashmap value by the field.
- `HGET key field value` Hset will return a hashmap value by the field.
- `HDEL key` Hdel will delete a hashmap value by the field.
- `HINCR key field` Hincr will incr a hashmap value by the field.
- `HINCRBY key field number` Hincrby will incr number a hashmap value by the field.
- `HEXISTS key field` Hexists will check the hashmap field is exists.
- `HSIZE key` Hsize will return the hashmap size.
- `HKEYS key start end limit` Hkeys will list the hashmap fields in the range.
- `HRKEYS key start end limit` Hrkeys will reverse list the hashmap fields in the range.
- `HGETALL key` Hgetall will list all fields/value in the hashmap.
- `HLIST key start end limit` Hlist will list all hashmap in the range.
- `HRLIST key start end limit` Hrlist will reverse list all hashmap in the range.
- `HSCAN key start end limit` Hscan will list fields/value of the hashmap in the range.
- `HRSCAN key start end limit` Hrscan will reverse list fields/value of the hashmap in the range.
- `HCLEAR key` Hclear will remove all value in the hashmap.


## TODO

### Geo data



## Example

```bash
go run example/server/main.go
```

### Benchmark

```bash
redis-benchmark -p 9999 -t set,get -r 100000 -n 100000 -c 150
```

