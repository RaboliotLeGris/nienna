using System.Data;
using System.IO;
using Newtonsoft.Json;

namespace pulsar.events
{
    public static class EventParser
    {
        public static Event Parse(string payload)
        {
            var parsedEvent = JsonConvert.DeserializeObject<Event>(payload);

            if (parsedEvent.Type == "" )
            {
                throw new DataException("Missing mandatory field \"event\" in event");
            }
            if (parsedEvent.Slug == "")
            {
                throw new DataException("Missing mandatory field \"slug\" in event");
            }

            if (!IsValidEventType(parsedEvent))
            {
                throw new InvalidDataException("Event \"" + parsedEvent.Type + "\" is not a valid event type");
            }
            
            return parsedEvent;
        }
        
        public static bool IsValidEventType(Event e)
        {
            return e.Type switch
            {
                "EventVideoProcessingSucceed" => true,
                "EventVideoProcessingFail" => true,
                _ => false,
            };
        }
    }
}