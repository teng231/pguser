start:
	go build \
	GRPC_PORT="6000" \
	HTTP_PORT="6001" \
	DB_PATH="postgresql://postgres:123456@localhost:5432/pguser" \
	DB_NAME="pguser" \
	REDIS_ADDR="localhost:6379" \
	RSA_PRIVATE_PATH="assets/id_rsa" \
	RSA_PUBLIC_PATH="assets/id_rsa.pub" \
	SECRET="assets/secret" \
	ACCESS_TOKEN_DURATION="129600" \
	REFRESH_TOKEN_DURATION="172800" \
	SERVICE_NAME="exchange-service" \
	SYNCER_SERVER_URL="http://localhost:1001/syncer-server" \
	KAFKA_BROKERS="localhost:9092" \
	PUBSUB_SUBSCRIPTION="exchange" ./pguser start