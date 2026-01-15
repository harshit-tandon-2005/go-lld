# go-lld
Low Level Design based on golang
The repository includes learing about SOLID Principles and commonly used Design Patterns


## SOLID Principles

To run the code for SOLID priciples use the command:

`    
    cd design-principles/{specific-directory}
`

`
    go run . {argyment}
`

argument- [1, 2]

### Problem: 
go run main.go compiles only that file; types in other files were missing.
### Solution: 
run the package (go run .) or include the other files so the compiler sees BadImplementation and GoodImplementation.