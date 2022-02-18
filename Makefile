start:
	go build && \
	GRPC_PORT="6000" \
	HTTP_PORT="6001" \
	DB_PATH="host=0.0.0.0 user=postgres password=mysecretpassword dbname=pguser port=5432" \
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

migrationDb:
	go build && \
	GRPC_PORT="6000" \
	HTTP_PORT="6001" \
	DB_PATH="host=0.0.0.0 user=postgres password=mysecretpassword dbname=pguser port=5432" \
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
	PUBSUB_SUBSCRIPTION="exchange" ./pguser migrationDb