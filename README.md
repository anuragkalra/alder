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
`ok  	github.com/anuragkalra/alder/main	0.867s	coverage: 75.2% of statements`

#### Program Design
All program logic is in main/. Data structures and funcs are described in individual files for debts, payments, and payment plans. In the mock/ directory, I organized the http mock responders, as well as their corresponding responses. For this assignment I prioritized clean and simple design, readability, and program correctness. The program does make some assumptions which I will detail below.

##### Assumptions
- For debts, the **remaining_amount** field is calculated based on **amount_to_pay** coming from the corresponding Payment Plan, minus the sum of all payments.
- For the **next_payment_due_date** associated with a Debt object, when the first payment is sent after the original **start_date** associated with the Payment Plan, we compute the value by taking the most recent payment and adding the corresponding time interval.
- When payments have been contributed to a payment plan prior to the **start_date** associated, the original Payment Plan **start_date** will be used to compute the **next_payment_due_date** for the Debt.

##### Improvements
- When we are working with APIs, we should prepare for all sorts of responses, reliability, and behaviors. For this assignment I generally assumed that the API provided was reliable and returns predictable response formats. If working with a more complex API, it would be important to cover a larger range of behaviors and respond accordingly. In the situation where an API requires authentication, we may have to provide API keys, complete an Oauth exchange, or interface with JWT. We would have to deal specifically with HTTP status codes 401, 403, etc. Given that we are only retrieving information (making GET requests) we can generally expect 200 OK responses. If we were updating/deleting/posting data we would need to account for those HTTP responses as well.
- The API urls provided are coded as static strings in this program. In the future this program could be easily extended to accept urls as command line parameters.
- While computing some of the fields on the Debt objects, some information from the Payments is traversed more than once. While this isn't necessary to compute the result, this provided a more readable and extendable program. If we were expecting to process larger amounts of data and efficiency was paramount, I could rewrite this code block.
