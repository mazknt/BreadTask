# BreadTask

##確認手順

- `.env`をルートに配置し、以下を設定する
  - SPACE_ID
  - HONEY_SOY_CRAN(蜂蜜豆乳クランベリーの entry_id)
  - BLACK_SESAME_POTE(黒ゴマポテロールの entry_id)
  - SHICHIMI_SALT_FOCACCIA(黒七味と岩塩のフォカッチャの entry_id)
  - ACCESS_TOKEN
  - PROJECT_ID(Firestore の project id)
  - CREDENTIAL_OPTION(Firestore の credential option)
- `go build`でビルドする
- `./task1 fetchBread`で今回作成した関数`fetchBread`を呼び出す
- Firestore にデータが保存されているかを確認する
