package es

var mappings = `
{
    "properties": {
        "title": {
            "type": "string",
            "index": "not_analyzed"
        },
        "author_url": {
            "type": "string",
            "index": "not_analyzed"
        },
        "author": {
            "type": "string"
        },
        "reply": {
            "type": "integer"
        },
        "last_reply_time": {
            "type": "date"
        },
        "topic_content": {
            "properties": {
                "update_time": {
                    "type": "date"
                },
                "content": {
                    "type": "string",
                    "index": "not_analyzed"
                },
                "with_pic": {
                    "type": "boolean"
                },
                "pic_urls": {
                    "type": "string",
                    "index": "not_analyzed"
                },
                "like": {
                    "type": "integer"
                }
            }
        }
    }
}
`
