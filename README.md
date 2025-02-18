# Cnyes stock news
這是一個可以一行指令查詢個股相關新聞的工具，新聞來源為鉅亨網

## Quick start
可以查詢個股新聞或分類新聞，多個代碼請用`,`分隔，輸出過去 24 小時內發布的新聞標題及其連結
### 個股新聞 stock command
一個個股最多顯示 30 篇新聞

**指令**
```
./cnyes-stock-news news stock 2330,2317
```
輸出
```
------2317------
1. 02/18 外資近5日買超個股 https://news.cnyes.com/news/id/5864849
2. 02/18 外資近3日買超個股 https://news.cnyes.com/news/id/5864847
3. 02/17 外資買超247億元敲進鴻海2.47萬張 同步回補台積電近萬張 https://news.cnyes.com/news/id/5864678
------2330------
1. 02/18 〈台股開盤〉創開春新高後觀望氣氛濃 生醫、塑化股重挫 中小型股活躍 https://news.cnyes.com/news/id/5864958        
2. 02/18 國產全建材集團策略穩健成長 旗下國宇、惠普營運續創新高 https://news.cnyes.com/news/id/5864908
3. 02/18 外資近5日賣超個股 https://news.cnyes.com/news/id/5864852
```

### 分類新聞 category command
一個分類最多顯示 50 篇新聞

分類代號請從鉅亨網分類新聞的 url 取得。例如美股雷達的 url 為：
```
https://news.cnyes.com/news/cat/us_stock
```
則分類代號為 `us_stock`

**指令**
```
./cnyes-stock-news news category tw_stock,us_stock
```
輸出
```
------us_stock------
1. 02/18 〈吳田玉開講〉日月光大搶AI商機 面板級封裝Q3試量產 https://news.cnyes.com/news/id/5865099
2. 02/18 大決戰！Grok-3發表在即OpenAI緊盯 傳奧特曼可能推GPT-4.5狙擊馬斯克xAI https://news.cnyes.com/news/id/5865049     
3. 02/17 〈吳田玉開講〉AI將加劇國與國間抗衡 日月光靠三策略應對 https://news.cnyes.com/news/id/5864657
------tw_stock------
1. 02/18 高雄左營店供電異常 新光三越：恢復正常供電 https://news.cnyes.com/news/id/5865158
2. 02/18 【一分鐘看圖論市】川普追求貿易公平，美國出台對等關稅 https://news.cnyes.com/news/id/5865108
3. 02/18 〈吳田玉開講〉日月光大搶AI商機 面板級封裝Q3試量產 https://news.cnyes.com/news/id/5865099
```

## 安裝
```
git clone git@github.com:FallPrediction/cnyes-stock-news.git
```
編譯檔請查看 [release 頁面](https://github.com/FallPrediction/cnyes-stock-news/releases)

## 額外功能
news command 有三個 flag 可使用：
```
-e, --exclude strings   新聞標題不可包含特定字串，以 , 分隔
-i, --include strings   新聞標題須包含特定字串，以 , 分隔
-r, --regex string      新聞標題須符合正規表達式
```
若正規表達式不合規則，正常輸出新聞。

若任一條件不符合就不輸出新聞，舉例：
```
./cnyes-stock-news news stock 2330 -i="台股,賣超" -e="賣超"

外資近5日買超個股 <- 不會輸出這篇新聞
台股開盤 <- 會輸出這篇新聞
```
