using System;
using RabbitMQ.Client.Events;

namespace pulsar.clients
{
    public interface IAmqpClient
    {
        public IAmqpClient Connect();
        public void DeclareQueues(params string[] queues);
        public string AddConsumer(string queue, EventHandler<BasicDeliverEventArgs> f);
        public void DeleteQueues(params string[] queues);
        public void Publish(string queueName, string contentType, string payload);
    }
}