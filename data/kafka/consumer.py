#!/usr/bin/env python2

from kafka import KafkaConsumer

TOPIC_NAME = 'saed'



consumer = KafkaConsumer(
        TOPIC_NAME,
        bootstrap_servers='kafka:9092'
        )


for message in consumer:
    print(message)
