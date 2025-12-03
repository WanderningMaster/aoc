defmodule Aoc.Y2025.Day3.Part2 do
  require Logger

  def parse(line) do
    line
    |> String.split("", trim: true)
    |> Enum.map(&String.to_integer/1)
  end

  def arrange(bank) when is_list(bank) do
    bank
    |> Enum.with_index()
    |> arrange([], 0, 12)
    |> Enum.reverse()
    |> Enum.reduce(0, fn digit, acc -> acc * 10 + digit end)
  end

  def arrange(_, l, _, left) when left == 0, do: l

  def arrange(bank, l, offset, left) do
    w_end = length(bank) - left

    {nVal, idx} =
      bank
      |> Enum.slice(offset..w_end)
      |> Enum.max_by(fn {x, _} -> x end)

    arrange(bank, [nVal | l], idx + 1, left - 1)
  end

  def run do
    File.read!("lib/aoc/2025/day3/in.txt")
    |> String.split("\n", trim: true)
    |> Enum.map(&parse/1)
    |> Enum.map(&arrange/1)
    |> Enum.sum()
  end
end
