package main

import (
	"fmt"
	"time"

	"github.com/valyala/quicktemplate"
)

type EmailPlaceholder struct {
	CompanyName       string
	CompanyLocation   string
	JobSeekerName     string
	JobTitle          string
	JobDescription    string
	InterviewDate     string
	InterviewLocation string
}

func main() {
	str := "we are {%s CompanyName %} {%s CompanyLocation %} {%s JobSeekerName %} {% JobTitle %} {%s JobDescription %} {%s InterviewDate %} {%s InterviewLocation %}"
	p := EmailPlaceholder{
		CompanyName:       "demo",
		CompanyLocation:   "chengdu",
		JobSeekerName:     "luoji",
		JobTitle:          "",
		JobDescription:    "none",
		InterviewDate:     time.Now().Format("2006-01-02 15:04:05"),
		InterviewLocation: "chengdu",
	}

	buffer := quicktemplate.AcquireByteBuffer()
	writer := quicktemplate.AcquireWriter(buffer)

	//var buf bytes.Buffer
	//writer := quicktemplate.AcquireWriter(&buf)
	//writer.N().S(str)
	writer.E().S(p.CompanyName)
	writer.N().S(str)
	quicktemplate.ReleaseWriter(writer)
	fmt.Println("result:", string(buffer.B))
}
