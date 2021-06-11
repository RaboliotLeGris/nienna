using System;
using System.Text;
using System.Threading;
using pulsar.clients;
using pulsar.daos;
using pulsar.events;

namespace pulsar
{
    class Program
    {
        static void Main(string[] args)
        {
            Config config = new Config();

            Console.WriteLine("Pulsar starting");
            if (config.GetLogLevel() == "DEBUG")
            {
                Console.WriteLine("CONFIG - " + config.GetAmqpuri() + " | " + config.GetDbUri() + " | " + config.GetLogLevel());
            }
                              

            AmqpClient amqpClient = new AmqpClient(config.GetAmqpuri());
            ISqlClient sqlClient = new SqlClient(config.GetDbUri()).Connect();

            Launch(config, amqpClient, sqlClient, Loop);
        }

        static void Launch(Config config, IAmqpClient amqpClient, ISqlClient sqlClient, Action loop)
        {
            amqpClient.Connect();
            amqpClient.DeclareQueues("nienna_jobs_result");
            amqpClient.AddConsumer("nienna_jobs_result", (sender, ea) =>
            {
                try
                {
                    Event e = EventParser.Parse(Encoding.UTF8.GetString(ea.Body.ToArray()));
                    VideoDao videoDao = new VideoDao(sqlClient);
                    switch (e.Type)
                    {
                        case "EventVideoProcessingSucceed":
                            videoDao.UpdateStatus(e.Slug, "READY");
                            break;
                        case "EventVideoProcessingFail":
                            videoDao.UpdateStatus(e.Slug, "FAILURE");
                            break;
                        default:
                            Console.WriteLine("Unrecognized event " + e.Type);
                            break;
                    }
                }
                catch (Exception exception)
                {
                    Console.WriteLine("Got exception while handling events: " + exception);
                }
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