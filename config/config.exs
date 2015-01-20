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
  secret_key_base: "PSerot/o4yVdegbNz81qfJLD2glw4bZxuC58NmtAi3fOWbCMxZBxOWt+cc5/ERcSOwjWOELLI/K4KxmFDH4LZg==",
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
