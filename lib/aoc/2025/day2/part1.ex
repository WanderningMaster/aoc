defmodule Aoc.Y2025.Day2.Part1 do
  def invalid_id?(id) when is_binary(id) and rem(byte_size(id), 2) == 1, do: 0

  def invalid_id?(id) when is_binary(id) do
    invalid_id?(String.split_at(id, 1), 1)
  end

  def invalid_id?({_, ""}, _), do: 0

  def invalid_id?({x, x}, _) do
    String.to_integer(x <> x)
  end

  def invalid_id?({id, rest}, at) when is_number(at) do
    invalid_id?(String.split_at(id <> rest, at + 1), at + 1)
  end

  def parse(range) do
    [s, e] = String.split(range, "-")
    String.to_integer(s)..String.to_integer(e)
  end

  def run do
    File.read!("lib/aoc/2025/day2/in.txt")
    |> String.split([",", "\n"], trim: true)
    |> Enum.map(&parse/1)
    |> Enum.reduce(0, fn x, acc ->
      acc +
        Enum.reduce(x, 0, fn n, r_acc ->
          r_acc + invalid_id?(Integer.to_string(n))
        end)
    end)
  end
end
