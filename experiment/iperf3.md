##### submariner

```
subctl export service -n test iperf3-server
kubectl exec -it -n test iperf3-client -- /bin/bash

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 1M
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec  0.048 ms  0/29292 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec  0.042 ms  0/29292 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec  0.050 ms  0/29292 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec  0.065 ms  0/29292 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec  0.055 ms  0/29292 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec  0.035 ms  0/29295 (0%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 2M
s1:  [SUM]   0.00-5.01   sec  3.57 MBytes  5.97 Mbits/sec  1.237 ms  106/58518 (0.18%)  receiver
s1:  [SUM]   0.00-5.00   sec  3.58 MBytes  6.00 Mbits/sec  0.035 ms  0/58587 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  3.58 MBytes  6.00 Mbits/sec  0.052 ms  0/58584 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  3.58 MBytes  6.00 Mbits/sec  0.071 ms  0/58584 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  3.58 MBytes  6.00 Mbits/sec  0.088 ms  0/58584 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  3.58 MBytes  6.00 Mbits/sec  0.045 ms  0/58587 (0%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 3M
s1:  [SUM]   0.00-5.00   sec  5.35 MBytes  8.98 Mbits/sec  0.043 ms  176/87882 (0.2%)  receiver
s1:  [SUM]   0.00-5.00   sec  5.36 MBytes  9.00 Mbits/sec  0.061 ms  0/87879 (0%)  receiver
s1:  [SUM]   0.00-5.00   sec  5.34 MBytes  8.96 Mbits/sec  0.030 ms  369/87867 (0.42%)  receiver
s1:  [SUM]   0.00-5.00   sec  5.35 MBytes  8.97 Mbits/sec  0.043 ms  249/87879 (0.28%)  receiver
s1:  [SUM]   0.00-5.00   sec  5.35 MBytes  8.98 Mbits/sec  0.081 ms  220/87876 (0.25%)  receiver
s1:  [SUM]   0.00-5.00   sec  5.27 MBytes  8.84 Mbits/sec  0.056 ms  1516/87879 (1.7%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 4M
s1:  [SUM]   0.00-5.00   sec  7.12 MBytes  11.9 Mbits/sec  0.053 ms  590/117180 (0.5%)  receiver
s1:  [SUM]   0.00-5.00   sec  7.12 MBytes  11.9 Mbits/sec  0.092 ms  527/117171 (0.45%)  receiver
s1:  [SUM]   0.00-5.00   sec  7.14 MBytes  12.0 Mbits/sec  0.082 ms  237/117174 (0.2%)  receiver
s1:  [SUM]   0.00-5.00   sec  7.13 MBytes  12.0 Mbits/sec  0.041 ms  282/117171 (0.24%)  receiver
s1:  [SUM]   0.00-5.00   sec  7.09 MBytes  11.9 Mbits/sec  0.052 ms  1078/117171 (0.92%)  receiver
s1:  [SUM]   0.00-5.00   sec  7.09 MBytes  11.9 Mbits/sec  0.484 ms  856/117011 (0.73%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 5M
s1:  [SUM]   0.00-5.00   sec  8.85 MBytes  14.9 Mbits/sec  0.030 ms  1415/146469 (0.97%)  receiver
s1:  [SUM]   0.00-5.00   sec  8.85 MBytes  14.8 Mbits/sec  0.063 ms  1444/146463 (0.99%)  receiver
s1:  [SUM]   0.00-5.00   sec  8.85 MBytes  14.8 Mbits/sec  0.083 ms  1313/146346 (0.9%)  receiver
s1:  [SUM]   0.00-5.00   sec  8.84 MBytes  14.8 Mbits/sec  0.055 ms  1589/146466 (1.1%)  receiver
s1:  [SUM]   0.00-5.00   sec  8.89 MBytes  14.9 Mbits/sec  0.091 ms  780/146466 (0.53%)  receiver
s1:  [SUM]   0.00-5.00   sec  8.89 MBytes  14.9 Mbits/sec  0.043 ms  794/146466 (0.54%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 6M
s1:  [SUM]   0.00-5.21   sec  10.6 MBytes  17.1 Mbits/sec  0.099 ms  2334/175765 (1.3%)  receiver
s1:  [SUM]   0.00-5.00   sec  10.4 MBytes  17.5 Mbits/sec  0.063 ms  5223/175761 (3%)  receiver
s1:  [SUM]   0.00-5.00   sec  10.5 MBytes  17.6 Mbits/sec  0.050 ms  4082/175753 (2.3%)  receiver
s1:  [SUM]   0.00-5.00   sec  10.5 MBytes  17.7 Mbits/sec  0.048 ms  3352/175779 (1.9%)  receiver
s1:  [SUM]   0.00-5.00   sec  10.6 MBytes  17.7 Mbits/sec  0.081 ms  2636/175753 (1.5%)  receiver
s1:  [SUM]   0.00-5.00   sec  10.5 MBytes  17.6 Mbits/sec  0.021 ms  3701/175764 (2.1%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 7M
s1:  [SUM]   0.00-5.02   sec  12.0 MBytes  20.0 Mbits/sec  0.113 ms  8795/205043 (4.3%)  receiver
s1:  [SUM]   0.00-5.01   sec  11.2 MBytes  18.7 Mbits/sec  0.118 ms  21923/205068 (11%)  receiver
s1:  [SUM]   0.00-5.02   sec  10.9 MBytes  18.2 Mbits/sec  0.111 ms  26359/205045 (13%)  receiver
s1:  [SUM]   0.00-5.02   sec  10.9 MBytes  18.1 Mbits/sec  0.096 ms  27220/205047 (13%)  receiver
s1:  [SUM]   0.00-5.01   sec  10.4 MBytes  17.4 Mbits/sec  0.144 ms  34678/204955 (17%)  receiver
s1:  [SUM]   0.00-5.00   sec  11.1 MBytes  18.6 Mbits/sec  0.058 ms  23438/205041 (11%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 8M
s1:  [SUM]   0.00-5.21   sec  9.98 MBytes  16.1 Mbits/sec  0.101 ms  70862/234333 (30%)  receiver
s1:  [SUM]   0.00-5.01   sec  10.1 MBytes  17.0 Mbits/sec  0.122 ms  68065/234347 (29%)  receiver
s1:  [SUM]   0.00-5.01   sec  10.9 MBytes  18.3 Mbits/sec  0.105 ms  55010/234330 (23%)  receiver
s1:  [SUM]   0.00-5.02   sec  10.5 MBytes  17.5 Mbits/sec  0.179 ms  62693/234349 (27%)  receiver
s1:  [SUM]   0.00-5.21   sec  10.1 MBytes  16.3 Mbits/sec  0.146 ms  66418/232473 (28%)  receiver
s1:  [SUM]   0.00-5.00   sec  11.0 MBytes  18.4 Mbits/sec  0.030 ms  54845/234342 (23%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 9M
s1:  [SUM]   0.00-5.21   sec  11.7 MBytes  18.8 Mbits/sec  0.096 ms  72319/263651 (27%)  receiver
s1:  [SUM]   0.00-5.01   sec  10.9 MBytes  18.2 Mbits/sec  0.149 ms  85292/263625 (32%)  receiver
s1:  [SUM]   0.00-5.21   sec  10.5 MBytes  17.0 Mbits/sec  0.105 ms  90965/263615 (34%)  receiver
s1:  [SUM]   0.00-5.02   sec  10.2 MBytes  17.0 Mbits/sec  0.145 ms  97091/263634 (37%)  receiver
s1:  [SUM]   0.00-5.01   sec  10.1 MBytes  16.8 Mbits/sec  0.084 ms  98904/263613 (38%)  receiver
s1:  [SUM]   0.00-5.00   sec  11.5 MBytes  19.3 Mbits/sec  0.035 ms  75253/263628 (29%)  receiver

iperf3 -c iperf3-server-2.test.svc.clusterset.local -u -t 5 -J -i 5 -l 64 -P 3 -b 10M
s1:  [SUM]   0.00-5.21   sec  10.1 MBytes  16.3 Mbits/sec  0.118 ms  126302/292312 (43%)  receiver
s1:  [SUM]   0.00-5.01   sec  9.20 MBytes  15.4 Mbits/sec  0.132 ms  142133/292829 (49%)  receiver
s1:  [SUM]   0.00-5.02   sec  10.7 MBytes  17.8 Mbits/sec  0.132 ms  118389/292962 (40%)  receiver
s1:  [SUM]   0.00-5.21   sec  10.2 MBytes  16.5 Mbits/sec  0.112 ms  125402/292879 (43%)  receiver
s1:  [SUM]   0.00-5.00   sec  9.63 MBytes  16.1 Mbits/sec  0.056 ms  135172/292917 (46%)  receiver
s1:  [SUM]   0.00-5.00   sec  11.5 MBytes  19.2 Mbits/sec  0.047 ms  105097/292948 (36%)  receiver


iperf3 -c iperf3-server-2.test.svc.clusterset.local -t 5 -J -i 5 -l 64  -P 3 -b 1M
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.0 Mbits/sec    0             sender
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.0 Mbits/sec                  receiver
812

iperf3 -c iperf3-server-2.test.svc.clusterset.local -t 5 -J -i 5 -l 64  -P 3 -b 10M
s1:  [SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec    0             sender
s1:  [SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec                  receiver
1071

iperf3 -c iperf3-server-2.test.svc.clusterset.local -t 5 -J -i 5 -l 64  -P 3 -b 50M
[SUM]   0.00-5.00   sec  89.4 MBytes   150 Mbits/sec    6             sender
[SUM]   0.00-5.00   sec  89.4 MBytes   150 Mbits/sec                  receiver
3452

iperf3 -c iperf3-server-2.test.svc.clusterset.local -t 5 -J -i 5 -l 64  -P 3 -b 100M
[SUM]   0.00-5.00   sec   115 MBytes   193 Mbits/sec   50             sender
[SUM]   0.00-5.00   sec   115 MBytes   193 Mbits/sec                  receiver
17534

iperf3 -c iperf3-server-2.test.svc.clusterset.local -t 5 -J -i 5 -l 64  -P 3 -b 1000M
s1:  [SUM]   0.00-5.00   sec   122 MBytes   204 Mbits/sec   69             sender
s1:  [SUM]   0.00-5.00   sec   122 MBytes   204 Mbits/sec                  receiver
24011
```

