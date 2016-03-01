defmodule Bob do
  def hey(input) do
    cond do
      Regex.match?(~r/^\s*$/, input) -> "Fine. Be that way!"
      String.ends_with?(input, "?") -> "Sure."
      String.upcase(input) == input and Regex.match?(~r/\p{L}/, input) -> "Whoa, chill out!"
      true -> "Whatever."
    end
  end
end
