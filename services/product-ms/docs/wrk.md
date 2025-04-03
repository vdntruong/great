# wrk - a HTTP benchmarking tool

[GitHub Repo](https://github.com/wg/wrk)

## Basic Usage

```bash
wrk -t12 -c1000 -d30s http://127.0.0.1:8081/products
```

This runs a benchmark for 30 seconds, using 12 threads, and keeping 400 HTTP connections open.

```
Running 30s test @ http://localhost:8081/products
  12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    68.90ms   49.26ms 399.53ms   66.23%
    Req/Sec   292.06    374.22     2.38k    91.73%
  104037 requests in 30.10s, 31.65MB read
  Socket errors: connect 757, read 111, write 0, timeout 0
Requests/sec:   3456.91
Transfer/sec:      1.05MB
```

## Command Line Options

```
-c, --connections: total number of HTTP connections to keep open with
                   each thread handling N = connections/threads

-d, --duration:    duration of the test, e.g. 2s, 2m, 2h

-t, --threads:     total number of threads to use

-s, --script:      LuaJIT script, see SCRIPTING

-H, --header:      HTTP header to add to request, e.g. "User-Agent: wrk"

    --latency:     print detailed latency statistics

    --timeout:     record a timeout if a response is not received within
                   this amount of time.
```
