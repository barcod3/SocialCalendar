APP = social-calendar-api

build:
	go build ./cmd/${APP}

run:
	go run ./cmd/${APP}

test:
	go test ./...

install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.43.0

lint: install-lint
	bin/golangci-lint run

docker-build:
	docker build --tag ${APP} .   

docker-run:
	docker run -p 8080:8080 -d \
		-e SOCIAL_CALENDAR_ID=${SOCIAL_CALENDAR_ID} \
		-e SOCIAL_CALENDAR_SECRET=${SOCIAL_CALENDAR_SECRET} \
		-e SOCIAL_CALENDAR_USERNAME=${SOCIAL_CALENDAR_USERNAME} \
		-e SOCIAL_CALENDAR_PASSWORD=${SOCIAL_CALENDAR_PASSWORD} \
		--name socialCalendarApi \
		${APP}