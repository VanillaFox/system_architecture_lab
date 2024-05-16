## 8 трэдов, 100 подключений, 2 минуты
| With cache | Avg ms | Stdev ms | Max ms | Total requests |
| --- | --- | --- | --- | --- |
| without | 43 | 45 | 559 | 332990 |
| with | 34 | 27 | 666 | 366558 |

## 16 трэдов, 100 подключений, 2 минуты
| With cache | Avg ms | Stdev ms | Max ms | Total requests |
| --- | --- | --- | --- | --- |
| without | 42 | 43 | 609 | 335015 |
| with | 42 | 43 | 608 | 371035 |

## 32 трэдов, 100 подключений, 2 минуты
| With cache | Avg ms | Stdev ms | Max ms | Total requests |
| --- | --- | --- | --- | --- |
| without | 43 | 45 | 639 | 331316 |
| with | 33 | 26 | 578 | 366600 |

## 16 трэдов, 200 подключений, 2 минуты
| With cache | Avg ms | Stdev ms | Max ms | Total requests |
| --- | --- | --- | --- | --- |
| without | 83 | 73 | 874 | 311632 |
| with | 34 | 26 | 579 | 372508 |

## 16 трэдов, 400 подключений, 2 минуты
| With cache | Avg ms | Stdev ms | Max ms | Total requests |
| --- | --- | --- | --- | --- |
| without | 94 | 88 | 1100 | 312502 |
| with | 80 | 40 | 820 | 356509 |

## Замеры без кэша

```
wrk -t8 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    42.97ms   44.75ms 559.27ms   86.60%
    Req/Sec   348.25     74.60   700.00     68.68%
  332990 requests in 2.00m, 56.53MB read
Requests/sec:   2772.73
Transfer/sec:    481.98KB
```

```
wrk -t16 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    42.29ms   43.37ms 608.46ms   86.34%
    Req/Sec   175.13     50.21   505.00     68.82%
  335015 requests in 2.00m, 56.87MB read
Requests/sec:   2789.62
Transfer/sec:    484.91KB
```

```
wrk -t32 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  32 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    43.06ms   44.67ms 638.86ms   86.52%
    Req/Sec    86.65     35.06   420.00     68.22%
  331316 requests in 2.00m, 56.24MB read
Requests/sec:   2758.84
Transfer/sec:    479.56KB
```

```
wrk -t16 -c200 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    83.80ms   73.81ms 874.21ms   67.23%
    Req/Sec   162.89     46.01   570.00     69.72%
  311632 requests in 2.00m, 52.90MB read
  Socket errors: connect 0, read 49, write 0, timeout 0
Requests/sec:   2594.79
Transfer/sec:    451.05KB
```

```
wrk -t16 -c400 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    94.46ms   88.52ms   1.10s    73.77%
    Req/Sec   177.75     88.22   494.00     67.78%
  312502 requests in 2.00m, 53.05MB read
  Socket errors: connect 165, read 96, write 0, timeout 0
Requests/sec:   2602.29
Transfer/sec:    452.35KB
```

## Замеры с кэшом

```
wrk -t8 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.71ms   26.91ms 666.23ms   90.02%
    Req/Sec   383.32     61.56     0.96k    77.96%
  366558 requests in 2.00m, 62.22MB read
Requests/sec:   3052.19
Transfer/sec:    530.56KB
```

```
wrk -t16 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.36ms   26.61ms 722.44ms   89.04%
    Req/Sec   194.06     39.57   580.00     76.46%
  371035 requests in 2.00m, 62.98MB read
Requests/sec:   3090.10
Transfer/sec:    537.15KB
```

```
wrk -t32 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  32 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.54ms   26.10ms 578.79ms   90.91%
    Req/Sec    95.91     22.67   505.00     71.72%
  366600 requests in 2.00m, 62.23MB read
Requests/sec:   3052.53
Transfer/sec:    530.62KB
```

```
wrk -t16 -c200 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    63.34ms   32.05ms 649.21ms   81.55%
    Req/Sec   194.72     36.94   424.00     72.15%
  372508 requests in 2.00m, 63.23MB read
  Socket errors: connect 0, read 50, write 0, timeout 0
Requests/sec:   3101.76
Transfer/sec:    539.17KB
```

```
wrk -t16 -c400 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    80.09ms   40.96ms 820.50ms   83.22%
    Req/Sec   186.39     91.67   616.00     60.39%
  356509 requests in 2.00m, 60.52MB read
  Socket errors: connect 165, read 90, write 0, timeout 0
Requests/sec:   2968.64
Transfer/sec:    516.03KB
```