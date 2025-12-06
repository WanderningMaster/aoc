defmodule Aoc.Y2025.Day6.Part1 do
  def do_math(group, op) do
    case op do
      "+" -> group |> Tuple.to_list() |> Enum.sum()
      "*" -> group |> Tuple.to_list() |> Enum.reduce(1, fn x, acc -> x * acc end)
    end
  end

  def run do
    {nums, [ops]} =
      File.read!("lib/aoc/2025/day6/in.txt")
      |> String.split("\n", trim: true)
      |> Enum.map(&String.split/1)
      |> Enum.map(fn x -> Enum.join(x, " ") end)
      |> Enum.split(-1)

    ops = ops |> String.split(" ", trim: true)

    nums =
      nums
      |> Enum.map(fn x ->
        String.split(x, " ", trim: true)
        |> Enum.map(&String.to_integer/1)
      end)

    nums
    |> Enum.zip()
    |> Enum.with_index()
    |> Enum.map(fn {group, idx} -> do_math(group, Enum.at(ops, idx)) end)
    |> Enum.sum()
  end
end
