defmodule Aoc.Y2025.Day1.Part2 do
  def wrap(n, step) do
    full_rotations = abs(div(step, 100))
    leftover = rem(abs(step), 100)

    zero_dist = if step > 0, do: 100 - n, else: n

    extra =
      if zero_dist > 0 and zero_dist <= leftover, do: 1, else: 0

    r = rem(n + step, 100)
    wrapped = if r < 0, do: r + 100, else: r

    # IO.puts(
    #   "#{n}+#{step} -> #{wrapped} zero_dist=#{zero_dist} leftover=#{leftover} extra=#{extra} full=#{full_rotations} "
    # )

    {full_rotations + extra, wrapped}
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
        {w, acc} = wrap(acc, x)
        {acc, c + w}
      end)

    answer
  end
end
