defmodule Aoc.Y2025.Day3.Part1 do
  require Logger

  def parse(line) do
    line |> String.split("", trim: true)
  end

  def arrange(bank) when is_list(bank) do
    {v, idx} =
      bank
      |> Enum.with_index()
      |> Enum.max_by(fn {x, _idx} -> String.to_integer(x) end)

    {v1, idx1} =
      case Enum.split(bank, idx) do
        {_, m} when length(m) > 1 ->
          m
          |> Enum.with_index(idx)
          |> Enum.drop(1)
          |> Enum.max_by(fn {x, _idx} -> String.to_integer(x) end)

        {s, m} when length(m) <= 1 ->
          s
          |> Enum.with_index()
          |> Enum.max_by(fn {x, _idx} -> String.to_integer(x) end)
      end

    if idx > idx1, do: String.to_integer(v1 <> v), else: String.to_integer(v <> v1)
  end

  def run do
    File.read!("lib/aoc/2025/day3/in.txt")
    |> String.split("\n", trim: true)
    |> Enum.map(&parse/1)
    |> Enum.map(&arrange/1)
    |> Enum.sum()
  end
end
