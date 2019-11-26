#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
#Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

Gui, Add, Text, x30, 対象ファイル:
Gui, Add, Edit, x100 yp+0 vTarget w400,
Gui, Add, Button, gExec x465 yp+20, &実行
Gui, Add, Button, gButton終了 x415 yp+0, &終了
Gui, Show, , 半スペチェック
return

Exec:
Gui, Submit, NoHide
SplitPath, Target, file, dir, ext, name_no_ext, drive

;フォルダ指定が空の場合は警告
if Target =
{
	MsgBox, 0x30, 警告, フルパスでファイルを指定してください。
	return
}
;フォルダ指定がドライブ直下の場合も警告
Else if (RegExMatch(Target, "i)^[a-zA-Z]:?\\?$"))
{
	MsgBox, 0x10, エラー！, ドライブ直下です。
	return
}

Sleep, 30

;exeを叩く
cmd = go_check.exe "%Target%"
Run, %cmd%

return

;ウィンドウにフォルダがドロップされたときに実行
GuiDropFiles:
StringSplit, fn, A_GuiEvent, `n    ;フォルダを一つにする	
GuiControl, , Target, %fn1%        ;エディットボックスに一つめのフォルダ名を設定
return

;[終了]と[x]ボタン
Button終了:
GuiClose:
ExitApp