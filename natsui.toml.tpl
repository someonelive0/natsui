# Natsui config file template


title = "Natsui"
license = "Copyright @ 2025"


[nats]
    name = "localhost"
    servers = [ "nats://127.0.0.1:4222" ]
    user = ""
    password = ""
    # timeout for connect and read timeout in seconds, default 10 seconds
    timeout = 10

