module Main where

import Data.Fixed
import Math.NumberTheory.Powers.Squares

trip a b = case exactSquareRoot (a^2+b^2) of
    Just c | a + b + c == 1000  ->  a * b * c
    _ | b == 1000          -> 0
      | a == b             -> trip 1 (b+1)
      | otherwise          -> trip (a+1) b

main = print (trip 3 4)
