using System;
using System.Collections;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;

namespace pulsar.clients
{
    public class AmqpClient
    {
        private readonly string _uri;
        private IConnection _conn;
        private IModel _channel;

        public AmqpClient(string uri)
        {
            if (!uri.StartsWith("amqp://"))
            {
                throw new ArgumentException("Invalid AMQP URI: " + uri);
            }

            this._uri = uri;
        }

        ~AmqpClient()
        {
            this._channel.Close();
            this._conn.Close();
        }

        public AmqpClient Connect()
        {
            ConnectionFactory factory = new ConnectionFactory();
            factory.Uri = new Uri(this._uri);
            factory.ClientProvidedName = "pulsar";
            factory.AutomaticRecoveryEnabled = true;
            factory.NetworkRecoveryInterval = TimeSpan.FromSeconds(10);

            this._conn = factory.CreateConnection();
            this._channel = this._conn.CreateModel();

            return this;
        }

        public void DeclareQueues(params string[] queues)
        {
            foreach (string queue in queues)
            {
                this._channel.QueueDeclare(queue, false, false, false, null);
            }
        }

        public string AddConsumer(string queue, EventHandler<BasicDeliverEventArgs> f)
        {
            var consumer = new EventingBasicConsumer(this._channel);
            consumer.Received += f;

            return this._channel.BasicConsume(queue, true, consumer);
        }

        public void DeleteQueues(params string[] queues)
        {
            foreach (string queue in queues)
            {
                this._channel.QueueDelete(queue);
            }
        }

        public void Publish(string queueName, string contentType, string payload)
        {
            byte[] messageBodyBytes = System.Text.Encoding.UTF8.GetBytes(payload);
            IBasicProperties props = this._channel.CreateBasicProperties();
            props.ContentType = contentType;
            props.DeliveryMode = 2;
            this._channel.BasicPublish("", queueName, props, messageBodyBytes);
        }
    }
}