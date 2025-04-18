/**
PROMPT FOR CURSOR:

Could you provide a simple example of using the golang standard library and its http client to call an api endpoint with post data/json.

Do not create a file, do not modify or add code to any file.
Please do not produce any diffs.

Just provide a simple example so I can see how the http client in the golang std library works.

Please use this api spec to create the example:
curl https://api.anthropic.com/v1/messages \
     --header "x-api-key: $ANTHROPIC_API_KEY" \
     --header "anthropic-version: 2023-06-01" \
     --header "content-type: application/json" \
     --data \
'{
    "model": "claude-3-7-sonnet-20250219",
    "max_tokens": 1024,
    "messages": [
        {"role": "user", "content": "Hello, world"}
    ]
}'
**/


