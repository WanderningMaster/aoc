defmodule Aoc.Y2025.Day1.Part1 do
  def wrap(n, step) do
    r = rem(n + step, 100)
    if r < 0, do: r + 100, else: r
  end

  def parse(line) do
    {dir, astr} = String.next_grapheme(line)

    with {amount, _} <- Integer.parse(astr) do
      case dir do
        "L" -> -amount
        "R" -> amount
      end
    else
      :error -> {:error, :not_an_integer}
    end
  end

  def run do
    ptr = 50

    {_, answer} =
      File.read!("lib/aoc/2025/day1/in.txt")
      |> String.split(["\r\n", "\n"], trim: true)
      |> Enum.map(&parse/1)
      |> Enum.reduce({ptr, 0}, fn x, {acc, c} ->
        acc = wrap(acc, x)
        c = if acc == 0, do: c + 1, else: c

        {acc, c}
      end)

    answer
  end
end
