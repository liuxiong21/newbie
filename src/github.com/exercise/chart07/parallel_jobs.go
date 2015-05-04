package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

type Job struct{
	Time string
	Name string
}

func RunParallelJobs(){
	jobList := []*Job{&Job{"2015-05-03 09:00:00","send mail"},
		&Job{"2015-05-03 19:00:00","compute access convert rate"}}
	done := make(chan bool,len(jobList))
	jobs := make(chan *Job)
	go func(){
		for index,job:= range jobList{
			jobs <- job
			job.Name = "tttt_"+strconv.Itoa(index)
			fmt.Println("Push new Job to channel")
		}
		close(jobs)
	}()
	go func(){
		for job:= range jobs{
			time.Sleep(time.Second*time.Duration(rand.Intn(10)))
			fmt.Printf("%s job handle finished start at %s\n",job.Name,job.Time)
			done <- true
		}
	}()
	for i:=0;i<len(jobList);i++{
		<-done
	}
}

func RunParallelJobs2(){
	jobList := []Job{Job{"2015-05-03 09:00:00","send mail"},
		Job{"2015-05-03 19:00:00","compute access convert rate"}}
	done := make(chan bool,len(jobList))
	jobs := make(chan Job)
	go func(){
		for index,job:= range jobList{
			jobs <- job
			job.Name = "tttt_"+strconv.Itoa(index)
			fmt.Println("Push new Job to channel")
		}
		close(jobs)
	}()
	go func(){
		for job:= range jobs{
			time.Sleep(time.Second*time.Duration(rand.Intn(10)))
			fmt.Printf("%s job handle finished start at %s\n",job.Name,job.Time)
			done <- true
		}
	}()
	for i:=0;i<len(jobList);i++{
		<-done
	}
}

