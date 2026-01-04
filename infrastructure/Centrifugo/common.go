package centrifugo

//──────────────────────────────────────────────────────────────────────────────────────────────────

const _RETRY_POLICY string = `{
    "methodConfig": [{
        "waitForReady": true,
        "retryPolicy": {
            "maxAttempts": 4,
            "initialBackoff": "0.1s",
            "maxBackoff": "1s",
            "backoffMultiplier": 2.0,
            "retryableStatusCodes": ["UNAVAILABLE", "RESOURCE_EXHAUSTED"]
        }
    }]
}`
