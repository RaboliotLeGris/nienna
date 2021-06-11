using Newtonsoft.Json;

namespace pulsar.events
{
    
    /*
     * Potential Type value:
     * * EventVideoProcessingSucceed
     * * EventVideoProcessingFail
     * * EventVideoReadyForProcessing
     */
    public class Event
    {
        [JsonProperty("event")]
        public string Type { get; }
        [JsonProperty("slug")]
        public string Slug { get; }
        [JsonProperty("content")]
        public string Content { get; }
        
        public Event(string type, string slug, string content)
        {
            Type = type;
            Slug = slug;
            Content = content;
        }
    }
}