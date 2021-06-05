using System;
using System.Text;
using System.Threading;
using pulsar.clients;

namespace pulsar
{
    class Program
    {
        static void Main(string[] args)
        {
            Config config = new Config();


            AmqpClient amqpClient = new AmqpClient(config.GetAmqpuri()).Connect();
            
            amqpClient.DeclareQueues("nienna_jobs_result");

            amqpClient.AddConsumer("nienna_jobs_result", (sender, ea) =>
            {
                Console.WriteLine("EVENT: " + Encoding.UTF8.GetString(ea.Body.ToArray()));
            });
            
            while (true)
            {
                // noop
                Thread.Sleep(1000);
            }
        }
    }
}