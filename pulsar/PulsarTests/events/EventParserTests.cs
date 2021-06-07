using System;
using System.Data;
using System.IO;
using NUnit.Framework;
using pulsar.events;

namespace PulsarTests.events
{
    public class EventParserTests
    {
        [Test]
        public void MustParseVideoProcessingSucceed()
        {
            string payload = @"
                {
                    ""event"": ""EventVideoProcessingSucceed"",
                    ""slug"": ""SomeSlug"",
                    ""content"": """"
                }
            ";
            Event expected = new Event("EventVideoProcessingSucceed", "SomeSlug", "");
            var result = EventParser.Parse(payload);
            Assert.That(result.Type, Is.EqualTo(expected.Type));
            Assert.That(result.Slug, Is.EqualTo(expected.Slug));
            Assert.That(result.Content, Is.EqualTo(expected.Content));
        }
        
        [Test]
        public void MustParseVideoProcessingFail()
        {
            string payload = @"
                {
                    ""event"": ""EventVideoProcessingFail"",
                    ""slug"": ""SomeSlug"",
                    ""content"": ""Fall on the ground""
                }
            ";
            Event expected = new Event("EventVideoProcessingFail", "SomeSlug", "Fall on the ground");
            
            var result = EventParser.Parse(payload);
            
            Assert.That(result.Type, Is.EqualTo(expected.Type));
            Assert.That(result.Slug, Is.EqualTo(expected.Slug));
            Assert.That(result.Content, Is.EqualTo(expected.Content));
        }
        
        [Test]
        public void InvalidJsonMustThrow()
        {
            string payload = @"
                {
                    ""event"": ""EventVideoProcessingFail"",
                    ""slug"": ""SomeSlug"",
                    ""content"": ""Fall on the ground"",somestuffthatshouldbehere
                }
            ";
            
            var ex = Assert.Throws<Newtonsoft.Json.JsonReaderException>(delegate { EventParser.Parse(payload); });
            Assert.True(ex.Message.StartsWith("Invalid character after parsing property name"));
        }
        
        [Test]
        public void MissingMandatoryEventField()
        {
            string payload = @"
                {
                    ""event"": """",
                    ""slug"": ""SomeSlug"",
                    ""content"": ""Fall on the ground""
                }
            ";
            
            var ex = Assert.Throws<DataException>(delegate { EventParser.Parse(payload); });
            Assert.True(ex.Message.StartsWith("Missing mandatory field \"event\" in event"));
        }
        
        [Test]
        public void MissingMandatorySlugField()
        {
            string payload = @"
                {
                    ""event"": ""EventVideoProcessingFail"",
                    ""slug"": """",
                    ""content"": ""Fall on the ground""
                }
            ";
            
            var ex = Assert.Throws<DataException>(delegate { EventParser.Parse(payload); });
            Assert.True(ex.Message.StartsWith("Missing mandatory field \"slug\" in event"));
        }
        
        [Test]
        public void InvalidTypField()
        {
            string payload = @"
                {
                    ""event"": ""SomeRandomEventType"",
                    ""slug"": ""SomeSlug"",
                    ""content"": ""Fall on the ground""
                }
            ";
            
            var ex = Assert.Throws<InvalidDataException>(delegate { EventParser.Parse(payload); });
            Assert.That(ex.Message, Is.EqualTo("Event \"SomeRandomEventType\" is not a valid event type"));
        }
        
    }
}