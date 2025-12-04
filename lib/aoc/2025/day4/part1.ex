defmodule Aoc.Y2025.Day4.Part1 do
  def get_adjacent_cells(matrix, {row, col}) do
    num_rows = length(matrix)
    num_cols = length(hd(matrix))

    directions = [
      {-1, 0},
      {1, 0},
      {0, -1},
      {0, 1},
      {-1, -1},
      {-1, 1},
      {1, -1},
      {1, 1}
    ]

    is_valid = fn r, c -> r >= 0 && r < num_rows && c >= 0 && c < num_cols end

    Enum.flat_map(directions, fn {dr, dc} ->
      new_row = row + dr
      new_col = col + dc

      if is_valid.(new_row, new_col) do
        value = Enum.at(Enum.at(matrix, new_row), new_col)
        [{value, {new_row, new_col}}]
      else
        []
      end
    end)
  end

  def parse(row) do
    row |> String.split("", trim: true)
  end

  def run do
    matrix =
      File.read!("lib/aoc/2025/day4/in.txt")
      |> String.split("\n", trim: true)
      |> Enum.map(&parse/1)

    {_, count} =
      Enum.reduce(0..(length(matrix) - 1), {matrix, 0}, fn row, {matrix_acc, count_acc} ->
        Enum.reduce(0..(length(hd(matrix)) - 1), {matrix_acc, count_acc}, fn col,
                                                                             {acc, acc_count} ->
          curr_cell = Enum.at(Enum.at(matrix_acc, row), col)

          case curr_cell do
            "@" ->
              nearby =
                get_adjacent_cells(matrix, {row, col})
                |> Enum.filter(fn {v, _} -> v == "@" end)
                |> Enum.count()

              if nearby < 4 do
                updated =
                  List.update_at(acc, row, fn r ->
                    List.update_at(r, col, fn _ -> "x" end)
                  end)

                {updated, acc_count + 1}
              else
                {acc, acc_count}
              end

            _ ->
              {acc, acc_count}
          end
        end)
      end)

    count
  end
end
