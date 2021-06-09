using System;
using System.Text;
using System.Threading;
using pulsar.clients;
using pulsar.events;

namespace pulsar
{
    class Program
    {
        static void Main(string[] args)
        {
            Config config = new Config();
            
            Console.WriteLine("Pulsar starting");
            Console.WriteLine("CONFIG - " + config.GetAmqpuri() + " | " + config.GetDbUri() + " | " + config.GetLogLevel());

            AmqpClient amqpClient = new AmqpClient(config.GetAmqpuri());
            
            Launch(config, amqpClient, null, Loop);
        }

        static void Launch(Config config, IAmqpClient amqpClient, ISqlClient sqlClient, Action loop)
        {
            amqpClient.Connect();
            amqpClient.DeclareQueues("nienna_jobs_result");
            amqpClient.AddConsumer("nienna_jobs_result", (sender, ea) =>
            {
                Event e = EventParser.Parse(Encoding.UTF8.GetString(ea.Body.ToArray()));
                Console.WriteLine("EVENT: " + e.Type);
            });

            loop();
        }

        static void Loop()
        {
            while (true)
            {
                // noop
                Thread.Sleep(1000);
            }
        }
    }
}