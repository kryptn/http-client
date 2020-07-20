post {
    url: "http://localhost:8080/post"

    json {
        username: {env MYUSER MyUsername}
        password: {env PASSWORD MyPassword123}
    }
}