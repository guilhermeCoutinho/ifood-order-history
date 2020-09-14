run:
	@go run *.go $(STARTING_DATE) $(END_DATE)

docker-build:
	@docker build . -t ifood-oh

docker-run:
	@docker run --rm -e "CURL_REQUEST=$$(eval pbpaste)"  ifood-oh