package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func main() {
	// 引数1が処理対象ファイル
	arg := os.Args[1]

	// 読み込み用のファイルをオープンする
	fp, err := os.Open(arg)
	if err != nil {
		panic(err)
	}

	// 書き込みファイル名
	result := "result.html"
	result_apath, _ := filepath.Abs(result)
	//fmt.Println("Check results: " + result_apath)

	// 書き込み用のファイルを作成しつつオープンする
	fw, err2 := os.Create(result)
	if err2 != nil {
		panic(err2)
	}

	// 読み込み用にオープンしたファイルを変数に入れる
	scanner := bufio.NewScanner(fp)
	// 書き込み用にオープンしたファイルを変数に入れる
	writer := bufio.NewWriter(fw)

	// htmlのヘッダー
	header := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"  xml:lang="ja" lang="ja">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<style Type="text/css">
	body{
		margin-right: auto;
		margin-left: auto;
		width: 1200px;
		font-family: 'Arial', 'Hiragino Kaku Gothic Pro', 'ヒラギノ角ゴ Pro W3', 'MS PGothic', 'ＭＳ Ｐゴシック', sans-serif;
	}

	.table {
		width: 1200px;
		border-collapse: collapse;
	}

	tr:nth-child(2n+1) {
		background: #f5f5f5;
	}
</style>
<title>Error results</title>
<body>
<table class="table" border="0" cellpadding="5">
<tr><th bgcolor="#FFE4E1"><p><font color="#880000">ERROR</font></p></th></tr>`

	// htmlのfooter
	footer := `</table>
</body>
</html>`

	// ヘッダーを出力
	fmt.Fprint(writer, header)

	// 半スペチェック
	for scanner.Scan() {
		str := scanner.Text()
		// 正規表現と文字列をcheck_regexp関数に投げる
		check_regexp(`( 。)`, str, writer)
		check_regexp(`(、、)`, str, writer)
		check_regexp(`(。 )`, str, writer)
		check_regexp(`(、 )`, str, writer)
		check_regexp(`(てて|にに|をを|はは|すす|よよ|んん|ュュ|ョョ|ッッ|ンン|っっ)`, str, writer)
		check_regexp(`( 、)`, str, writer)
		check_regexp(`(。。)`, str, writer)
		check_regexp(`(\p{Han} \p{Katakana})`, str, writer)
		check_regexp(`(\p{Katakana} \p{Han})`, str, writer)
		check_regexp(`(\p{Han} \p{Hiragana})`, str, writer)
		check_regexp(`(\p{Hiragana} \p{Han})`, str, writer)
		check_regexp(`(\p{Katakana} \p{Hiragana})`, str, writer)
		check_regexp(`(\p{Hiragana} \p{Katakana})`, str, writer)
		check_regexp(`([a-zA-Z)]\p{Katakana})`, str, writer)
		check_regexp(`(\p{Katakana}[a-zA-Z(])`, str, writer)
		check_regexp(`([a-zA-Z)]\p{Han})`, str, writer)
		check_regexp(`(\p{Han}[a-zA-Z(])`, str, writer)
		check_regexp(`([a-zA-Z)]\p{Hiragana})`, str, writer)
		check_regexp(`(\p{Hiragana}[a-zA-Z(])`, str, writer)
		check_regexp(`( 「|」 |「 | 」)`, str, writer)
	}

	// フッターを出力
	fmt.Fprint(writer, footer)

	writer.Flush()

	fp.Close()
	fw.Close()

	// チェック結果をブラウザで表示
	exec.Command("cmd.exe", "/c", "start ", "\"", result_apath, "\"").Run()
}

func check_regexp(reg, str string, writer *bufio.Writer) {
	if regexp.MustCompile(reg).MatchString(str) == true {
		re := regexp.MustCompile(reg)
		str := re.ReplaceAllString(str, "<font color=\"red\">$1</font>")
		// result.htmlに書き込む
		fmt.Fprint(writer, "<tr><td><p>"+str+"</p></td></tr>"+"\n")
	}
}
