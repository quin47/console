build:
	cd web && yarn build && cd ../ && rice embed-go  && go build