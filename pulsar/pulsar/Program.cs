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
            
            AmqpClient amqpClient = new AmqpClient(config.GetAmqpuri()).Connect();
            amqpClient.DeclareQueues("nienna_jobs_result");
            amqpClient.AddConsumer("nienna_jobs_result", (sender, ea) =>
            {
                Event e = EventParser.Parse(Encoding.UTF8.GetString(ea.Body.ToArray()));
                Console.WriteLine("EVENT: " + e.Type);
            });
            
            while (true)
            {
                // noop
                Thread.Sleep(1000);
            }
        }
    }
}