#!/usr/bin/env python2

from kafka import KafkaProducer

TOPIC_NAME = 'saed'


producer = KafkaProducer(
        bootstrap_servers='kafka:9092'
        )

for a in range(3):
    producer.send(TOPIC_NAME, b'Test Message')
    producer.flush()
