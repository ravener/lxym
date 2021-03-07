
package pages

const Purescript = "PureScript is a small strongly, statically typed language compiling to Javascript.\n\n* Learn more at [http://www.purescript.org/](http://www.purescript.org/)\n* Documentation: [http://pursuit.purescript.org/](http://pursuit.purescript.org/)\n* Book: Purescript by Example, [https://leanpub.com/purescript/](https://leanpub.com/purescript/)\n\nAll the noncommented lines of code can be run in the PSCI REPL, though some will\nrequire the `--multi-line-mode` flag.\n\n```haskell\n\n--\n-- 1. Primitive datatypes that corresponds to their Javascript\n-- equivalents at runtime.\n\nimport Prelude\n-- Numbers\n1.0 + 7.2*5.5 :: Number -- 40.6\n-- Ints\n1 + 2*5 :: Int -- 11\n-- Types are inferred, so the following works fine\n9.0/2.5 + 4.4 -- 8.0\n-- But Ints and Numbers don't mix, so the following won't\n5/2 + 2.5 -- Expression 2.5 does not have type Int\n-- Hexadecimal literals\n0xff + 1 -- 256\n-- Unary negation\n6 * -3 -- -18\n6 * negate 3 -- -18\n-- Modulus, from purescript-math (Math)\n3.0 % 2.0 -- 1.0\n4.0 % 2.0 -- 0.0\n-- Inspect the type of an expression in psci\n:t 9.5/2.5 + 4.4 -- Prim.Number\n\n-- Booleans\ntrue :: Boolean -- true\nfalse :: Boolean -- false\n-- Negation\nnot true -- false\n23 == 23 -- true\n1 /= 4 -- true\n1 >= 4 -- false\n-- Comparisons < <= > >=\n-- are defined in terms of compare\ncompare 1 2 -- LT\ncompare 2 2 -- EQ\ncompare 3 2 -- GT\n-- Conjunction and Disjunction\ntrue && (9 >= 19 || 1 < 2) -- true\n\n-- Strings\n\"Hellow\" :: String -- \"Hellow\"\n-- Multiline string without newlines, to run in psci use the --multi-line-mode flag\n\"Hellow\\\n\\orld\" -- \"Helloworld\"\n-- Multiline string with newlines\n\"\"\"Hello\nworld\"\"\" -- \"Hello\\nworld\"\n-- Concatenate\n\"such \" <> \"amaze\" -- \"such amaze\"\n\n--\n-- 2. Arrays are Javascript arrays, but must be homogeneous\n\n[1,1,2,3,5,8] :: Array Int -- [1,1,2,3,5,8]\n[1.2,2.0,3.14] :: Array Number -- [1.2,2.0,3.14]\n[true, true, false] :: Array Boolean -- [true,true,false]\n-- [1,2, true, \"false\"] won't work\n-- `Cannot unify Prim.Int with Prim.Boolean`\n-- Cons (prepend)\n1 : [2,4,3] -- [1,2,4,3]\n\n-- Requires purescript-arrays (Data.Array)\n-- and purescript-maybe (Data.Maybe)\n\n-- Safe access return Maybe a\nhead [1,2,3] -- Just (1)\ntail [3,2,1] -- Just ([2,1])\ninit [1,2,3] -- Just ([1,2])\nlast [3,2,1] -- Just (1)\n-- Array access - indexing\n[3,4,5,6,7] !! 2 -- Just (5)\n-- Range\n1..5 -- [1,2,3,4,5]\nlength [2,2,2] -- 3\ndrop 3 [5,4,3,2,1] -- [2,1]\ntake 3 [5,4,3,2,1] -- [5,4,3]\nappend [1,2,3] [4,5,6] -- [1,2,3,4,5,6]\n\n--\n-- 3. Records are Javascript objects, with zero or more fields, which\n-- can have different types.\n-- In psci you have to write `let` in front of the function to get a\n-- top level binding.\nlet book = {title: \"Foucault's pendulum\", author: \"Umberto Eco\"}\n-- Access properties\nbook.title -- \"Foucault's pendulum\"\n\nlet getTitle b = b.title\n-- Works on all records with a title (but doesn't require any other field)\ngetTitle book -- \"Foucault's pendulum\"\ngetTitle {title: \"Weekend in Monaco\", artist: \"The Rippingtons\"} -- \"Weekend in Monaco\"\n-- Can use underscores as shorthand\n_.title book -- \"Foucault's pendulum\"\n-- Update a record\nlet changeTitle b t = b {title = t}\ngetTitle (changeTitle book \"Ill nome della rosa\") -- \"Ill nome della rosa\"\n\n--\n-- 4. Functions\n-- In psci's multiline mode\nlet sumOfSquares :: Int -> Int -> Int\n    sumOfSquares x y = x*x + y*y\nsumOfSquares 3 4 -- 25\nlet myMod x y = x % y\nmyMod 3.0 2.0 -- 1.0\n-- Infix application of function\n3 `mod` 2 -- 1\n\n-- function application has higher precedence than all other\n-- operators\nsumOfSquares 3 4 * sumOfSquares 4 5 -- 1025\n\n-- Conditional\nlet abs' n = if n>=0 then n else -n\nabs' (-3) -- 3\n\n-- Guarded equations\nlet abs'' n | n >= 0    = n\n            | otherwise = -n\n\n-- Pattern matching\n\n-- Note the type signature, input is a list of numbers. The pattern matching\n-- destructures and binds the list into parts.\n-- Requires purescript-lists (Data.List)\nlet first :: forall a. List a -> a\n    first (Cons x _) = x\nfirst (toList [3,4,5]) -- 3\nlet second :: forall a. List a -> a\n    second (Cons _ (Cons y _)) = y\nsecond (toList [3,4,5]) -- 4\nlet sumTwo :: List Int -> List Int\n    sumTwo (Cons x (Cons y rest)) = x + y : rest\nfromList (sumTwo (toList [2,3,4,5,6])) :: Array Int -- [5,4,5,6]\n\n-- sumTwo doesn't handle when the list is empty or there's only one element in\n-- which case you get an error.\nsumTwo [1] -- Failed pattern match\n\n-- Complementing patterns to match\n-- Good ol' Fibonacci\nlet fib 1 = 1\n    fib 2 = 2\n    fib x = fib (x-1) + fib (x-2)\nfib 10 -- 89\n\n-- Use underscore to match any, where you don't care about the binding name\nlet isZero 0 = true\n    isZero _ = false\n\n-- Pattern matching on records\nlet ecoTitle {author = \"Umberto Eco\", title = t} = Just t\n    ecoTitle _ = Nothing\n\necoTitle book -- Just (\"Foucault's pendulum\")\necoTitle {title: \"The Quantum Thief\", author: \"Hannu Rajaniemi\"} -- Nothing\n-- ecoTitle requires both field to type check:\necoTitle {title: \"The Quantum Thief\"} -- Object lacks required property \"author\"\n\n-- Lambda expressions\n(\\x -> x*x) 3 -- 9\n(\\x y -> x*x + y*y) 4 5 -- 41\nlet sqr = \\x -> x*x\n\n-- Currying\nlet myAdd x y = x + y -- is equivalent with\nlet myAdd' = \\x -> \\y -> x + y\nlet add3 = myAdd 3\n:t add3 -- Prim.Int -> Prim.Int\n\n-- Forward and backward function composition\n-- drop 3 followed by taking 5\n(drop 3 >>> take 5) (1..20) -- [4,5,6,7,8]\n-- take 5 followed by dropping 3\n(drop 3 <<< take 5) (1..20) -- [4,5]\n\n-- Operations using higher order functions\nlet even x = x `mod` 2 == 0\nfilter even (1..10) -- [2,4,6,8,10]\nmap (\\x -> x + 11) (1..5) -- [12,13,14,15,16]\n\n-- Requires purescript-foldable-traversable (Data.Foldable)\n\nfoldr (+) 0 (1..10) -- 55\nsum (1..10) -- 55\nproduct (1..10) -- 3628800\n\n-- Testing with predicate\nany even [1,2,3] -- true\nall even [1,2,3] -- false\n\n```"