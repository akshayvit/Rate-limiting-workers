package main
import (
    "fmt"
    "time"
)

func worker1(id int,jobs <-chan int,results  chan<-int) {
    for j:=range jobs {
        fmt.Println("started job 1 - worker",id,"with job id",j)
        time.Sleep(time.Second*2)
        fmt.Println("ended job  1 - worker",id,"with job id",j)
        results<-j*4
    }
}

func worker2(id int , jobs<-chan int,results chan<-int) {
    for j:=range jobs {
        fmt.Println("started job 2 - worker",id,"with job id",j)
        time.Sleep(time.Second*2)
        fmt.Println("ended job 2 - worker",id,"with job id",j)
        results<-j*2
    }
}

func main() {
   const jobsn=5
   jobs:=make(chan int,jobsn)
   results:=make(chan int,jobsn)
   
   for j:=1;j<=1;j++ {
       go worker1(j,jobs,results)
   }
   t1:=time.Now()
   for j:=1;j<=jobsn;j++ {
       jobs<-j
   }
   
   close(jobs)
   for j:=1;j<=jobsn;j++ {
       fmt.Println(<-results)
   }
   t2:=time.Now()
   fmt.Println("Time needed is :,",t2.Sub(t1))
}