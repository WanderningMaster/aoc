defmodule Misc.Mathutils do
  def divisors(n) when n > 0 do
    limit = :math.sqrt(n) |> trunc()

    divisors =
      for x <- 1..limit, rem(n, x) == 0 do
        y = div(n, x)
        if x == y, do: [x], else: [x, y]
      end

    divisors
    |> List.flatten()
    |> Enum.sort()
  end
end
