# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.
use Mix.Config

# Configures the router
config :phoenix, Batman.Router,
  url: [host: "localhost"],
  http: [port: System.get_env("PORT")],
  secret_key_base: "9EKliUb78P6QZcaQ4PoDbAUin9qBUAyfbt0cECWE4j2hHtK+4S9MlP106pgpKBHZov+VrgR4NOQJQKo3ibxniQ==",
  debug_errors: false,
  error_controller: Batman.PageController

# Session configuration
config :phoenix, Batman.Router,
  session: [store: :cookie,
            key: "_batman_key"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env}.exs"
