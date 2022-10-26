# platform-service-library

아름다운 가게 golang backend service를 위한 공통 라이브러리

1. conversion, format 체크등을 위한 공통 라이브러리
2. aws/sns : aws sns 서비스에 메세지를 publish 하기 위한 라이브러리
3. logging : echo framework을 기반으로 한 미들웨어에서 사용하기 위한 로깅 라이브러리
   - logging 정보 : trace id, service id, parent service id, service name, query param, path param, body 정보 등
     다양한 로깅 정보 저장
   - 기존 로깅 서비스와 연계한 service 내 behavior에 대한 stack trace 저장
   - trace 정보에서 sql query 별도 관리 가능
   - msa 서비스간 tracking을 위한 trace id, service id, parent service id 관리
   - console output 및 aws sns를 기반으로 한 db saving 기능
   - db saving : aws sns를 이용한 서비스 분리
     현재 서비스와 별도 서비스로 로깅 분리 - async 작업으로 기존 서비스에 영향주지 않음
     behavior_logs 테이블 및 behavior_log_details 저장

## Installation

go get -u github.com/beautiful-store/platform-service-library

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://gitlab.com/beautifulstore/platform-service-library.git
```
