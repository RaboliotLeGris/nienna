using System;
using System.Text;
using System.Threading;
using NUnit.Framework;
using pulsar.clients;
using RabbitMQ.Client.Exceptions;

namespace PulsarTests.clients
{
    public class AmqpClientTests
    {
        private static string _amqpUri;

        [OneTimeSetUp]
        public void SetAmqpUri()
        {
            _amqpUri = Environment.GetEnvironmentVariable("TEST_IN_DOCKER") != null
                ? "amqp://nienna:nienna123@rabbitmq:5672"
                : "amqp://nienna:nienna123@localhost:5672";
            Console.WriteLine();
        }

        [Test]
        public void WrongUriMustThrow()
        {
            var ex = Assert.Throws<ArgumentException>(delegate { new AmqpClient("Some wrong uri"); });
            Assert.That(ex.Message, Is.EqualTo("Invalid AMQP URI: Some wrong uri"));
        }

        [Test]
        public void MustConnectToAmqpServer()
        {
            var amqpClient = new AmqpClient(_amqpUri);
            amqpClient.Connect();
            Assert.Pass();
        }

        [Test]
        public void MustFailConnectToAmqpServer()
        {
            var amqpClient = new AmqpClient("amqp://raboland/fezf:fezef");
            var ex = Assert.Throws<BrokerUnreachableException>(delegate { amqpClient.Connect(); });
            Assert.That(ex.Message, Is.EqualTo("None of the specified endpoints were reachable"));
        }

        [Test]
        public void DeclareQueuesMustSucceed()
        {
            var amqpClient = new AmqpClient(_amqpUri);
            amqpClient.Connect();
            amqpClient.DeclareQueues("test_pulsar1", "test_pulsar2");
            amqpClient.DeleteQueues("test_pulsar1", "test_pulsar2");
            Assert.Pass();
        }

        [Test]
        public void PublishMustSucceed()
        {
            var amqpClient = new AmqpClient(_amqpUri);
            amqpClient.Connect();
            amqpClient.DeclareQueues("test_pulsar");
            amqpClient.Publish("test_pulsar", "text/plain", "some data");
            amqpClient.DeleteQueues("test_pulsar");
        }

        [Test]
        public void AddConsumerMustSucceed()
        {
            var amqpClient = new AmqpClient(_amqpUri);
            amqpClient.Connect();
            amqpClient.DeclareQueues("test_pulsar");

            var gotEvent = false;

            amqpClient.AddConsumer("test_pulsar", (sender, ea) =>
            {
                var event_payload = Encoding.UTF8.GetString(ea.Body.ToArray());
                Console.WriteLine("EVENT: " + event_payload);
                if (event_payload == "some data")
                {
                    gotEvent = true;
                }
            });
            amqpClient.Publish("test_pulsar", "text/plain", "some data");
            Thread.Sleep(500);
            Assert.IsTrue(gotEvent);
        }
    }
}