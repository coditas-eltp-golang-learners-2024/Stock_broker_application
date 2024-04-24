# Define the directories of your microservices
MICROSERVICES := authentication watchlist orders portfolio alerts profile payments stocks search global-search scanners corporate-action mutual-funds
DIR_PATH := ./src/app
GO_VERSION := 1.20

generate-swagger:
	@for service in $(MICROSERVICES); do \
		echo "Generating Swagger for $$service"; \
		(cd $(DIR_PATH)/$$service && swag init -g main.go;) \
	done

go-mod-tidy:
	@for service in $(MICROSERVICES); do \
		echo "Running GO mod tidy for $$service"; \
		(cd $(DIR_PATH)/$$service && go mod tidy;) \
	done

test-all:
	@for service in $(MICROSERVICES); do \
		echo "Running test cases for $$service"; \
		(cd $(DIR_PATH)/$$service && go test ./tests/... -coverpkg=./... -coverprofile $$service.out -covermode count;) \
	done

# Copy api.yml in test setup enviroment
copy-api-yml:
	cp ./src/configs/apis.yml ./src/setupTest/testConfigs/apis.yml

# Generate swagger files and run GO mod tidy
all: generate-swagger go-mod-tidy copy-api-yml test-all

authentication: 
	cd $(DIR_PATH)/authentication && go mod tidy

watchlist: 
	cd $(DIR_PATH)/watchlist && go mod tidy
	cd $(DIR_PATH)/watchlist && go test ./tests/... -coverpkg=./... -coverprofile watchlist.out -covermode count

orders: 
	cd $(DIR_PATH)/orders && go mod tidy
	cd $(DIR_PATH)/orders && go test ./tests/... -coverpkg=./... -coverprofile orders.out -covermode count

portfolio: 
	cd $(DIR_PATH)/portfolio && go mod tidy
	cd $(DIR_PATH)/portfolio && go test ./tests/... -coverpkg=./... -coverprofile portfolio.out -covermode count

alerts: 
	cd $(DIR_PATH)/alerts && go mod tidy
	cd $(DIR_PATH)/alerts && go test ./tests/... -coverpkg=./... -coverprofile alerts.out -covermode count

profile: 
	cd $(DIR_PATH)/profile && go mod tidy
	cd $(DIR_PATH)/profile && go test ./tests/... -coverpkg=./... -coverprofile profile.out -covermode count

payments: 
	cd $(DIR_PATH)/payments && go mod tidy
	cd $(DIR_PATH)/payments && go test ./tests/... -coverpkg=./... -coverprofile payments.out -covermode count

stocks: 
	cd $(DIR_PATH)/stocks && go mod tidy
	cd $(DIR_PATH)/stocks && go test ./tests/... -coverpkg=./... -coverprofile stocks.out -covermode count

search: 
	cd $(DIR_PATH)/search && go mod tidy
	cd $(DIR_PATH)/search && go test ./tests/... -coverpkg=./... -coverprofile search.out -covermode count

global-search: 
	cd $(DIR_PATH)/global-search && go mod tidy
	cd $(DIR_PATH)/global-search && go test ./tests/... -coverpkg=./... -coverprofile global-search.out -covermode count

scanners: 
	cd $(DIR_PATH)/scanners && go mod tidy
	cd $(DIR_PATH)/scanners && go test ./tests/... -coverpkg=./... -coverprofile scanners.out -covermode count

corporate-action: 
	cd $(DIR_PATH)/corporate-action && go mod tidy
	cd $(DIR_PATH)/corporate-action && go test ./tests/... -coverpkg=./... -coverprofile corporate-action.out -covermode count

mutual-funds: 
	cd $(DIR_PATH)/mutual-funds && go mod tidy
	cd $(DIR_PATH)/mutual-funds && go test ./tests/... -coverpkg=./... -coverprofile mutual-funds.out -covermode count

create_service:
	./scripts/createMicroservice.sh $(servicename) $(port) $(GO_VERSION)

.PHONY: create_service