##### istio

```
kubectl exec -it -n sample iperf3-client -- /bin/bash

iperf3 -c iperf3-server-2.sample -t 5 -J -i 5 -l 64  -P 3 -b 1M
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.0 Mbits/sec    0             sender
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.0 Mbits/sec                  receiver
213

iperf3 -c iperf3-server-2.sample -t 5 -J -i 5 -l 64  -P 3 -b 10M
s1:  [SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec    0             sender
s1:  [SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec                  receiver
259

iperf3 -c iperf3-server-2.sample -t 5 -J -i 5 -l 64  -P 3 -b 50M
[SUM]   0.00-5.00   sec  89.3 MBytes   150 Mbits/sec    0             sender
[SUM]   0.00-5.01   sec  89.3 MBytes   150 Mbits/sec                  receiver
967

iperf3 -c iperf3-server-2.sample -t 5 -J -i 5 -l 64  -P 3 -b 100M
[SUM]   0.00-5.00   sec   107 MBytes   179 Mbits/sec    0             sender
[SUM]   0.00-5.00   sec   107 MBytes   179 Mbits/sec                  receiver
2275

iperf3 -c iperf3-server-2.sample -t 5 -J -i 5 -l 64  -P 3 -b 1000M
[SUM]   0.00-5.00   sec   111 MBytes   186 Mbits/sec    0             sender
[SUM]   0.00-5.00   sec   111 MBytes   186 Mbits/sec                  receiver
3200
```

