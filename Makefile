
default: image run

test:
	curl -Method POST -Uri 127.0.0.1:8080 -ContentType 'application/json' -Body '{"a":2,"b":3}'

image:
	docker build -t my_application .

run:
	docker run -p 8080:8080 my_application


local:
	$Env:GOOS = "linux"; $Env:GOARCH = "amd64"
	go build

