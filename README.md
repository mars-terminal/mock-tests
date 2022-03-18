# mock-tests

    https://gist.github.com/thiagozs/4276432d12c2e5b152ea15b3f8b0012e - Tutorial


There im testing user create function

For create mock files

    mockgen -source=user.go -package=service -destination=user_mock.go

Also you can run tests in  
    
    service/user/read_test.go