##### skupper

```
skupper expose -n skupper service/iperf3-server-2 --address iperf3-server-20

kubectl exec -it -n skupper iperf3-client -- /bin/bash

iperf3 -c iperf3-server-20 -t 5 -J -i 5 -l 64 -P 3 -b 1M
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.0 Mbits/sec    0             sender
s1:  [SUM]   0.00-5.00   sec  1.79 MBytes  3.0 Mbits/sec                  receiver
223

iperf3 -c iperf3-server-20 -t 5 -J -i 5 -l 64 -P 3 -b 10M
[SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec    0             sender
[SUM]   0.00-5.01   sec  17.9 MBytes  29.9 Mbits/sec                  receiver
575

iperf3 -c iperf3-server-20 -t 5 -J -i 5 -l 64 -P 3 -b 50M
[SUM]   0.00-5.00   sec  78.4 MBytes   132 Mbits/sec   80             sender
[SUM]   0.00-5.01   sec  70.3 MBytes   118 Mbits/sec                  receiver
7558

iperf3 -c iperf3-server-20 -t 5 -J -i 5 -l 64 -P 3 -b 100M
[SUM]   0.00-5.00   sec  94.4 MBytes   158 Mbits/sec  106             sender
[SUM]   0.00-5.07   sec  88.1 MBytes   146 Mbits/sec                  receiver
10224

iperf3 -c iperf3-server-20 -t 5 -J -i 5 -l 64 -P 3 -b 1000M
[SUM]   0.00-5.00   sec  91.1 MBytes   153 Mbits/sec  192             sender
[SUM]   0.00-5.01   sec  90.1 MBytes   151 Mbits/sec                  receiver
19612
```

