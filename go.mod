module github.com/kleecollage/GoFirstSteps-Golang

go 1.23.2

replace github.com/kleecollage/GoFirstSteps-Golang/greetings => ./greetings

require rsc.io/quote v1.5.2

require (
	github.com/kleecollage/GoFirstSteps-Golang/greetings v0.0.0-20241014225602-6a5a94832ab6
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
