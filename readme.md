# epay-auto-punch
ePayWorkへの勤怠打刻を行います。(※作者の勤務先以外での動作確認はしていません。)

実行した日の出勤時間の入力が無ければ現在時刻を出勤時刻として入力、あれば現在時刻を退勤時間として打刻します。

# 必要なもの
### Google Chrome

(略)

### golang

(略)

### agouti

```
go get github.com/sclevine/agouti
```

### ChromeDriver

macの場合は以下

```
brew cask install chromedriver
```

# 実行方法

### 環境変数に以下を設定

- ePayWorkCopCd : ePayWorkログイン時の企業コード
- ePayWorkEmpCd : ePayWorkログイン時の社員番号
- ePayWorkPassword : ePayWorkログイン時のパスワード

### 以下を実行

```
go run main.go
```