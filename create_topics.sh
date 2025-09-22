#/bin/bash

KAFKA_TOPICS="opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka1:9092"
PARTITIONS_DEFAULT=3
REPLICATION_FACTOR_DEFAULT=1
MESSAGE_TTL_MS_DEFAULT=86400000

$KAFKA_TOPICS --list

echo 'Creating kafka topics...'
$KAFKA_TOPICS --create --if-not-exists --topic users.create --partitions $PARTITIONS_DEFAULT --replication-factor $REPLICATION_FACTOR_DEFAULT --config retention.ms=$MESSAGE_TTL_MS_DEFAULT &&
$KAFKA_TOPICS --create --if-not-exists --topic users.get --partitions $PARTITIONS_DEFAULT --replication-factor $REPLICATION_FACTOR_DEFAULT --config retention.ms=$MESSAGE_TTL_MS_DEFAULT

EXIT_CODE=$?

echo 'List of topics:'
$KAFKA_TOPICS --list

echo "exit code: $EXIT_CODE"
exit $EXIT_CODE
