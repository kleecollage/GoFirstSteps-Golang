module github.com/kleecollage/GoFirstSteps-Golang

go 1.23.2

replace github.com/kleecollage/GoFirstSteps-Golang/greetings => ./greetings

require (
	github.com/kleecollage/GoFirstSteps-Golang/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2
	rsc.io/sampler v1.3.0 // indirect
)
