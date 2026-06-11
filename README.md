# db-benchmark-api

Go + Gin を使用した、MySQL と PostgreSQL の簡易ベンチマーク用 API です。

## 概要

本プロジェクトは、同一データセット（約100万件の顧客データ）を用いて、MySQL と PostgreSQL の検索・集計性能を比較することを目的としています。

APIを通して各DBにアクセスし、応答速度や取得結果を確認します。

---

## 使用技術

- Go
- Gin
- MySQL（Docker）
- PostgreSQL（Docker）

---

## データセットについて

本プロジェクトでは、以下の公開サンプルデータセットを使用しています。

- データセット名: Datablist Sample CSV datasets
- 提供元: https://www.datablist.com/
- データ生成: Faker によるランダム生成データ
- 利用データ: customers-1000000.csv
- 用途: DB（MySQL / PostgreSQL）性能比較の学習・検証

※本データは実在の個人情報を含まないダミーデータです。

---

## 目的

- GoによるAPI開発の基礎理解
- Ginフレームワークの学習
- RDB（MySQL / PostgreSQL）の比較検証
- 大規模データでのクエリ性能の確認

---

## 今後の実装予定

- MySQL / PostgreSQL 接続API
- レコード件数取得API
- 条件検索API
- 実行時間比較（ベンチマーク）
