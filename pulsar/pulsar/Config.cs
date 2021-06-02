using System;

namespace pulsar
{
    public class Config
    {
        private const string LOG_LEVEL_KEY = "LOG_LEVEL";
        private const string DB_URI_KEY = "DB_URI";
        private const string AMQP_URI_KEY = "AMQP_URI";
        
        private String logLevel;
        private String dbURI;
        private String amqpURI;


        public Config()
        {
            this.logLevel = Environment.GetEnvironmentVariable(LOG_LEVEL_KEY);
            if (this.logLevel == null)
            {
                throw new ArgumentException("Env var LOG_LEVEL must not be null");
            }
            this.dbURI = Environment.GetEnvironmentVariable(DB_URI_KEY);
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

        public String getLogLevel()
        {
            return this.logLevel;
        }

        public String getDBURI()
        {
            return this.dbURI;
        }

        public String getAMQPURI()
        {
            return this.amqpURI;
        }
    }
}