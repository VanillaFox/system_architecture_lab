```
wrk -t8 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    39.88ms   15.76ms 392.37ms   77.82%
    Req/Sec   304.53     40.18   440.00     74.83%
  291073 requests in 2.00m, 49.41MB read
Requests/sec:   2424.94
Transfer/sec:    421.52KB
```

```
wrk -t16 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    40.38ms   28.51ms   1.37s    93.45%
    Req/Sec   153.15     23.30   280.00     73.77%
  293126 requests in 2.00m, 49.76MB read
Requests/sec:   2440.83
Transfer/sec:    424.29KB
```

```
wrk -t32 -c100 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  32 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    40.26ms   33.15ms   1.44s    95.22%
    Req/Sec    77.34     14.96   170.00     71.47%
  296287 requests in 2.00m, 50.30MB read
Requests/sec:   2467.86
Transfer/sec:    428.98KB
```

```
wrk -t16 -c200 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    85.70ms   31.11ms   1.60s    83.60%
    Req/Sec   141.28     23.90   260.00     74.17%
  270296 requests in 2.00m, 45.88MB read
  Socket errors: connect 0, read 51, write 0, timeout 0
Requests/sec:   2250.86
Transfer/sec:    391.26KB
```

```
wrk -t16 -c400 -d120s -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3MiLCJ1c2VybmFtZSI6Iml2YW4ifQ.3suRbMkRKSz8CrGyujU9syW-x4UuRYebGEdC5WubV5Q' http://localhost:8080/api/v1/users/username/ivan
Running 2m test @ http://localhost:8080/api/v1/users/username/ivan
  16 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   106.78ms   39.60ms   1.99s    82.54%
    Req/Sec   136.00     87.75   370.00     55.48%
  259390 requests in 2.00m, 44.03MB read
  Socket errors: connect 165, read 83, write 0, timeout 98
Requests/sec:   2159.81
Transfer/sec:    375.44KB
```

---

