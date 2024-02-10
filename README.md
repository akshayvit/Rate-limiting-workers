<h1>Rate-limiting-workers</h1>

<p>Has drastically reduced the latency for executing the jobs with concurrent programmes on a request with self made rate limiter put</p>

Accessing resource : "https://api.btcmarkets.net/v3/markets/BTC-AUD/ticker"

Ways to run :

1. go run server.go
2. time go run server.go ( To watch the time )

Output:



WITH <b>5 WORKERS</b> PARALELLY WORKING ON THE JOB (FIVE WORKERS):


3. time go run main.go

akshaychatterjee@Akshays-MacBook-Pro go-worker-echo-rate % time go run server.go
The best bid for BTC-AUD at 2024-02-10T09:35:31.348000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:35:31.943000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:35:31.943000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:35:31.943000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:35:31.943000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:35:41.016000Z body is 72368.88
The best bid for BTC-AUD at 2024-02-10T09:35:51.041000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:36:01.072000Z body is 72369.95
The best bid for BTC-AUD at 2024-02-10T09:36:11.077000Z body is 72368.89
The best bid for BTC-AUD at 2024-02-10T09:36:21.056000Z body is 72357.36
The best bid for BTC-AUD at 2024-02-10T09:36:31.061000Z body is 72346.03
The best bid for BTC-AUD at 2024-02-10T09:36:41.036000Z body is 72346.03
The best bid for BTC-AUD at 2024-02-10T09:36:51.052000Z body is 72346.03
The best bid for BTC-AUD at 2024-02-10T09:37:01.043000Z body is 72346.03
The best bid for BTC-AUD at 2024-02-10T09:37:11.041000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:37:21.037000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:37:30.136000Z body is 72335.79
The best bid for BTC-AUD at 2024-02-10T09:37:41.030000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:37:51.063000Z body is 72335.83
The best bid for BTC-AUD at 2024-02-10T09:38:00.645000Z body is 72336.5
The best bid for BTC-AUD at 2024-02-10T09:38:11.077000Z body is 72337.78
The best bid for BTC-AUD at 2024-02-10T09:38:21.054000Z body is 72336.25
The best bid for BTC-AUD at 2024-02-10T09:38:31.066000Z body is 72337.46
The best bid for BTC-AUD at 2024-02-10T09:38:41.051000Z body is 72338.28
The best bid for BTC-AUD at 2024-02-10T09:38:51.130000Z body is 72339.67
The best bid for BTC-AUD at 2024-02-10T09:39:00.804000Z body is 72336.42
The best bid for BTC-AUD at 2024-02-10T09:39:08.665000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:39:21.055000Z body is 72336.32
The best bid for BTC-AUD at 2024-02-10T09:39:31.060000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:39:41.017000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:39:51.048000Z body is 72336.18
The best bid for BTC-AUD at 2024-02-10T09:40:01.030000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:40:11.050000Z body is 72335.89
The best bid for BTC-AUD at 2024-02-10T09:40:21.044000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:40:31.035000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:40:41.047000Z body is 72335.8
The best bid for BTC-AUD at 2024-02-10T09:40:51.073000Z body is 72332.81
The best bid for BTC-AUD at 2024-02-10T09:41:01.059000Z body is 72332.81
The best bid for BTC-AUD at 2024-02-10T09:41:11.073000Z body is 72332.81
The best bid for BTC-AUD at 2024-02-10T09:41:20.948000Z body is 72332.82
The best bid for BTC-AUD at 2024-02-10T09:41:31.045000Z body is 72332.84
The best bid for BTC-AUD at 2024-02-10T09:41:41.055000Z body is 72332.81
The best bid for BTC-AUD at 2024-02-10T09:41:51.080000Z body is 72319.82
The best bid for BTC-AUD at 2024-02-10T09:42:01.053000Z body is 72319.69
The best bid for BTC-AUD at 2024-02-10T09:42:10.969000Z body is 72329.36
The best bid for BTC-AUD at 2024-02-10T09:42:21.111000Z body is 72329.36
The best bid for BTC-AUD at 2024-02-10T09:42:31.039000Z body is 72329.45
The best bid for BTC-AUD at 2024-02-10T09:42:40.180000Z body is 72329.37
The best bid for BTC-AUD at 2024-02-10T09:42:51.079000Z body is 72329.93
The best bid for BTC-AUD at 2024-02-10T09:43:01.089000Z body is 72331.14
The best bid for BTC-AUD at 2024-02-10T09:43:11.081000Z body is 72332.34
The best bid for BTC-AUD at 2024-02-10T09:43:21.067000Z body is 72333.5
The best bid for BTC-AUD at 2024-02-10T09:43:30.745000Z body is 72334.58
The best bid for BTC-AUD at 2024-02-10T09:43:41.112000Z body is 72329.47
The best bid for BTC-AUD at 2024-02-10T09:43:51.059000Z body is 72329.94
The best bid for BTC-AUD at 2024-02-10T09:44:01.077000Z body is 72331.22
The best bid for BTC-AUD at 2024-02-10T09:44:11.056000Z body is 72332.64
The best bid for BTC-AUD at 2024-02-10T09:44:21.089000Z body is 72333.48
The best bid for BTC-AUD at 2024-02-10T09:44:31.019000Z body is 72334.86
The best bid for BTC-AUD at 2024-02-10T09:44:41.077000Z body is 72329.98
The best bid for BTC-AUD at 2024-02-10T09:44:51.064000Z body is 72329.66
The best bid for BTC-AUD at 2024-02-10T09:45:01.052000Z body is 72329.37
The best bid for BTC-AUD at 2024-02-10T09:45:10.765000Z body is 72329.36
The best bid for BTC-AUD at 2024-02-10T09:45:20.995000Z body is 72329.5
The best bid for BTC-AUD at 2024-02-10T09:45:30.366000Z body is 72321.91
The best bid for BTC-AUD at 2024-02-10T09:45:40.584000Z body is 72321.91
The best bid for BTC-AUD at 2024-02-10T09:45:51.038000Z body is 72321.91
The best bid for BTC-AUD at 2024-02-10T09:46:01.041000Z body is 72321.91
The best bid for BTC-AUD at 2024-02-10T09:46:11.037000Z body is 72308.91
The best bid for BTC-AUD at 2024-02-10T09:46:21.042000Z body is 72289.17
The best bid for BTC-AUD at 2024-02-10T09:46:30.933000Z body is 72290.33
The best bid for BTC-AUD at 2024-02-10T09:46:40.353000Z body is 72271.98
The best bid for BTC-AUD at 2024-02-10T09:46:51.035000Z body is 72272
The best bid for BTC-AUD at 2024-02-10T09:47:01.046000Z body is 72273.8
The best bid for BTC-AUD at 2024-02-10T09:47:11.184000Z body is 72275.4
The best bid for BTC-AUD at 2024-02-10T09:47:21.156000Z body is 72272.14
The best bid for BTC-AUD at 2024-02-10T09:47:30.840000Z body is 72274
The best bid for BTC-AUD at 2024-02-10T09:47:41.059000Z body is 72275.46
The best bid for BTC-AUD at 2024-02-10T09:47:51.053000Z body is 72276.74
The best bid for BTC-AUD at 2024-02-10T09:48:00.906000Z body is 72272.21
The best bid for BTC-AUD at 2024-02-10T09:48:02.265000Z body is 72271.49
The best bid for BTC-AUD at 2024-02-10T09:48:21.040000Z body is 72261.68
The best bid for BTC-AUD at 2024-02-10T09:48:31.071000Z body is 72253.07
The best bid for BTC-AUD at 2024-02-10T09:48:41.068000Z body is 72253.51
The best bid for BTC-AUD at 2024-02-10T09:48:51.021000Z body is 72253.55
The best bid for BTC-AUD at 2024-02-10T09:49:01.076000Z body is 72260.91
The best bid for BTC-AUD at 2024-02-10T09:49:10.789000Z body is 72310.44
The best bid for BTC-AUD at 2024-02-10T09:49:21.083000Z body is 72311.68
The best bid for BTC-AUD at 2024-02-10T09:49:31.043000Z body is 72310.44
The best bid for BTC-AUD at 2024-02-10T09:49:41.062000Z body is 72271.89
The best bid for BTC-AUD at 2024-02-10T09:49:51.072000Z body is 72273.19
The best bid for BTC-AUD at 2024-02-10T09:50:01.067000Z body is 72274.47
The best bid for BTC-AUD at 2024-02-10T09:50:05.709000Z body is 72271.81
The best bid for BTC-AUD at 2024-02-10T09:50:21.077000Z body is 72273.05
The best bid for BTC-AUD at 2024-02-10T09:50:31.241000Z body is 72272.13
The best bid for BTC-AUD at 2024-02-10T09:50:40.937000Z body is 72273.87
The best bid for BTC-AUD at 2024-02-10T09:50:51.024000Z body is 72272.05
The best bid for BTC-AUD at 2024-02-10T09:51:01.048000Z body is 72271.83
The best bid for BTC-AUD at 2024-02-10T09:51:11.079000Z body is 72261.23
The best bid for BTC-AUD at 2024-02-10T09:51:21.135000Z body is 72262.62
The best bid for BTC-AUD at 2024-02-10T09:51:31.079000Z body is 72264.07
The best bid for BTC-AUD at 2024-02-10T09:51:40.074000Z body is 72265.12
The best bid for BTC-AUD at 2024-02-10T09:51:50.976000Z body is 72261.23
The best bid for BTC-AUD at 2024-02-10T09:52:00.286000Z body is 72260.81
The best bid for BTC-AUD at 2024-02-10T09:52:11.083000Z body is 72261.83
The best bid for BTC-AUD at 2024-02-10T09:52:21.056000Z body is 72262.93
The best bid for BTC-AUD at 2024-02-10T09:52:31.080000Z body is 72264.27
The best bid for BTC-AUD at 2024-02-10T09:52:41.089000Z body is 72265.65
The best bid for BTC-AUD at 2024-02-10T09:52:51.068000Z body is 72255.5
The best bid for BTC-AUD at 2024-02-10T09:53:01.016000Z body is 72255.39
The best bid for BTC-AUD at 2024-02-10T09:53:11.090000Z body is 72255.98
The best bid for BTC-AUD at 2024-02-10T09:53:20.155000Z body is 72255.11
The best bid for BTC-AUD at 2024-02-10T09:53:30.506000Z body is 72214.85
The best bid for BTC-AUD at 2024-02-10T09:53:41.079000Z body is 72216.06
The best bid for BTC-AUD at 2024-02-10T09:53:51.085000Z body is 72215.17
The best bid for BTC-AUD at 2024-02-10T09:54:00.476000Z body is 72214.85
The best bid for BTC-AUD at 2024-02-10T09:54:10.978000Z body is 72216.17
The best bid for BTC-AUD at 2024-02-10T09:54:21.644000Z body is 72215.41
The best bid for BTC-AUD at 2024-02-10T09:54:31.071000Z body is 72215.23
The best bid for BTC-AUD at 2024-02-10T09:54:41.094000Z body is 72215.7
The best bid for BTC-AUD at 2024-02-10T09:54:51.052000Z body is 72214.92
The best bid for BTC-AUD at 2024-02-10T09:55:01.068000Z body is 72215.42
The best bid for BTC-AUD at 2024-02-10T09:55:11.094000Z body is 72214.92
The best bid for BTC-AUD at 2024-02-10T09:55:21.048000Z body is 72167.45
The best bid for BTC-AUD at 2024-02-10T09:55:31.041000Z body is 72168.07
The best bid for BTC-AUD at 2024-02-10T09:55:41.037000Z body is 72169.99
The best bid for BTC-AUD at 2024-02-10T09:55:51.063000Z body is 72171.95
^Csignal: interrupt
go run server.go  0.41s user 0.22s system 0% cpu 20:28.16 total