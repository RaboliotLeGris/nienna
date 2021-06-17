using System;
using System.IO;

namespace PulsarTests.helpers
{
    public class Schema
    {
        public static string Get()
        {
            string schema = Path.Combine(Directory.GetCurrentDirectory(), "../../../helpers/schema.sql");
            string text = File.ReadAllText(schema);
            return text;
        }
    }
}