Setup instructions:
1. Makefile is present in the root directory. On the terminal type the command "make", and it will up the server.
2. Make sure that the project is placed inside the directory /go/src
3. Dep has been used for dependency management, total dependencies used are: 
    a. github.com/gorilla/handlers (for http handling)
    b. github.com/gorilla/mux (for http routing)
    c. github.com/sonyarouje/simdb (for handling JSON as DB)
4. Front-end code is present in the "front" directory.It will automatically be served once the server is up
5. Model test cases are present in the "model" package
6. For any other queries, shoot me a mail on prasangmisra@gmail.com. I am also available on 7698189874