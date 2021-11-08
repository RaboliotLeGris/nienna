using System;
using NUnit.Framework;
using pulsar;

namespace PulsarTests
{
    public class ConfigTests
    {
        [TearDown]
        public void Clean()
        {
            Environment.SetEnvironmentVariable("NIENNA_LOG_LEVEL", null);
            Environment.SetEnvironmentVariable("DB_PARAMS", null);
            Environment.SetEnvironmentVariable("AMQP_URI", null);
        }

        [Test]
        public void AllEnvVarSet()
        {
            Environment.SetEnvironmentVariable("NIENNA_LOG_LEVEL", "log_level");
            Environment.SetEnvironmentVariable("DB_PARAMS", "db_uri");
            Environment.SetEnvironmentVariable("AMQP_URI", "amqp_uri");

            Config config = new Config();

            Assert.That(config.GetLogLevel(), Is.EqualTo("log_level"));
            Assert.That(config.GetDbUri(), Is.EqualTo("db_uri"));
            Assert.That(config.GetAmqpuri(), Is.EqualTo("amqp_uri"));
        }

        [Test]
        public void MissingAmqpUri()
        {
            Environment.SetEnvironmentVariable("NIENNA_LOG_LEVEL", "log_level");
            Environment.SetEnvironmentVariable("DB_PARAMS", "db_uri");

            var ex = Assert.Throws<ArgumentException>(delegate { new Config(); });
            Assert.That(ex.Message, Is.EqualTo("Env var AMQP_URI must not be null"));
        }

        [Test]
        public void MissingDbUri()
        {
            Environment.SetEnvironmentVariable("NIENNA_LOG_LEVEL", "log_level");

            var ex = Assert.Throws<ArgumentException>(delegate { new Config(); });
            Assert.That(ex.Message, Is.EqualTo("Env var DB_URI must not be null"));

        }

        [Test]
        public void MissingLogLevel()
        {
            var ex = Assert.Throws<ArgumentException>(delegate { new Config(); });
            Assert.That(ex.Message, Is.EqualTo("Env var NIENNA_LOG_LEVEL must not be null"));
        }
    }
}