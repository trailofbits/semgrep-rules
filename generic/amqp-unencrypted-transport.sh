#!/bin/bash

# ruleid: amqp-unencrypted-transport
echo "Hello, World!" | amqp-publish --url=amqp://guest:guest@example.com:5672 --routing-key=test_queue

# ok: amqp-unencrypted-transport
echo "Hello, World!" | amqp-publish --url=amqps://guest:guest@example.com:5672 --routing-key=test_queue
