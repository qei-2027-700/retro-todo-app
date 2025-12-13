これは、個人・私生活用のタスク管理・振返用アプリケーションです。
WEBアプリですが、将来的にはスマホアプリ開発も考えています。

バックエンドがGo/Echo
フロントエンドがVue3/Nuxt4
の構成のモノレポ構成です。

## バックエンド
/backend
httpメソッドの`PATCH` は使用しないでください。

## フロントエンド
/frontend-nuxt

tailwindcss を利用しています。
classは静的な場合は必ず `class=` で定義し、 `:class=` は動的な場合のみ利用して。

### 自動importについて
nuxt4なので、下記を意識して
- vueファイルにてref, computed は自動importされるので記載しないで
- utils, coposables, components, layoutsは自動importされるはずです

### components ディレクトリについて
下記のように分割しています。
- common
- ui
- 機能別

### フォーマット
if-elseやtry-catchは下記のフォーマットで記載してください。

```ts
if (result.success) {
  // 
}
else {
  //
}
```

