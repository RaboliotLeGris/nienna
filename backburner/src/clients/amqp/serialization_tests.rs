#[cfg(test)]
mod serialization_tests {
    use crate::clients::amqp::serialization::EventSerialization;

    #[test]
    fn should_parse_json_to_event_serialization_struct() {
        // given
        let given = r#"
        {
            "event": "EventVideoReadyForProcessing",
            "slug": "randomstring",
            "filename": "randomefilename.mp4"
        }"#;

        // then
        let e: EventSerialization = serde_json::from_str(given).unwrap();

        // expect
        let expected = EventSerialization::new(String::from("EventVideoReadyForProcessing"), String::from("randomstring"), String::from("randomefilename.mp4"));

        assert_eq!(expected, e)
    }

    #[test]
    fn should_fail_to_parse_json_to_event_serialization_struct() {
        // given
        // Should fail because there is an extra comma and key should be with ""
        let given = r#"
        {
            event: "EventVideoReadyForProcessing",
            slug: "randomstring",
            filename: "filename.mp4",
        }"#;

        // then
        let e: serde_json::error::Result<EventSerialization> = serde_json::from_str(given);

        // expect
        assert!(e.is_err());
    }

    #[test]
    fn should_fail_due_to_unrecognized_event_struct() {
        // given
        let given = r#"
        {
            "event": "UnrecognizedEvent",
            "slug": "randomstring",
            "filename": "randomfilename.mp4"
        }"#;
        let e: serde_json::error::Result<EventSerialization> = serde_json::from_str(given);

        // then
        let res = EventSerialization::check_event(&e.unwrap());

        // expect
        assert_eq!(res, false, "test should fail, event is unrecognized")
    }
}