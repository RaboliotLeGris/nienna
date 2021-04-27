import amqp
import psycopg2
from psycopg2 import pool
import os
import json
from videoDAO import VideoDAO


def create_queues(ch, pool):
    def on_jobs_result_message(message):
        conn = pool.getconn()
        if conn:
            payload = json.loads(message.body)
            slug = payload["slug"]
            if payload["event"] == "EventVideoProcessingSucceed":
                VideoDAO(conn).update_status(slug, "READY")
                conn.commit()
                ch.basic_ack(message.delivery_tag)
            elif payload["event"] == "EventVideoProcessingFail":
                VideoDAO(conn).update_status(slug, "FAILURE")
                conn.commit()
                ch.basic_ack(message.delivery_tag)
            else:
                ch.basic_reject(message.delivery_tag)

    ch.basic_consume(consumer_tag='pulsar', queue='nienna_jobs_result', callback=on_jobs_result_message)


def decompose_amqp_uri(uri):
    if uri is None:
        return None, None, None
    splitted_uri = uri[7:].split("@")
    return splitted_uri[1], splitted_uri[0].split(":")[0], splitted_uri[0].split(":")[1]


def main():
    db_pool = psycopg2.pool.SimpleConnectionPool(1, 20, os.getenv("DB_URI"))

    host, user, passwd = decompose_amqp_uri(os.getenv("AMQP_URI"))
    with amqp.Connection(host, user, passwd) as c:
        ch = c.channel()

        print("Creating queues")
        create_queues(ch, db_pool)

        print("Starting Pulsar")
        while True:
            c.drain_events()


if __name__ == '__main__':
    main()
