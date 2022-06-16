## Entain BE Technical Test

Notes from Lachlan:
1. Would have liked to have had this back to the team @ entain sooner, ran into issues running the protoc commands which I lost a couple hours on for the first day - first PR will have version increments to external modules which was a side effect of me attempting to fix the issues I was having. Running the API and GRPC services all work fine with these version increments
2. Test files would ideally be in a single folder called Tests
3. used gofmt -s -w . for linting
4. In terms of formatting, testing etc, I was using a few go plugins - these are: gotests@v1.6.0, gomodifytags@v1.16.0, staticcheck@v0.3.2, gopls@v0.8.4
5. There were a couple of places in the code I would normally refactor but I held off on makin those changes to achieve submiting the coding challenge in a reasonable timeframe 
6. In relation to point 5, tests in the sports service was minimal to achieve compeleting the challenge in a reasonable timeframe without being excessive. In real life, i would have a full suite for the sports service on top of what was delivered for the racing service. 
7. Documentation exists as comments throughout code where I felt necessary. I didnt find many instances where the code wasnt self-descriptive or was in the same style as the existing codebase. As I mentioned in point 5, I could have refactored more and normally would in real life but again, reasonable timeframes and changing the codebase style too much would probably add noise to each PR and hide the actual changes rather than provide value and in that case, more documentation/comments would be required.





