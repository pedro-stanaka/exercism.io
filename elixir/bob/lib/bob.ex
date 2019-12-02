defmodule Bob do
  def hey(input) do
    input = String.trim(input)

    case shouting?(input) do
      true -> answer_shouting(input)
      false -> answer_normal_dialogue(input)
    end
  end

  defp shouting?(input) do
    upcase_input = String.upcase(input)
    diff = String.myers_difference(input, upcase_input)

    cond do
      is_numeric_string(input) ->
        false

      true ->
        case Keyword.fetch(diff, :del) do
          {:ok, _} -> false
          :error -> true
        end
    end
  end

  def is_numeric_string(input) do
    clean_input = String.replace(input, ~r/[[:punct:][:blank:]]/, "")

    case Integer.parse(clean_input) do
      :error -> false
      {_, sufix} -> String.length(sufix) == 0
    end
  end

  defp answer_shouting(input) do
    cond do
      input == "" -> "Fine. Be that way!"
      String.ends_with?(input, "?") -> "Calm down, I know what I'm doing!"
      true -> "Whoa, chill out!"
    end
  end

  defp answer_normal_dialogue(input) do
    cond do
      input == "" -> "Fine. Be that way!"
      String.ends_with?(input, "?") -> "Sure."
      true -> "Whatever."
    end
  end
end
