## TrueAccord Take Home Assignment
### Anurag Kalra

#### Technologies
For this assignment, I chose to write it using Go. I used Go because it is fast, relatively easy to build/compile/run,  provides effective native error handling, and robust testing ecosystem.

#### Dependencies

- github.com/stretchr/testify/assert	
	- Package assert provides a set of comprehensive testing tools for use with the normal Go testing system.
- github.com/jarcoal/httpmock
	- Easy mocking of http responses from external resources.
	

#### Installing, Running, and Testing the Program:

##### Installation and Running:
`$ git clone https://github.com/anuragkalra/alder.git`
`$ cd alder/main`
`$ go install`
`$ go build`
`$ ./main`

##### Running Tests:
`$ go test`

##### Running Tests with Cover Profile (May depend on machine):
`$ go test -timeout 30s -coverprofile=/var/folders/zk/pd98np1s5_z5c312_4k1vhtw0000gn/T/vscode-gokDyP22/go-code-cover github.com/anuragkalra/alder/main`

##### Expected Output:
`ok      github.com/anuragkalra/alder/main       1.080s  coverage: 86.2% of statements`