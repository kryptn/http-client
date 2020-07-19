post {
    url: "http://localhost:8000/auth"

    json {
        username: {env USERNAME MyUsername}
        password: {env PASSWORD MyPassword123}
    }
}