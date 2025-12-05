defmodule Aoc.Y2025.Day5.Part1 do
  defmacro load_ranges!(filename) do
    [ranges, _] =
      filename
      |> File.read!()
      |> String.split("\n\n")

    ranges =
      ranges
      |> String.split("\n", trim: true)
      |> Enum.map(fn line ->
        [a, b] = line |> String.split("-") |> Enum.map(&String.to_integer/1)

        a..b
      end)

    v = Macro.var(:v, nil)
    x = Macro.var(:x, nil)

    clauses =
      Enum.map(ranges, fn range ->
        escaped_range = Macro.escape(range)
        in_call = {:in, [], [v, escaped_range]}
        when_ast = {:when, [], [v, in_call]}
        {:->, [], [[when_ast], true]}
      end)

    default_clause =
      {:->, [], [[{:_, [], Elixir}], false]}

    case_ast =
      {:case, [],
       [
         x,
         [do: clauses ++ [default_clause]]
       ]}

    fn_ast =
      {:fn, [],
       [
         {:->, [], [[x], case_ast]}
       ]}

    fn_ast
  end

  def run do
    [_, values] =
      File.read!("lib/aoc/2025/day5/in.txt")
      |> String.split("\n\n")

    values
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.map(load_ranges!("lib/aoc/2025/day5/in.txt"))
    |> Enum.count(fn x -> x != false end)
  end
end
