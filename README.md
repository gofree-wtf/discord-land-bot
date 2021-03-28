discord-land-bot
================

관심 지역의 부동산을 실시간으로 조회하여 디스코드 채널에 뿌려주는 봇


## Usage

TODO


## Development

### Data Flow

- Context #1
  - 주기적으로 공공 API 데이터를 MongoDB로 덤프
- Context #2
  - 디스코드 연결
  - MongoDB에 있는 거래건을 시간 순으로 디스코드 채널에 출력
  - 출력한 거래건은 마킹 (앱을 재시작해도 다시 출력하지 않게끔)

### XML to Struct

XML을 Struct로 한땀한땀 만들기 귀찮으므로 `zek` 툴을 사용한다.

```bash
$ echo "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?><response><header><resultCode>99</resultCode><resultMsg>SERVICE KEY IS NOT REGISTERED ERROR.</resultMsg></header></response>" | zek -e
// Response was generated 2021-03-28 18:00:47 by inter6 on gofree-mac.local.
type Response struct {
        XMLName xml.Name `xml:"response"`
        Text    string   `xml:",chardata"`
        Header  struct {
                Text       string `xml:",chardata"`
                ResultCode string `xml:"resultCode"` // 99
                ResultMsg  string `xml:"resultMsg"`  // SERVICE KEY IS NOT REGIST...
        } `xml:"header"`
} 
```
