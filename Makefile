create:
	cp dbconfig.yml.template dbconfig.yml
	
migrate:
	env akeome=develop go run cmd/migrate.go

run:
	env akeome=develop go run main.go


docker/build:
	make -f .circleci/ci.mk go/build
	make -f .circleci/ci.mk docker/build

docker/run:
	docker run -it --rm 241556795328.dkr.ecr.ap-northeast-1.amazonaws.com/lavender
