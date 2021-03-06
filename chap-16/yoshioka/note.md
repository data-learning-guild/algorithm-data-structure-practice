# 16章考察メモ
## 16.2
- 参考URL
  - https://drken1215.hatenablog.com/entry/2021/08/05/013400
- "s-t間の最短路長を伸ばす"
  - シンプルに最短路を求めて、最短路を構成する辺のうち最小コストの辺を伸ばせばいい気がする
  - 複数の最短路がある場合は、それではうまくいかない
    - 複数ある場合は、すべての最短路を伸ばさないと行けない
- 一方で最短路に関係ない辺をといても仕方がない
  - "最短路に含まれる可能性のある辺を抽出する"
    - Floyd-Warshall法などによって、全頂点間の最短路 (2頂点 u,v 間の距離を dist[u][v] とする) を求めておく
      - フロイドワーシャル法：p264 14.7節
      - すでにある2点間について、間にさらにノードを入れてみて、ソッチのほうが短いならそれを採用する
    - 辺 e=(u,v) が「いずれかの最短経路に含まれるか」を判定するためには、dist[s][u] + c(e) + dist[v][t] == dist[s][t] かどうかを調べればよい
      - 距離テーブルとDPコストが一致するなら最短路を構成する辺
- "最短路に含まれる可能性のある辺のグラフ"を上記プロセスで構築
  - このグラフでの最小カットが最小コストと一致する
    - 最小カット=s-t間で、もっともコストが少ない場所

## 16.4
- 互いに素：二つの整数が1または－1以外の公約数をもたないこと。
- 互いに素ではないケースだけ、エッジをつくる
  - 最小カットが最大のマッチング数
- 互いに素ではないことをどうやって調べるか？
  - ユーグリットの互除法（GCD）の結果が１のとき、互いに素である
    - https://qiita.com/drken/items/0c88a37eec520f82b788#4-%E4%BA%92%E3%81%84%E3%81%AB%E7%B4%A0%E3%81%B8%E3%81%AE%E5%B8%B0%E7%9D%80
    - code4-4にて実装済み

## 16.6
- プロジェクト選択問題
  - 整数xは複数ありうる
  - 最大値を知りたい
    - うまく負の値の宝石を効率よく割る
- 宝石を割る/割らない = {0,1}
  - 操作を行ったあとのsum(a_n)を求める
- 制約条件
  - 整数xが約数; j % i == 0
  - すべて割る; x_i = 0 && x_jのコストは∞
- 割ってしまう正の宝石と割らなかった負の宝石を最小化する
  - 得られずはずの最大値から、割ってしまった値を引く