defmodule Misc.Strings do
  def chunk_by(str, size)
      when is_binary(str) and
             is_number(size) do
    str
    |> String.codepoints()
    |> Enum.chunk_every(size)
    |> Enum.map(&Enum.join/1)
  end
end
