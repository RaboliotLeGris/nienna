using System;

namespace pulsar
{
    public class Config
    {
        private const string LOG_LEVEL_KEY = "NIENNA_LOG_LEVEL";
        private const string DB_PARAMS_KEY = "DB_PARAMS";
        private const string AMQP_URI_KEY = "AMQP_URI";
        
        private string logLevel;
        private string dbURI;
        private string amqpURI;


        public Config()
        {
            this.logLevel = Environment.GetEnvironmentVariable(LOG_LEVEL_KEY);
            if (this.logLevel == null)
            {
                throw new ArgumentException("Env var NIENNA_LOG_LEVEL must not be null");
            }
            this.dbURI = Environment.GetEnvironmentVariable(DB_PARAMS_KEY);
            if (this.dbURI == null)
            {
                throw new ArgumentException("Env var DB_URI must not be null");
            }
            this.amqpURI = Environment.GetEnvironmentVariable(AMQP_URI_KEY);
            if (this.amqpURI == null)
            {
                throw new ArgumentException("Env var AMQP_URI must not be null");
            }
        }

        public string GetLogLevel()
        {
            return this.logLevel;
        }

        public string GetDbUri()
        {
            return this.dbURI;
        }

        public string GetAmqpuri()
        {
            return this.amqpURI;
        }
    }
}