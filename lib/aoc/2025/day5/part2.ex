defmodule Aoc.Y2025.Day5.Part2 do
  def parse(range) do
    [a, b] = range |> String.split("-") |> Enum.map(&String.to_integer/1)
    a..b
  end

  def collapse_ranges(ranges) do
    ranges
    |> Enum.map(fn x -> {x.first, x.last} end)
    |> Enum.sort()
    |> Enum.reduce([], fn {f, l}, acc ->
      case acc do
        [] ->
          [{f, l}]

        [{af, al} | rest] ->
          if f <= al + 1 do
            [{af, max(al, l)} | rest]
          else
            [{f, l} | acc]
          end
      end
    end)
  end

  def run do
    [ranges, _] =
      File.read!("lib/aoc/2025/day5/in.txt")
      |> String.split("\n\n")

    ranges
    |> String.split("\n", trim: true)
    |> Enum.map(&parse/1)
    |> collapse_ranges()
    |> Enum.reduce(0, fn {a, b}, acc ->
      acc + (b - a + 1)
    end)
  end
end
