defmodule Mix.Tasks.Aoc do
  use Mix.Task

  @impl Mix.Task
  def run(args) do
    Mix.Task.run("app.start")

    case args do
      [year, day] ->
        run_parts(year, normalize_day(day), ["part1", "part2"]) |> exit_status()

      [year, day, part] ->
        parts = normalize_part(part)
        run_parts(year, normalize_day(day), parts) |> exit_status()

      _ ->
        Mix.shell().info(@moduledoc || "Usage: mix aoc <year> <day> [part]")
        Mix.raise("Usage: mix aoc <year> <day> [part]")
    end
  end

  defp run_parts(year, day, parts) do
    Enum.map(parts, fn part ->
      header = "==> #{year}/#{day}/#{part}"
      Mix.shell().info(header)

      mod = module_for(year, day, part)

      cond do
        not Code.ensure_loaded?(mod) ->
          Mix.shell().error("Missing module: #{inspect(mod)}")
          {:error, :missing_module}

        not function_exported?(mod, :run, 0) ->
          Mix.shell().error("Missing function run/0 in #{inspect(mod)}")
          {:error, :missing_function}

        true ->
          try do
            value = apply(mod, :run, [])
            Mix.shell().info(inspect(value))
            {:ok, value}
          rescue
            e ->
              Mix.shell().error(
                "Error running #{inspect(mod)}.run/0: " <>
                  Exception.format(:error, e, __STACKTRACE__)
              )

              {:error, e}
          end
      end
    end)
  end

  defp normalize_part(part) when is_binary(part) do
    case part do
      "1" -> ["part1"]
      "2" -> ["part2"]
      _ -> Mix.raise("Unknown part: #{part}. Use 1, 2")
    end
  end

  defp normalize_day(day) when is_binary(day) do
    "day#{day}"
  end

  defp module_for(year, day, part) do
    Module.concat([
      Aoc,
      String.to_atom("Y#{year}"),
      String.to_atom(String.capitalize(day)),
      String.to_atom(String.capitalize(part))
    ])
  end

  defp exit_status(results) do
    if Enum.any?(results, fn r -> match?({:error, _}, r) end) do
      :erlang.halt(1)
    else
      :ok
    end
  end
end
