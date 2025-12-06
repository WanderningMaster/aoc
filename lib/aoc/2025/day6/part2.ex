defmodule Aoc.Y2025.Day6.Part2 do
  def do_math(group, op) do
    case op do
      "+" -> group |> Enum.sum()
      "*" -> group |> Enum.reduce(1, fn x, acc -> x * acc end)
    end
  end

  def parse(lines) do
    rows = lines |> Enum.map(&String.graphemes/1)

    rows
    |> Enum.zip()
    |> Enum.map(&Tuple.to_list/1)
    |> Enum.map(fn col ->
      col
      |> Enum.join("")
      |> String.trim()
      |> String.replace(" ", "")
    end)
    |> Enum.reduce([[]], fn
      "", acc -> [[] | acc]
      item, [head | tail] -> [[item | head] | tail]
    end)
    |> Enum.map(fn x -> x |> Enum.map(&String.to_integer/1) end)
    |> Enum.reverse()
  end

  def run do
    {nums, [ops]} =
      File.read!("lib/aoc/2025/day6/in.txt")
      |> String.split("\n", trim: true)
      |> Enum.split(-1)

    ops = ops |> String.split(" ", trim: true)

    nums
    |> parse()
    |> Enum.with_index()
    |> Enum.map(fn {group, idx} -> do_math(group, Enum.at(ops, idx)) end)
  end
end
