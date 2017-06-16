## REST API for parsing image boards like 4chan/2ch

### Why?
I have written this as I love .webms, .gifs etc and everytime I should
manually download them.

This REST API parse your thread URL and return a JSON with all webm files
related to thread.

To obtain webm links you should make request to url

`http://goarparser.com/parse_data`

with parameter

```json
{
    "thread_link": "https://2ch.hk/{board_name}/res/{thread_id}.json"
}
```

Response:

```json
{
    "Files": [
        {
            "name": "{webm_id}.webm",
            "path": "....webm",
            "thumbnail": "....jpg"
        },
        ...
}
```