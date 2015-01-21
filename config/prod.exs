use Mix.Config

# ## SSL Support
#
# To get SSL working, you will need to set:
#
#     https: [port: 443,
#             keyfile: System.get_env("SOME_APP_SSL_KEY_PATH"),
#             certfile: System.get_env("SOME_APP_SSL_CERT_PATH")]
#
# Where those two env variables point to a file on
# disk for the key and cert.

config :phoenix, Batman.Router,
  url: [host: "example.com"],
  http: [port: System.get_env("PORT")],
  secret_key_base: "9EKliUb78P6QZcaQ4PoDbAUin9qBUAyfbt0cECWE4j2hHtK+4S9MlP106pgpKBHZov+VrgR4NOQJQKo3ibxniQ=="

config :logger,
  level: :info
