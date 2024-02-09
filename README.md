<h1>Concurrent program through workers</h1>

<p>HAs drastically reduced the latency for executing the jobs with concurrent programmes with a shared resource</p>

Ways to run :

1. go run main.go
2. time go run main.go ( To watch the time )

Output:

WITHOUT WORKER PARALELLY WORKING ON THE JOB (ONLY ONE WORKER):

1.  time go run main.go
started job 1 - worker 1 with job id 1
ended job  1 - worker 1 with job id 1
started job 1 - worker 1 with job id 2
4
ended job  1 - worker 1 with job id 2
started job 1 - worker 1 with job id 3
8
ended job  1 - worker 1 with job id 3
started job 1 - worker 1 with job id 4
12
ended job  1 - worker 1 with job id 4
started job 1 - worker 1 with job id 5
16
ended job  1 - worker 1 with job id 5
2 
Time needed is :, 10.005481417s
go run main.go  0.11s user 0.15s system 2% cpu 10.747 total


WITH TWO WORKERS PARALELLY WORKING ON THE JOB (TWO WORKERS):


2. time go run main.go
started job 2 - worker 1 with job id 2
started job 1 - worker 1 with job id 1
ended job  1 - worker 1 with job id 1
started job 1 - worker 1 with job id 3
4
ended job 2 - worker 1 with job id 2
started job 2 - worker 1 with job id 4
4
ended job 2 - worker 1 with job id 4
ended job  1 - worker 1 with job id 3
started job 2 - worker 1 with job id 5
12
8
ended job 2 - worker 1 with job id 5
10
Time needed is :, 6.003951s


WITH SIX WORKERS PARALELLY WORKING ON THE JOB (SIX WORKERS):


3. time go run main.go
started job 2 - worker 2 with job id 3
started job 1 - worker 2 with job id 1
started job 1 - worker 3 with job id 4
started job 2 - worker 1 with job id 2
started job 2 - worker 4 with job id 5
ended job  1 - worker 3 with job id 4
ended job 2 - worker 1 with job id 2
16
4
ended job  1 - worker 2 with job id 1
ended job 2 - worker 4 with job id 5
ended job 2 - worker 2 with job id 3
10
4
6
Time needed is :, 2.001423166s
go run main.go  0.10s user 0.11s system 7% cpu 2.709 total