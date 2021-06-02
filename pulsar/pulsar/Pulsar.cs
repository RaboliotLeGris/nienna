using System;

namespace pulsar
{
    class Pulsar
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            Config config = new Config();
            Console.WriteLine(config.getDBURI());
        }
    }
}