module Problem1 where
import Data.List
import Text.Printf

main :: IO()
main = do
    rawInput <- readFile "input.txt"
    let rawInputLines = lines rawInput
    let inputLineNumbers = map getLineNumbers rawInputLines
    let inputLinesAsPairs = map arrayIntoPair inputLineNumbers
    let (listA, listB) = unzip inputLinesAsPairs

    let sortedListA = sort listA
    let sortedListB = sort listB
    let sortedListPairs = zip sortedListA sortedListB

    let distances = map (\pair -> abs $ snd pair - fst pair) sortedListPairs
    let totalDistance = sum distances

    let sameCounts = map (\numA -> length $ filter (== numA) sortedListB) sortedListA
    let similarityScores = zipWith (*) sortedListA sameCounts

    printf "Solution 1: %d\nSolution 2: %d\n" totalDistance (sum similarityScores)

arrayIntoPair :: [Int] -> (Int, Int)
arrayIntoPair arr = (head arr, arr !! 1)

getLineNumbers :: String -> [Int]
getLineNumbers str = map read $ words str