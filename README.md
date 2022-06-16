## Entain BE Technical Test

Notes from Lachlan:
1. Would have liked to have had this back to the team @ entain sooner, ran into issues running the protoc commands which I lost a couple hours on for the first day - first PR will have version increments to external modules which was a side effect of me attempting to fix the issues I was having. Running the API and GRPC services all work fine with these version increments
2. Test files would ideally be in a single folder called Tests
3. In terms of linting, testing etc, I was using a few go plugins - these are: gotests@v1.6.0, gomodifytags@v1.16.0, staticcheck@v0.3.2, gopls@v0.8.4
4. There were a couple of places in the code I would normally refactor but I held off on making excessive changes as I am intending on submitting both backend and frontend coding challenges. 
5. In relation to point 4, tests in the sports service was minimal to achieve compeleting the frontend challenge too. In real life, i would have a full suite for the sports service





