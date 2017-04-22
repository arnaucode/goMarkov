# goMarkov
markov chains text generator written in Go from scratch


```go
states := markov.train(text)
generatedText := markov.generateText(states, firstWord, count)
fmt.Println(generatedText)
```
(in the text variable, goes the text content, can be loaded from a .txt file)



![Argos](https://raw.githubusercontent.com/arnaucode/goMarkov/master/goMarkov.gif "goMarkov")
