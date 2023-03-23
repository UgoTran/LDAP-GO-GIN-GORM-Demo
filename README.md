# Module Role & Permission

### 1. Technical
   - go 1.20++
   - gin(web framework)
   - redis
   - IDE: IntelliJ / VSCode, DBeaver, Postman
### 2. Frameworks 
   - **[go gin]** run on {project_root}: **go get -u github.com/gin-gonic/gin**
### 3. Go document
   - document: https://gin-gonic.com/docs/quickstart/
   - https://www.golang-book.com/books/intro/8
   - https://hopding.com/interesting-golang-features
   - common error:  import cycle not allowed
   - repository: https://pkg.go.dev/
   - file **go.mod** define package, module and dependencies
   - *GOROOT & GOPATH:* https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html
   - Build container: https://www.codementor.io/@iedesnoek/deploying-a-go-app-with-gin-to-an-azure-container-app-22auximxoe
   - link repository: https://gitlab.hivetech.vn/microservice/backend/internal_role_permission_go
   - jira: https://j2.ossigroup.net/projects/PT/summary
   - confluence: https://c2.ossigroup.net/pages/viewpage.action?pageId=809399
   - funcs: If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
   - **JWT**: https://blog.logrocket.com/jwt-authentication-go/
### 4. Run & build
   - Run project: **go run src/main.go**
   - Build application: **go build -o target/RolePermissionApptilication src/main.go**
   - Build directory: build; **run** standalone file: **./RolePermissionApplication**
   - Docker build local:**docker build -t rp_app .**
   - Docker run container: **docker run -p 8200:8100 rp_app**
### 5. Project structure
   - src: 
     - *test*: unit test file
     - *resource*: config file yml, static resource
     - *com.hivetech.role_permission*: source code
     - **main.go**: file run application on *IDE*
   - Dockerfile*: build container for each env
   - go.mod: go library
### 6. Config env variable
    - ENV
### 7. External Library
    - Go Modules <pt_role_permission_go>
    - Go SDK 1.20
- LOG: **go get github.com/sirupsen/logrus**
- LDAP: **go get github.com/go-ldap/ldap/v3**
- Security
  - JWT: **go get github.com/golang-jwt/jwt**