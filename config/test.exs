use Mix.Config

config :phoenix, Batman.Router,
  http: [port: System.get_env("PORT") || 4001],