##### X-ClustersLink

```
subctl export service -n ns2 iperf3-server
kubectl exec -it -n ns2 iperf3-client -- /bin/bash

iperf3 -c 242.5.255.254 -t 5 -J -i 5 -l 64 -P 3 -b 1M
[SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec    0             sender
[SUM]   0.00-5.00   sec  1.79 MBytes  3.00 Mbits/sec                  receiver
340

iperf3 -c 242.5.255.254 -t 5 -J -i 5 -l 64 -P 3 -b 10M
[SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec    0             sender
[SUM]   0.00-5.00   sec  17.9 MBytes  30.0 Mbits/sec                  receiver
566

iperf3 -c 242.5.255.254 -t 5 -J -i 5 -l 64 -P 3 -b 50M
[SUM]   0.00-5.00   sec  89.4 MBytes   150 Mbits/sec    0             sender
[SUM]   0.00-5.00   sec  89.4 MBytes   150 Mbits/sec                  receiver


iperf3 -c 242.5.255.254 -t 5 -J -i 5 -l 64 -P 3 -b 100M
[SUM]   0.00-5.00   sec   111 MBytes   186 Mbits/sec    0             sender
[SUM]   0.00-5.01   sec   109 MBytes   183 Mbits/sec                  receiver
715

iperf3 -c 242.5.255.254 -t 5 -J -i 5 -l 64 -P 3 -b 1000M
[SUM]   0.00-5.00   sec   118 MBytes   197 Mbits/sec    3             sender
[SUM]   0.00-5.00   sec   117 MBytes   197 Mbits/sec                  receiver
1013
```

```
vpc-gw pod cpu使用率

不转发
CPU:   7% usr   2% sys   0% nic  88% idle   0% io   0% irq   0% sirq

1 pod iperf
CPU:  18% usr  24% sys   0% nic  42% idle   0% io   0% irq  13% sirq
2 pod iperf
CPU:  16% usr  21% sys   0% nic  25% idle   0% io   0% irq  36% sirq
3 pod iperf
CPU:  22% usr  34% sys   0% nic   2% idle   0% io   0% irq  41% sirq
4 pod iperf
CPU:  21% usr  30% sys   0% nic   5% idle   0% io   0% irq  42% sirq
5 pod iperf
CPU:  22% usr  34% sys   0% nic   0% idle   0% io   0% irq  42% sirq
6 pod iperf
CPU:  18% usr  31% sys   0% nic   4% idle   0% io   0% irq  45% sirq
7 pod iperf
CPU:  20% usr  32% sys   0% nic   0% idle   0% io   0% irq  47% sirq
8 pod iperf
CPU:  19% usr  31% sys   0% nic   0% idle   0% io   0% irq  48% sirq
9 pod iperf
CPU:  21% usr  29% sys   0% nic   4% idle   0% io   0% irq  45% sirq
10 pod iperf
CPU:  10% usr  40% sys   0% nic   0% idle   0% io   0% irq  50% sirq

[SUM]   0.00-10.00  sec  86.3 MBytes  72.4 Mbits/sec   93             sender
[SUM]   0.00-10.00  sec  86.3 MBytes  72.4 Mbits/sec                  receiver

理论带宽上限不大于1G左右
```

