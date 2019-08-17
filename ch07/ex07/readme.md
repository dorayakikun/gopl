> 20.0のデフォルト値はを含んでいないのに、ヘルプメッセージがを含んでいる理由を説明しなさい。

標準出力時に `func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }` が暗黙的に呼び出されているため。

`print.go` の `handleMethods` で `Stringer` 実装している場合のケースが実行されている。