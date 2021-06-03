using System;

namespace pulsar
{
    class Program
    {
        static void Main(string[] args)
        {
            Config config = new Config();
            Console.WriteLine(config.getDBURI());
            while (true)
            {
                // noop
            }
        }
    }
}