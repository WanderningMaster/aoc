defmodule Aoc.Y2025.Day2.Part2 do
  defp invalid_id_enum_any_uniq?(id) when is_binary(id) do
    {_, divs} = Misc.Mathutils.divisors(String.length(id)) |> List.pop_at(-1)

    invalid? =
      Enum.any?(divs, fn x ->
        Enum.uniq(Misc.Strings.chunk_by(id, x)) |> length() == 1
      end)

    if invalid?, do: String.to_integer(id), else: 0
  end

  defp invalid_id?(id) when is_binary(id) do
    len = String.length(id)

    {_, divs} = Misc.Mathutils.divisors(len) |> List.pop_at(-1)

    invalid? =
      Enum.any?(divs, fn chunk_size ->
        chunk = :binary.part(id, 0, chunk_size)
        repeats = div(len, chunk_size)

        Enum.all?(1..(repeats - 1), fn k ->
          start = k * chunk_size
          :binary.part(id, start, chunk_size) == chunk
        end)
      end)

    if invalid?, do: String.to_integer(id), else: 0
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
