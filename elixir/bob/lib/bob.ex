defmodule Bob do
  def hey(input) do
    cond do
      silence?(input) -> "Fine. Be that way!"
      question?(input) -> "Sure."
      shouting_question?(input) -> "Calm down, I know what I am doing!"
      shouting?(input) -> "Whoa, chill out!"
      true -> "Whatever."
    end
  end

  defp silence?(input) do
    String.trim(input) == ""
  end

  defp question?(input) do
    String.ends_with?(input, "?")
  end

  defp shouting_question?(input) do
    question?(input) and shouting?(input)
  end

  defp shouting?(input) do
    input == String.upcase(input)
  end
end
